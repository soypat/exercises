package exercises

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"

	latex "github.com/soypat/goldmark-latex"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func TestSmoke(t *testing.T) {
	// Smoke-test
	exercises, err := ParseDirExercises("testdata/guia6")
	if err != nil {
		t.Fatal(err)
	}
	if len(exercises) == 0 {
		t.Fatal("zero exercises found in directory")
	}
}

func TestGenerateExerciseGuide(t *testing.T) {
	const dir = "testdata/guia6"
	hdfp, err := os.Open(dir + "/000-header.md")
	if err != nil {
		t.Fatal(err)
	}
	mdHeader, err := io.ReadAll(hdfp)
	hdfp.Close()
	if err != nil {
		t.Fatal(err)
	}
	// markdown will contain markdown file text.
	markdown := new(bytes.Buffer)
	_, err = markdown.Write(append(mdHeader, '\n'))
	if err != nil {
		t.Fatal(err)
	}
	exercises, err := ParseDirExercises(dir)
	if err != nil {
		t.Fatal(err)
	}
	for i, ex := range exercises {
		problem := strings.ReplaceAll(ex.Problem, "\n", "\n\t")
		problem = strings.ReplaceAll(problem, "  \n\n", "  \n")
		// fmt.Fprintf(markdown, "%d. %s\n\n", i+1, ex.Problem)
		fmt.Fprintf(markdown, "%d. %s\n\n", i+1, problem)
	}
	fmt.Fprintln(markdown, "# Solucionario")
	for i, ex := range exercises {
		if i%2 == 1 || true {
			fmt.Fprintf(markdown, "### Posible solución al %d.\n\n```python\n%s\n```\n\n", i+1, ex.SolutionCode)
		}
	}

	// We now work with a byte slice with markdown contents.
	// We will save it to a file and then use it to render latex.
	mdbytes := markdown.Bytes()
	mdOut, err := os.Create("testdata/guia6_test.md")
	if err != nil {
		t.Fatal(err)
	}
	defer mdOut.Close()
	_, err = io.Copy(mdOut, bytes.NewReader(mdbytes))
	if err != nil {
		t.Fatal(err)
	}

	// ltx will contain LaTeX text content.
	ltx := new(bytes.Buffer)
	rd := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(latex.NewRenderer(latex.Config{
		NoHeadingNumbering: true,                                                                     // No heading numbers
		Preamble:           append(latex.DefaultPreamble(), []byte("\n\\usepackage{MnSymbol}\n")...), // add star symbols to preamble.
		DeclareUnicode: func(r rune) (raw string, isReplaced bool) {
			switch r {
			case '★':
				return `$\filledstar$`, true
			case '☆':
				return `$\smallstar$`, true
			}
			return "", false
		},
	}), 1000)))
	md := goldmark.New(goldmark.WithRenderer(rd))
	err = md.Convert(mdbytes, ltx) // Latex rendering happens here.
	if err != nil {
		t.Fatal(err)
	}
	// Save latex contents, straight from the buffer.
	const latexFilename = "testdata/guia6_test.tex"
	os.Remove("testdata/guia6_test.pdf")
	os.Remove("guia6_test.pdf")
	ltxOut, err := os.Create(latexFilename)
	if err != nil {
		t.Fatal(err)
	}
	defer ltxOut.Close()
	_, err = io.Copy(ltxOut, ltx) // Write all Latex to file.
	if err != nil {
		t.Fatal(err)
	}
	err = exec.Command("latexmk", "-h").Run()
	if err != nil {
		t.Log("latexmk not installed")
		return
	}
	cmd := exec.Command("latexmk", "-pdf", latexFilename)
	cmd.Stdin = strings.NewReader("QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ") // Force batch mode
	cmdOut, err := cmd.CombinedOutput()
	if err != nil {
		t.Error(string(cmdOut))
		t.Fatal("running latexmk failed with error:", err)
	}
}
