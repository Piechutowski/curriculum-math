// Command progress scans the curriculum module files for markdown checkboxes
// ("- [ ]" / "- [x]"), then:
//
//   - renders an SVG progress-bar card per module plus an overall card into progress/
//   - refreshes the "N / M complete" block in each module file and the README dashboard
//   - maintains a completion log: progress/log.csv records when each lesson was first
//     ticked, and LOG.md presents it newest-day-first (unticking a lesson removes it)
//   - writes the raw counts to progress/progress.csv
//   - counts numbered exercises on exercise rows (see exerciseLine) and, for
//     modules that have them, renders one card per chapter with a labelled
//     segment per exercise set (<id>-exercises-chNN.svg), plus a module-wide
//     <id>-exercises.svg whose segments are the chapters
//
// Run it from the repository root after ticking checkboxes:
//
//	go run ./tools/progress
package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// group is one "## " heading in a module file — a course section or book
// chapter. Its boundaries become milestone notches on the progress bar.
type group struct {
	done  int
	total int
}

type module struct {
	id      string // also the SVG file basename and the progress-marker id
	file    string
	title   string
	short   string // compact name used in the log
	unit    string // what a "## " heading is called in this module
	c1, c2  string // gradient start/end
	done    int
	total   int
	checked []string // labels of ticked lessons, in file order
	groups  []group  // one per "## " heading that contains checkboxes

	exDone, exTotal int
	exChapters      []*exChapter // "## " headings that contain exercise rows
}

// exChapter is one "## " heading's exercise rows; each row is one set (a
// section's exercises, or Review Questions, Supplementary Exercises, ...) and
// gets its own labelled segment on the chapter's card.
type exChapter struct {
	num   int    // from "## Ch. N", else the heading's position in the file
	name  string // "Ch. 1"
	sets  []exSet
	done  int
	total int
}

type exSet struct {
	label string // "1.1" for a section's Exercises row, initials ("RQ") otherwise
	done  int
	total int
}

var modules = []*module{
	{id: "01-algebra-trigonometry", file: "01-algebra-trigonometry-map.md", title: "01 · Algebra & Trigonometry — Greene", short: "Greene", unit: "sections", c1: "#58a6ff", c2: "#1f6feb"},
	{id: "02-discrete-mathematics", file: "02-discrete-mathematics-map.md", title: "02 · Discrete Mathematics — Rosen", short: "Rosen", unit: "chapters", c1: "#bc8cff", c2: "#8957e5"},
	{id: "03-discrete-math-lectures", file: "03-discrete-math-lectures.md", title: "03 · Discrete Math — Lecture Course", short: "Lectures", unit: "sections", c1: "#ff7b72", c2: "#da3633"},
	{id: "04-calculus", file: "04-calculus.md", title: "04 · Calculus — Thomas", short: "Calculus", unit: "sections", c1: "#ffa657", c2: "#f0883e"},
	{id: "05-linear-algebra", file: "05-linear-algebra.md", title: "05 · Linear Algebra — Kumaresan", short: "Linear Algebra", unit: "sections", c1: "#56d364", c2: "#2ea043"},
	{id: "06-probability", file: "06-probability.md", title: "06 · Probability — Veerarajan", short: "Probability", unit: "sections", c1: "#f778ba", c2: "#db61a2"},
	{id: "07-numerical-methods", file: "07-numerical-methods.md", title: "07 · Numerical Methods — Sastry", short: "Numerical", unit: "sections", c1: "#76e3ea", c2: "#39c5cf"},
	{id: "08-pl-theory", file: "08-pl-theory.md", title: "08 · PL Theory — Sipser · Pierce · Winskel", short: "PL Theory", unit: "sections", c1: "#e3b341", c2: "#d29922"},
}

// cardVersion changes whenever the SVG design changes, so the ?v= cache-buster
// in image URLs also changes and viewers do not keep an old-looking card.
const cardVersion = 4

const logCSV = "progress/log.csv"

type logEntry struct {
	ts     time.Time
	module string // module id
	label  string
}

var checkboxLine = regexp.MustCompile(`^\s*[-*] \[([ xX])\] (.+)$`)

// exerciseLine is an indented bullet under a lesson listing that lesson's
// numbered exercises, e.g. "  - Exercises: [x] 1 · [ ] 2 · [ ] 3". The bullet
// has no checkbox of its own, so checkboxLine never counts it as a lesson.
var exerciseLine = regexp.MustCompile(`^\s+[-*] ([A-Za-z][A-Za-z ]*): \[[ xX]\] 1\b`)

