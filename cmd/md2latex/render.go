package main

import (
	"bytes"
	"errors"
	"flag"
	"io"
	"log"
	"main/md2latex"
	"os"

	"github.com/russross/blackfriday/v2"
)

func main() {
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
	renderer := &md2latex.Renderer{}
	latex := blackfriday.Run(input, blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.FencedCode))
	outfp, err := os.Create("output.tex")
	if err != nil {
		return err
	}
	defer outfp.Close()
	_, err = io.Copy(outfp, bytes.NewBuffer(latex))
	return err
}
