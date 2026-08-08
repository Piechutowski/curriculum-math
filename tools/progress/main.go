// Command progress scans the curriculum module files for markdown checkboxes
// ("- [ ]" / "- [x]"), then:
//
//   - renders an SVG progress-bar card per module plus an overall card into progress/
//   - refreshes the "N / M complete" block in each module file and the README dashboard
//   - writes the raw numbers to progress/progress.csv
//
// Run it from the repository root after ticking checkboxes:
//
//	go run ./tools/progress
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type module struct {
	id    string // also the SVG file basename and the progress-marker id
	file  string
	title string
	c1,c2 string // gradient start/end
	done  int
	total int
}

var modules = []*module{
	{id: "01-algebra-trigonometry", file: "01-algebra-trigonometry-map.md", title: "01 · Algebra & Trigonometry — Greene", c1: "#58a6ff", c2: "#1f6feb"},
	{id: "02-discrete-mathematics", file: "02-discrete-mathematics-map.md", title: "02 · Discrete Mathematics — Rosen", c1: "#bc8cff", c2: "#8957e5"},
	{id: "03-calculus", file: "03-calculus.md", title: "03 · Calculus — Thomas", c1: "#ffa657", c2: "#f0883e"},
	{id: "04-linear-algebra", file: "04-linear-algebra.md", title: "04 · Linear Algebra — Kumaresan", c1: "#56d364", c2: "#2ea043"},
	{id: "05-probability", file: "05-probability.md", title: "05 · Probability — Veerarajan", c1: "#f778ba", c2: "#db61a2"},
	{id: "06-numerical-methods", file: "06-numerical-methods.md", title: "06 · Numerical Methods — Sastry", c1: "#76e3ea", c2: "#39c5cf"},
}

var checkboxRe = regexp.MustCompile(`(?m)^\s*[-*] \[([ xX])\] `)

func main() {
	if _, err := os.Stat("README.md"); err != nil {
		fatal("run this from the repository root (README.md not found)")
	}

	for _, m := range modules {
		src, err := os.ReadFile(m.file)
		if err != nil {
			fatal("reading %s: %v", m.file, err)
		}
		for _, hit := range checkboxRe.FindAllStringSubmatch(string(src), -1) {
			m.total++
			if hit[1] != " " {
				m.done++
			}
		}
	}

	var done, total int
	for _, m := range modules {
		done += m.done
		total += m.total
	}

	if err := os.MkdirAll("progress", 0o755); err != nil {
		fatal("creating progress dir: %v", err)
	}

	for _, m := range modules {
		writeFile(filepath.Join("progress", m.id+".svg"), card(m.title, m.done, m.total, m.c1, m.c2))
	}
	writeFile(filepath.Join("progress", "overall.svg"), card("Overall — all modules", done, total, "#e3b341", "#d29922"))
	writeFile(filepath.Join("progress", "progress.csv"), csv(done, total))

	for _, m := range modules {
		updateBlock(m.file, "progress:"+m.id, moduleBlock(m))
	}
	updateBlock("README.md", "progress:overall", readmeBlock(done, total))

	fmt.Printf("%-40s %9s %6s\n", "module", "done", "")
	for _, m := range modules {
		fmt.Printf("%-40s %4d /%4d %5d%%\n", m.title, m.done, m.total, pct(m.done, m.total))
	}
	fmt.Printf("%-40s %4d /%4d %5d%%\n", "overall", done, total, pct(done, total))
}

func pct(done, total int) int {
	if total == 0 {
		return 0
	}
	return int(float64(done)/float64(total)*100 + 0.5)
}

// card renders one self-contained progress-bar card (dark, GitHub-friendly).
func card(title string, done, total int, c1, c2 string) string {
	const w, h, pad, barH = 640, 56, 20, 8
	barW := w - 2*pad
	p := pct(done, total)
	fill := barW * done
	if total > 0 {
		fill /= total
	} else {
		fill = 0
	}
	if done > 0 && fill < barH {
		fill = barH // keep the rounded cap visible for tiny progress
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" role="img" aria-label="%s: %d%%">`+"\n", w, h, w, h, esc(title), p)
	fmt.Fprintf(&b, `  <defs><linearGradient id="g" x1="0" y1="0" x2="1" y2="0"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>`+"\n", c1, c2)
	fmt.Fprintf(&b, `  <rect width="%d" height="%d" rx="12" fill="#0d1117" stroke="#30363d"/>`+"\n", w, h)
	fmt.Fprintf(&b, `  <text x="%d" y="24" font-family="-apple-system,'Segoe UI',Helvetica,Arial,sans-serif" font-size="13" font-weight="600" fill="#e6edf3">%s</text>`+"\n", pad, esc(title))
	fmt.Fprintf(&b, `  <text x="%d" y="24" text-anchor="end" font-family="-apple-system,'Segoe UI',Helvetica,Arial,sans-serif" font-size="12" fill="#9198a1">%d / %d · %d%%</text>`+"\n", w-pad, done, total, p)
	fmt.Fprintf(&b, `  <rect x="%d" y="36" width="%d" height="%d" rx="4" fill="#21262d"/>`+"\n", pad, barW, barH)
	if fill > 0 {
		fmt.Fprintf(&b, `  <rect x="%d" y="36" width="%d" height="%d" rx="4" fill="url(#g)"/>`+"\n", pad, fill, barH)
	}
	b.WriteString("</svg>\n")
	return b.String()
}

func moduleBlock(m *module) string {
	return fmt.Sprintf("![%s progress](progress/%s.svg)\n\n**%d / %d lessons complete · %d%%**",
		esc(m.title), m.id, m.done, m.total, pct(m.done, m.total))
}

func readmeBlock(done, total int) string {
	var b strings.Builder
	fmt.Fprintf(&b, "![Overall progress](progress/overall.svg)\n\n")
	for _, m := range modules {
		fmt.Fprintf(&b, "[![%s](progress/%s.svg)](%s)\n", esc(m.title), m.id, m.file)
	}
	fmt.Fprintf(&b, "\n**Total: %d / %d lessons complete · %d%%** — raw numbers in [progress/progress.csv](progress/progress.csv)", done, total, pct(done, total))
	return b.String()
}

func csv(done, total int) string {
	var b strings.Builder
	b.WriteString("module,title,done,total,percent\n")
	for _, m := range modules {
		fmt.Fprintf(&b, "%s,%q,%d,%d,%d\n", m.id, m.title, m.done, m.total, pct(m.done, m.total))
	}
	fmt.Fprintf(&b, "overall,\"All modules\",%d,%d,%d\n", done, total, pct(done, total))
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