var chapterHeading = regexp.MustCompile(`^## Ch\. (\d+)`)

var exerciseBox = regexp.MustCompile(`\[([ xX])\] \d+`)

func main() {
	if _, err := os.Stat("README.md"); err != nil {
		fatal("run this from the repository root (README.md not found)")
	}

	byID := map[string]*module{}
	for _, m := range modules {
		byID[m.id] = m
		src, err := os.ReadFile(m.file)
		if err != nil {
			fatal("reading %s: %v", m.file, err)
		}
		var ch *exChapter // exercise rows of the current "## " heading
		var lesson string // content of the last lesson line, for set labels
		headings := 0
		for _, line := range strings.Split(string(src), "\n") {
			if strings.HasPrefix(line, "## ") {
				m.groups = append(m.groups, group{})
				headings++
				ch = &exChapter{num: headings, name: fmt.Sprintf("Part %d", headings)}
				if hit := chapterHeading.FindStringSubmatch(line); hit != nil {
					fmt.Sscan(hit[1], &ch.num)
					ch.name = "Ch. " + hit[1]
				}
				continue
			}
			if hit := exerciseLine.FindStringSubmatch(line); hit != nil {
				if ch == nil {
					ch = &exChapter{num: 0, name: "Intro"}
				}
				if len(ch.sets) == 0 {
					m.exChapters = append(m.exChapters, ch)
				}
				set := exSet{label: setLabel(hit[1], lesson)}
				for _, box := range exerciseBox.FindAllStringSubmatch(line, -1) {
					set.total++
					if box[1] != " " {
						set.done++
					}
				}
				ch.sets = append(ch.sets, set)
				ch.done += set.done
				ch.total += set.total
				m.exDone += set.done
				m.exTotal += set.total
				continue
			}
			hit := checkboxLine.FindStringSubmatch(line)
			if hit == nil {
				continue
			}
			if len(m.groups) == 0 { // checkbox before any heading
				m.groups = append(m.groups, group{})
			}
			lesson = label(hit[2])
			g := &m.groups[len(m.groups)-1]
			m.total++
			g.total++
			if hit[1] != " " {
				m.done++
				g.done++
				m.checked = append(m.checked, label(hit[2]))
			}
		}
		// Headings with no checkboxes (exit criteria, tables) are not milestones.
		kept := m.groups[:0]
		for _, g := range m.groups {
			if g.total > 0 {
				kept = append(kept, g)
			}
		}
		m.groups = kept
	}

	var done, total int
	for _, m := range modules {
		done += m.done
		total += m.total
	}

	if err := os.MkdirAll("progress", 0o755); err != nil {
		fatal("creating progress dir: %v", err)
	}

	added, removed := updateLog(byID)

	for _, m := range modules {
		writeFile(filepath.Join("progress", m.id+".svg"), card(m.title, m.done, m.total, m.c1, m.c2, m.groups, m.unit))
		if m.exTotal > 0 {
			chapters := make([]group, len(m.exChapters))
			for i, c := range m.exChapters {
				chapters[i] = group{done: c.done, total: c.total}
				writeFile(filepath.Join("progress", exCardName(m, c)), setCard(m.short+" "+c.name+" · exercises", c, m.c1, m.c2))
			}
			writeFile(filepath.Join("progress", m.id+"-exercises.svg"), card(m.title+" · exercises", m.exDone, m.exTotal, m.c1, m.c2, chapters, m.unit))
		}
	}
	// The overall bar's segments are the modules themselves.
	overallGroups := make([]group, len(modules))
	for i, m := range modules {
		overallGroups[i] = group{done: m.done, total: m.total}
	}
	writeFile(filepath.Join("progress", "overall.svg"), card("Overall — all modules", done, total, "#e3b341", "#d29922", overallGroups, "modules"))
	writeFile(filepath.Join("progress", "progress.csv"), countsCSV(done, total))

	for _, m := range modules {
		updateBlock(m.file, "progress:"+m.id, moduleBlock(m))
	}
	updateBlock("README.md", "progress:overall", readmeBlock(done, total))

	fmt.Printf("%-40s %9s %6s\n", "module", "done", "")
	for _, m := range modules {
		fmt.Printf("%-40s %4d /%4d %5d%%\n", m.title, m.done, m.total, pct(m.done, m.total))
		if m.exTotal > 0 {
			fmt.Printf("%-40s %4d /%4d %5d%%\n", "  exercises", m.exDone, m.exTotal, pct(m.exDone, m.exTotal))
		}
	}
	fmt.Printf("%-40s %4d /%4d %5d%%\n", "overall", done, total, pct(done, total))
	if added > 0 || removed > 0 {
		fmt.Printf("log: %d completion(s) added, %d removed\n", added, removed)
	}
}

