package md2latex

import (
	_ "embed"
	"fmt"
	"io"

	"github.com/russross/blackfriday/v2"
	latex "github.com/soypat/goldmark-latex"
)

type Flags int

const (
	NoFlags Flags = iota
	SkipHTML
)

func NewRenderer(p RendererParameters) blackfriday.Renderer {
	r := &renderer{
		RendererParameters: p,
		labelIDs:           make(map[string]int),
	}
	return r
}

type renderer struct {
	RendererParameters
	lastOutputLen int
	// Track heading IDs to prevent ID collision in a single generation.
	labelIDs map[string]int
}

type RendererParameters struct {
	// Flags allow customizing this renderer's behavior
	Flags Flags
	// Increase heading levels: if the offset is 1, \section (1) becomes \subsection (2) etc.
	// Negative offset is also valid.
	// Resulting levels are clipped between 1 and 6.
	HeadingLevelOffset int
	// Removes section numbering.
	NoHeadingNumbering bool
	// Replace the default preamble by setting this to a non-nil byte slice.
	Preamble []byte
}

var defaultHeader = latex.DefaultPreamble()

func (r *renderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {
	if r.Preamble == nil {
		w.Write(defaultHeader)
	} else {
		w.Write(r.Preamble)
	}
	w.Write([]byte("\n\\begin{document}\n"))
}

func (r *renderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {
	w.Write([]byte("\n\\end{document}\n"))
}

func (r *renderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	switch node.Type {
	case blackfriday.Text:
		if node.Parent.Type == blackfriday.Link {
			escLink(w, node.Literal)
		} else {
			escapeLaTeX(w, node.Literal)
		}
	case blackfriday.Softbreak:
		r.cr(w)
		// TODO: make it configurable via out(renderer.softbreak)
	case blackfriday.Hardbreak:
		r.out(w, hardBreak)
		r.cr(w)
	case blackfriday.Emph:
		if entering {
			r.out(w, italicStart)
		} else {
			r.out(w, braceEnd)
		}
	case blackfriday.Strong:
		if entering {
			r.out(w, boldStart)
		} else {
			r.out(w, braceEnd)
		}
	case blackfriday.Del:
		if entering {
			r.out(w, strikeStart)
		} else {
			r.out(w, braceEnd)
		}
	case blackfriday.Link:
		if entering {
			r.out(w, hrefStart)
			escLink(w, node.LinkData.Destination)
			r.out(w, braceEnd)
			r.out(w, braceStart)
		} else {
			r.out(w, braceEnd)
		}
	case blackfriday.Image:
		// Skip images.
	case blackfriday.Code:
		r.out(w, codeStart)
		escapeLaTeX(w, node.Literal)
		r.out(w, braceEnd)
	case blackfriday.Document:
		break
	case blackfriday.Paragraph:
		if skipParagraphTags(node) {
			break
		}
		if entering {
			// TODO: untangle this clusterfuck about when the newlines need
			// to be added and when not.
			if node.Prev != nil {
				switch node.Prev.Type {
				case blackfriday.HTMLBlock, blackfriday.List, blackfriday.Paragraph, blackfriday.Heading, blackfriday.CodeBlock, blackfriday.BlockQuote, blackfriday.HorizontalRule:
					r.cr(w)
				}
			}
			if node.Parent.Type == blackfriday.BlockQuote && node.Prev == nil {
				r.cr(w)
			}
			// r.out(w, pTag)
		} else if node.Parent.Type != blackfriday.BlockQuote && node.Parent.Type != blackfriday.Item {
			r.out(w, hardBreak)
			if !(node.Parent.Type == blackfriday.Item && node.Next == nil) {
				r.cr(w)
			}
		}
	case blackfriday.BlockQuote:
		if entering {
			r.cr(w)
			r.out(w, blockQuoteStart)
		} else {
			r.out(w, blockQuoteEnd)
			r.cr(w)
		}
	case blackfriday.HTMLBlock, blackfriday.HTMLSpan:
		if r.Flags&SkipHTML != 0 {
			break
		}
		panic("md2latex: HTML rendering not supported")

	case blackfriday.Heading:
		if entering {
			headingLevel := r.RendererParameters.HeadingLevelOffset + node.Level
			start := headingFromLevel(headingLevel, r.NoHeadingNumbering)
			r.cr(w)
			r.out(w, start)
			if headingLevel >= 5 {
				r.cr(w)
			}
		} else {
			r.out(w, braceEnd)
			if node.HeadingID != "" {
				_, exist := r.labelIDs[node.HeadingID]
				if exist {
					panic("repeated HeadingID: " + node.HeadingID)
				}
				r.labelIDs[node.HeadingID] = 1
				fmt.Fprintf(w, "\\label{%s}\n", node.HeadingID)
			}
			if !(node.Parent.Type == blackfriday.Item && node.Next == nil) {
				r.cr(w)
			}
		}
	case blackfriday.HorizontalRule:
		r.cr(w)
		r.out(w, hruleCommand)
		r.cr(w)
	case blackfriday.List:
		opener := ulStart
		closer := ulEnd
		if node.ListFlags&blackfriday.ListTypeOrdered != 0 {
			opener = olStart
			closer = olEnd
		}
		if node.ListFlags&blackfriday.ListTypeDefinition != 0 {
			opener = descriptionStart
			closer = descriptionEnd
		}
		if entering {
			r.cr(w)
			if node.Parent.Type == blackfriday.Item && node.Parent.Parent.Tight {
				r.cr(w)
			}
			r.out(w, opener)
			r.cr(w)
		} else {
			r.out(w, closer)
			if node.Parent.Type == blackfriday.Item && node.Next != nil {
				r.cr(w)
			}
			if node.Parent.Type == blackfriday.Document || node.Parent.Type == blackfriday.BlockQuote {
				r.cr(w)
			}
		}
	case blackfriday.Item:
		if entering {
			r.out(w, itemCommand)
		} else {
			r.cr(w)
		}
	case blackfriday.CodeBlock:
		// r.out(w, parBreak)
		fmt.Fprintf(w, "\n%%s\n\\begin{lstlisting}[language=%s]\n%s\n\\end{lstlisting}", node.Info, string(node.Literal))
		if node.Parent.Type != blackfriday.Item {
			r.cr(w)
		}
	case blackfriday.Table:
		if entering {
			r.cr(w)
			r.out(w, tableStart)
			r.out(w, []byte("\n% Tables unsupported!\n"))
		} else {
			r.out(w, tableEnd)
			r.cr(w)
		}
	default:
		panic("Unknown node type " + node.Type.String())
	}
	return blackfriday.GoToNext
}

func (r *renderer) cr(w io.Writer) {
	w.Write(nlByte)
}

func (r *renderer) out(w io.Writer, text []byte) {
	w.Write(text)
	r.lastOutputLen = len(text)
}

func skipParagraphTags(node *blackfriday.Node) bool {
	grandparent := node.Parent.Parent
	if grandparent == nil || grandparent.Type != blackfriday.List {
		return false
	}
	tightOrTerm := grandparent.Tight || node.Parent.ListFlags&blackfriday.ListTypeTerm != 0
	return grandparent.Type == blackfriday.List && tightOrTerm
}

func headingFromLevel(level int, nonumber bool) []byte {
	if level < 1 {
		level = 1
	} else if level > 6 {
		level = 6
	}
	return headingTable[level-1][bool2int(nonumber)]
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
