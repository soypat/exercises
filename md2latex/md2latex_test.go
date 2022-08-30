package md2latex_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/russross/blackfriday/v2"
	"github.com/soypat/exercises/md2latex"
)

func TestRunGolden(t *testing.T) {
	fpin, err := os.Open("testdata/golden.md")
	if err != nil {
		t.Fatal(err)
	}
	markdown, err := io.ReadAll(fpin)
	if err != nil {
		t.Fatal(err)
	}
	renderer := md2latex.NewRenderer(md2latex.RendererParameters{Flags: md2latex.SkipHTML})
	latex := blackfriday.Run(markdown, blackfriday.WithRenderer(renderer))
	outfp, err := os.Create("testdata/output.tex")
	if err != nil {
		t.Fatal(err)
	}
	defer outfp.Close()
	_, err = io.Copy(outfp, bytes.NewBuffer(latex))
	if err != nil {
		t.Fatal(err)
	}
}