// setLabel names an exercise set on its chapter card: a section's plain
// "Exercises" row takes the section number from its lesson line ("1.1 ..."),
// any other row is abbreviated to its initials ("Review Questions" -> "RQ").
func setLabel(row, lesson string) string {
	if row == "Exercises" {
		if f := strings.Fields(lesson); len(f) > 0 {
			return f[0]
		}
	}
	var b strings.Builder
	for _, w := range strings.Fields(row) {
		if w == "and" || w == "of" {
			continue
		}
		b.WriteByte(w[0])
	}
	return strings.ToUpper(b.String())
}

func exCardName(m *module, c *exChapter) string {
	return fmt.Sprintf("%s-exercises-ch%02d.svg", m.id, c.num)
}

// setCard renders a chapter's exercise card. Unlike card, segment widths are
// not strictly proportional: every set gets a minimum width so even a
// six-exercise set stays visible and labelled, and the rest of the bar is
// shared out by size. Each segment fills by its own set's progress, and its
// label sits under it. The header still states the exact count.
func setCard(title string, c *exChapter, c1, c2 string) string {
	const (
		w, h    = 640, 84
		pad     = 20
		barY    = 34
		barH    = 10
		notchW  = 2
		minSeg  = 26
		bgColor = "#0d1117"
		font    = `-apple-system,'Segoe UI',Helvetica,Arial,sans-serif`
	)
	barW := w - 2*pad
	n := len(c.sets)
	minW := minSeg
	if n*minW > barW {
		minW = barW / n
	}
	spare := barW - n*minW

	setsDone := 0
	for _, s := range c.sets {
		if s.done == s.total {
			setsDone++
		}
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" role="img" aria-label="%s: %d of %d exercises (%d of %d sets complete)">`+"\n",
		w, h, w, h, esc(title), c.done, c.total, setsDone, n)
	fmt.Fprintf(&b, `  <defs>`+"\n")
	fmt.Fprintf(&b, `    <linearGradient id="g" x1="0" y1="0" x2="1" y2="0"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient>`+"\n", c1, c2)
	fmt.Fprintf(&b, `    <clipPath id="c"><rect x="%d" y="%d" width="%d" height="%d" rx="4"/></clipPath>`+"\n", pad, barY, barW, barH)
	fmt.Fprintf(&b, `  </defs>`+"\n")
	fmt.Fprintf(&b, `  <rect width="%d" height="%d" rx="12" fill="%s" stroke="#30363d"/>`+"\n", w, h, bgColor)
	fmt.Fprintf(&b, `  <text x="%d" y="22" font-family="%s" font-size="13" font-weight="600" fill="#e6edf3">%s</text>`+"\n", pad, font, esc(title))
	fmt.Fprintf(&b, `  <text x="%d" y="22" text-anchor="end" font-family="%s" font-size="12" fill="#9198a1">%d / %d · %d%%</text>`+"\n", w-pad, font, c.done, c.total, pct(c.done, c.total))

	// Segment bounds: minW each, plus a share of the spare width by set size.
	xs := make([]int, n+1)
	cum := 0
	for i, st := range c.sets {
		xs[i] = pad + i*minW + spare*cum/c.total
		cum += st.total
	}
	xs[n] = pad + barW

	fmt.Fprintf(&b, `  <g clip-path="url(#c)">`+"\n")
	for i, st := range c.sets {
		x0, x1 := xs[i], xs[i+1]
		track := "#21262d" // alternate shades so neighbouring sets read apart
		if i%2 == 1 {
			track = "#2d333b"
		}
		fmt.Fprintf(&b, `    <rect x="%d" y="%d" width="%d" height="%d" fill="%s"/>`+"\n", x0, barY, x1-x0, barH, track)
		if st.done > 0 {
			segW := (x1 - x0) * st.done / st.total
			if segW < 2 {
				segW = 2
			}
			fmt.Fprintf(&b, `    <rect x="%d" y="%d" width="%d" height="%d" fill="url(#g)"/>`+"\n", x0, barY, segW, barH)
		}
	}
	fmt.Fprintf(&b, `  </g>`+"\n")
	for i, st := range c.sets {
		x0, x1 := xs[i], xs[i+1]
		if i > 0 {
			fmt.Fprintf(&b, `  <rect x="%d" y="%d" width="%d" height="%d" fill="%s"/>`+"\n", x0-notchW/2, barY, notchW, barH, bgColor)
		}
		color := "#7d8590"
		if st.done == st.total {
			color = c1
		}
		fmt.Fprintf(&b, `  <text x="%d" y="%d" text-anchor="middle" font-family="%s" font-size="9.5" fill="%s">%s</text>`+"\n", (x0+x1)/2, barY+barH+13, font, color, esc(st.label))
	}
	fmt.Fprintf(&b, `  <text x="%d" y="76" font-family="%s" font-size="10.5" fill="#7d8590">%d of %d sets complete</text>`+"\n", pad, font, setsDone, n)
	b.WriteString("</svg>\n")
	return b.String()
}

