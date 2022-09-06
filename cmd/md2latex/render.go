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
	latex "github.com/soypat/goldmark-latex"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var (
	usehtml      bool
	print        bool
	unhead       bool
	rendererFlag string
)

func main() {
	flag.BoolVar(&usehtml, "html", false, "Output html")
	flag.BoolVar(&print, "p", false, "Output to stdout")
	flag.BoolVar(&unhead, "unhead", false, "No heading numbering")
	flag.StringVar(&rendererFlag, "renderer", "goldmark", "Renderer used. Available: "+availableRenders())
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
	rd := renderers[rendererFlag]
	if rd == nil {
		return errors.New("renderer \"" + rendererFlag + "\" undefined")
	}
	output, err := rd(input)
	if err != nil {
		return err
	}
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

var renderers = map[string]func(input []byte) ([]byte, error){
	"blackfriday": bf,
	"goldmark":    gm,
}

func availableRenders() (out string) {
	out = "["
	for k := range renderers {
		out += `"` + k + `",`
	}
	return out[:len(out)-2] + "]"
}

func bf(input []byte) ([]byte, error) {
	params := md2latex.RendererParameters{
		Flags:              md2latex.SkipHTML,
		NoHeadingNumbering: unhead,
	}
	var rd blackfriday.Renderer
	if usehtml {
		rd = blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{})
	} else {
		rd = md2latex.NewRenderer(params)
	}
	output := blackfriday.Run(input, blackfriday.WithNoExtensions(), blackfriday.WithRenderer(rd), blackfriday.WithExtensions(blackfriday.FencedCode))
	return output, nil
}

func gm(input []byte) ([]byte, error) {
	var rd renderer.Renderer
	if usehtml {
		rd = goldmark.DefaultRenderer()
	} else {
		rd = renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(latex.NewRenderer(latex.Config{
			NoHeadingNumbering: unhead,
		}), 1000)))
	}
	md := goldmark.New(goldmark.WithRenderer(rd))
	var b bytes.Buffer
	err := md.Convert(input, &b)
	return b.Bytes(), err
}
