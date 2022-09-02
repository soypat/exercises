package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/russross/blackfriday/v2"
	"github.com/soypat/exercises/md2latex"
)

var (
	usehtml bool
	print   bool
	unhead  bool
)

func main() {
	flag.BoolVar(&usehtml, "html", false, "output html")
	flag.BoolVar(&print, "p", false, "Output to stdout.")
	flag.BoolVar(&unhead, "unhead", false, "No heading numbering")
	flag.Parse()
	args := flag.Args()
	err := run(args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("missing filename argument")
	}
	fp, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer fp.Close()
	input, err := io.ReadAll(fp)
	if err != nil {
		return err
	}
	params := md2latex.RendererParameters{
		Flags:              md2latex.SkipHTML,
		NoHeadingNumbering: unhead,
	}
	var renderer blackfriday.Renderer
	if usehtml {
		renderer = blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{})
	} else {
		renderer = md2latex.NewRenderer(params)
	}

	output := blackfriday.Run(input, blackfriday.WithNoExtensions(), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.FencedCode))
	if print {
		fmt.Println(string(output))
		return nil
	}
	outfp, err := os.Create("output.tex")
	if err != nil {
		return err
	}
	defer outfp.Close()
	_, err = io.Copy(outfp, bytes.NewBuffer(output))
	return err
}