// label extracts a compact lesson name from a checkbox line's content: the bold
// span when the lesson has a long description (modules 03-06), the whole line
// otherwise (Greene and Rosen entries).
func label(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "**") {
		if end := strings.Index(s[2:], "**"); end >= 0 {
			return strings.TrimSpace(s[2 : 2+end])
		}
	}
	return s
}

// updateLog reconciles progress/log.csv with the currently ticked lessons:
// newly ticked lessons are stamped with the current time, unticked ones are
// dropped. It then regenerates LOG.md. Returns (added, removed).
func updateLog(byID map[string]*module) (int, int) {
	prev := readLog()
	prevKey := map[string]logEntry{}
	for _, e := range prev {
		prevKey[e.module+"|"+e.label] = e
	}

	now := time.Now().UTC().Truncate(time.Minute)
	var next []logEntry
	added := 0
	for _, m := range modules {
		for _, l := range m.checked {
			if e, ok := prevKey[m.id+"|"+l]; ok {
				next = append(next, e)
			} else {
				next = append(next, logEntry{ts: now, module: m.id, label: l})
				added++
			}
		}
	}
	removed := len(prev) + added - len(next)

	sort.SliceStable(next, func(i, j int) bool { return next[i].ts.Before(next[j].ts) })

	var b strings.Builder
	w := csv.NewWriter(&b)
	w.Write([]string{"completed_at_utc", "module", "lesson"})
	for _, e := range next {
		w.Write([]string{e.ts.Format(time.RFC3339), e.module, e.label})
	}
	w.Flush()
	writeFile(logCSV, b.String())

	writeFile("LOG.md", logPage(next, byID))
	return added, removed
}

func readLog() []logEntry {
	f, err := os.Open(logCSV)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	if err != nil {
		fatal("reading %s: %v", logCSV, err)
	}
	defer f.Close()
	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fatal("parsing %s: %v", logCSV, err)
	}
	var out []logEntry
	for i, r := range rows {
		if i == 0 || len(r) != 3 { // header
			continue
		}
		ts, err := time.Parse(time.RFC3339, r[0])
		if err != nil {
			fatal("%s row %d: bad timestamp %q", logCSV, i+1, r[0])
		}
		out = append(out, logEntry{ts: ts, module: r[1], label: r[2]})
	}
	return out
}

// logPage renders LOG.md: days newest-first, lessons within a day in the order
// they were logged.
func logPage(entries []logEntry, byID map[string]*module) string {
	var b strings.Builder
	b.WriteString("# Completion Log\n\n")
	b.WriteString("When each lesson was completed — newest day first. Maintained by\n")
	b.WriteString("`go run ./tools/progress`; the raw data lives in [progress/log.csv](progress/log.csv).\n")
	b.WriteString("Do not edit this file by hand.\n\n")

	if len(entries) == 0 {
		b.WriteString("*Nothing logged yet — tick your first checkbox and rerun the tool.*\n")
		return b.String()
	}
	fmt.Fprintf(&b, "**%d lessons completed so far.**\n", len(entries))

	// Group by UTC date, preserving chronological order within each day.
	var days []string
	byDay := map[string][]logEntry{}
	for _, e := range entries {
		d := e.ts.Format("2006-01-02")
		if _, ok := byDay[d]; !ok {
			days = append(days, d)
		}
		byDay[d] = append(byDay[d], e)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(days)))

	for _, d := range days {
		es := byDay[d]
		noun := "lessons"
		if len(es) == 1 {
			noun = "lesson"
		}
		fmt.Fprintf(&b, "\n## %s — %d %s\n\n", d, len(es), noun)
		for _, e := range es {
			short := e.module
			if m, ok := byID[e.module]; ok {
				short = m.short
			}
			fmt.Fprintf(&b, "- **%s** — %s\n", short, e.label)
		}
	}
	return b.String()
}

func pct(done, total int) int {
	if total == 0 {
		return 0
	}
	return int(float64(done)/float64(total)*100 + 0.5)
}

// card renders one self-contained progress-bar card (dark, GitHub-friendly).
//
// The bar is notched at every group boundary, so each segment is one section of
// the course (or chapter of the book). Each segment fills independently, in its
// own place along the bar: studying out of order (trigonometry before the
// algebra sections, say) lights up the segments you actually did rather than
// filling from the left. Total filled area still equals done/total, so the bar
// reads as a percentage and as a map at the same time.
func card(title string, done, total int, c1, c2 string, groups []group, unit string) string {
	const (
		w, h    = 640, 64
		pad     = 20
		barY    = 34
		barH    = 10
		notchW  = 2 // background-colored gap drawn at each group boundary
		minFill = 2 // so a single ticked lesson in a long section still shows
		bgColor = "#0d1117"
	)
	barW := w - 2*pad
	p := pct(done, total)

	groupsDone := 0
	for _, g := range groups {
		if g.done == g.total {
			groupsDone++
		}
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" role="img" aria-label="%s: %d%% (%d of %d %s complete)">`+"\n",
		w, h, w, h, esc(title), p, groupsDone, len(groups), unit)
	fmt.Fprintf(&b, `  <defs>`+"\n")
	// Object-bounding-box units (the default): the gradient is scoped to each
	// filled rect, so the full c1→c2 sweep is visible however little is done.
	// Spanning it across the whole bar instead would wash the color out at low
	// percentages — at 19% you would only ever see the first fifth of the ramp.
	fmt.Fprintf(&b, `    <linearGradient id="g" x1="0" y1="0" x2="1" y2="0"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient>`+"\n", c1, c2)
	fmt.Fprintf(&b, `    <clipPath id="c"><rect x="%d" y="%d" width="%d" height="%d" rx="4"/></clipPath>`+"\n", pad, barY, barW, barH)
	fmt.Fprintf(&b, `  </defs>`+"\n")
	fmt.Fprintf(&b, `  <rect width="%d" height="%d" rx="12" fill="%s" stroke="#30363d"/>`+"\n", w, h, bgColor)
	fmt.Fprintf(&b, `  <text x="%d" y="22" font-family="-apple-system,'Segoe UI',Helvetica,Arial,sans-serif" font-size="13" font-weight="600" fill="#e6edf3">%s</text>`+"\n", pad, esc(title))
	fmt.Fprintf(&b, `  <text x="%d" y="22" text-anchor="end" font-family="-apple-system,'Segoe UI',Helvetica,Arial,sans-serif" font-size="12" fill="#9198a1">%d / %d · %d%%</text>`+"\n", w-pad, done, total, p)
	fmt.Fprintf(&b, `  <rect x="%d" y="%d" width="%d" height="%d" rx="4" fill="#21262d"/>`+"\n", pad, barY, barW, barH)

	// Per-segment fills, clipped to the rounded track so the ends stay clean.
	if total > 0 && done > 0 {
		fmt.Fprintf(&b, `  <g clip-path="url(#c)">`+"\n")
		if len(groups) == 0 { // no headings: one plain fill
			fmt.Fprintf(&b, `    <rect x="%d" y="%d" width="%d" height="%d" fill="url(#g)"/>`+"\n", pad, barY, barW*done/total, barH)
		}
		cum := 0
		for _, g := range groups {
			x0 := pad + barW*cum/total
			cum += g.total
			x1 := pad + barW*cum/total
			if g.done == 0 {
				continue
			}
			segW := (x1 - x0) * g.done / g.total
			if segW < minFill {
				segW = minFill
			}
			fmt.Fprintf(&b, `    <rect x="%d" y="%d" width="%d" height="%d" fill="url(#g)"/>`+"\n", x0, barY, segW, barH)
		}
		fmt.Fprintf(&b, `  </g>`+"\n")
	}

	// Notches cut through both track and fill, separating the segments.
	cum := 0
	for i, g := range groups {
		cum += g.total
		if i == len(groups)-1 || total == 0 {
			break // no notch at the far right edge
		}
		x := pad + barW*cum/total
		fmt.Fprintf(&b, `  <rect x="%d" y="%d" width="%d" height="%d" fill="%s"/>`+"\n", x-notchW/2, barY, notchW, barH, bgColor)
	}
	if len(groups) > 0 {
		fmt.Fprintf(&b, `  <text x="%d" y="55" font-family="-apple-system,'Segoe UI',Helvetica,Arial,sans-serif" font-size="10.5" fill="#7d8590">%d of %d %s complete</text>`+"\n",
			pad, groupsDone, len(groups), unit)
	}
	b.WriteString("</svg>\n")
	return b.String()
}

// Image URLs carry a ?v=done-total query that changes whenever the counts
// change, so GitHub's image cache (camo) fetches a fresh SVG instead of
// serving a stale bar.
func moduleBlock(m *module) string {
	groupsDone := 0
	for _, g := range m.groups {
		if g.done == g.total {
			groupsDone++
		}
	}
	s := fmt.Sprintf("![%s progress](progress/%s.svg?v=%d-%d-%d)\n\n**%d / %d lessons complete · %d%%** — %d of %d %s finished · [completion log](LOG.md)",
		esc(m.title), m.id, cardVersion, m.done, m.total,
		m.done, m.total, pct(m.done, m.total), groupsDone, len(m.groups), m.unit)
	if m.exTotal > 0 {
		s += fmt.Sprintf("\n\n**%d / %d exercises solved · %d%%**\n", m.exDone, m.exTotal, pct(m.exDone, m.exTotal))
		for _, c := range m.exChapters {
			s += fmt.Sprintf("\n![%s %s exercises](progress/%s?v=%d-%d-%d)", esc(m.short), c.name, exCardName(m, c), cardVersion, c.done, c.total)
		}
	}
	return s
}

func readmeBlock(done, total int) string {
	var b strings.Builder
	fmt.Fprintf(&b, "![Overall progress](progress/overall.svg?v=%d-%d-%d)\n\n", cardVersion, done, total)
	for _, m := range modules {
		fmt.Fprintf(&b, "[![%s](progress/%s.svg?v=%d-%d-%d)](%s)\n", esc(m.title), m.id, cardVersion, m.done, m.total, m.file)
		if m.exTotal > 0 {
			fmt.Fprintf(&b, "[![%s exercises](progress/%s-exercises.svg?v=%d-%d-%d)](%s)\n", esc(m.title), m.id, cardVersion, m.exDone, m.exTotal, m.file)
		}
	}
	fmt.Fprintf(&b, "\n**Total: %d / %d lessons complete · %d%%** — [completion log](LOG.md) · raw numbers in [progress/progress.csv](progress/progress.csv)", done, total, pct(done, total))
	return b.String()
}

func countsCSV(done, total int) string {
	var b strings.Builder
	b.WriteString("module,title,done,total,percent,exercises_done,exercises_total\n")
	var exDone, exTotal int
	for _, m := range modules {
		fmt.Fprintf(&b, "%s,%q,%d,%d,%d,%d,%d\n", m.id, m.title, m.done, m.total, pct(m.done, m.total), m.exDone, m.exTotal)
		exDone += m.exDone
		exTotal += m.exTotal
	}
	fmt.Fprintf(&b, "overall,\"All modules\",%d,%d,%d,%d,%d\n", done, total, pct(done, total), exDone, exTotal)
	return b.String()
}

// updateBlock replaces the content between "<!-- id -->" and "<!-- endprogress -->" in file.
func updateBlock(file, id, content string) {
	src, err := os.ReadFile(file)
	if err != nil {
		fatal("reading %s: %v", file, err)
	}
	re := regexp.MustCompile(`(?s)<!-- ` + regexp.QuoteMeta(id) + ` -->.*?<!-- endprogress -->`)
	if !re.Match(src) {
		fatal("%s: marker <!-- %s --> ... <!-- endprogress --> not found", file, id)
	}
	block := "<!-- " + id + " -->\n" + content + "\n<!-- endprogress -->"
	out := re.ReplaceAllString(string(src), strings.ReplaceAll(block, "$", "$$"))
	if out != string(src) {
		writeFile(file, out)
	}
}

func esc(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", `"`, "&quot;")
	return r.Replace(s)
}

func writeFile(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		fatal("writing %s: %v", path, err)
	}
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "progress: "+format+"\n", args...)
	os.Exit(1)
}
