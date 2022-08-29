package md2latex

import (
	"fmt"
	"io"

	"github.com/russross/blackfriday/v2"
)

type Flag int

const (
	NoFlags Flag = iota
	SkipHTML
)

var _ blackfriday.Renderer = &Renderer{}

type Renderer struct {
	Flags         Flag
	lastOutputLen int
}

func (r *Renderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	attrs := []string{}
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
	case blackfriday.HTMLSpan:
		// Skip HTML Span.
		break
		if r.Flags&SkipHTML != 0 {
			break
		}
		escapeLaTeX(w, node.Literal)
		// r.out(w, node.Literal)
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
				case HTMLBlock, List, Paragraph, Heading, CodeBlock, BlockQuote, HorizontalRule:
					r.cr(w)
				}
			}
			if node.Parent.Type == BlockQuote && node.Prev == nil {
				r.cr(w)
			}
			r.out(w, pTag)
		} else {
			r.out(w, pCloseTag)
			if !(node.Parent.Type == Item && node.Next == nil) {
				r.cr(w)
			}
		}
	case blackfriday.BlockQuote:
		if entering {
			r.cr(w)
			r.out(w, blockquoteTag)
		} else {
			r.out(w, blockquoteCloseTag)
			r.cr(w)
		}
	case blackfriday.HTMLBlock:
		if r.Flags&SkipHTML != 0 {
			break
		}
		r.cr(w)
		r.out(w, node.Literal)
		r.cr(w)
	case blackfriday.Heading:
		headingLevel := r.HTMLRendererParameters.HeadingLevelOffset + node.Level
		openTag, closeTag := headingTagsFromLevel(headingLevel)
		if entering {
			if node.IsTitleblock {
				attrs = append(attrs, `class="title"`)
			}
			if node.HeadingID != "" {
				id := r.ensureUniqueHeadingID(node.HeadingID)
				if r.HeadingIDPrefix != "" {
					id = r.HeadingIDPrefix + id
				}
				if r.HeadingIDSuffix != "" {
					id = id + r.HeadingIDSuffix
				}
				attrs = append(attrs, fmt.Sprintf(`id="%s"`, id))
			}
			r.cr(w)
			r.tag(w, openTag, attrs)
		} else {
			r.out(w, closeTag)
			if !(node.Parent.Type == Item && node.Next == nil) {
				r.cr(w)
			}
		}
	case blackfriday.HorizontalRule:
		r.cr(w)
		r.outHRTag(w)
		r.cr(w)
	case blackfriday.List:
		openTag := ulTag
		closeTag := ulCloseTag
		if node.ListFlags&ListTypeOrdered != 0 {
			openTag = olTag
			closeTag = olCloseTag
		}
		if node.ListFlags&ListTypeDefinition != 0 {
			openTag = dlTag
			closeTag = dlCloseTag
		}
		if entering {
			if node.IsFootnotesList {
				r.out(w, footnotesDivBytes)
				r.outHRTag(w)
				r.cr(w)
			}
			r.cr(w)
			if node.Parent.Type == Item && node.Parent.Parent.Tight {
				r.cr(w)
			}
			r.tag(w, openTag[:len(openTag)-1], attrs)
			r.cr(w)
		} else {
			r.out(w, closeTag)
			//cr(w)
			//if node.parent.Type != Item {
			//	cr(w)
			//}
			if node.Parent.Type == Item && node.Next != nil {
				r.cr(w)
			}
			if node.Parent.Type == Document || node.Parent.Type == BlockQuote {
				r.cr(w)
			}
			if node.IsFootnotesList {
				r.out(w, footnotesCloseDivBytes)
			}
		}
	case blackfriday.Item:
		openTag := liTag
		closeTag := liCloseTag
		if node.ListFlags&ListTypeDefinition != 0 {
			openTag = ddTag
			closeTag = ddCloseTag
		}
		if node.ListFlags&ListTypeTerm != 0 {
			openTag = dtTag
			closeTag = dtCloseTag
		}
		if entering {
			if itemOpenCR(node) {
				r.cr(w)
			}
			if node.ListData.RefLink != nil {
				slug := slugify(node.ListData.RefLink)
				r.out(w, footnoteItem(r.FootnoteAnchorPrefix, slug))
				break
			}
			r.out(w, openTag)
		} else {
			if node.ListData.RefLink != nil {
				slug := slugify(node.ListData.RefLink)
				if r.Flags&FootnoteReturnLinks != 0 {
					r.out(w, footnoteReturnLink(r.FootnoteAnchorPrefix, r.FootnoteReturnLinkContents, slug))
				}
			}
			r.out(w, closeTag)
			r.cr(w)
		}
	case blackfriday.CodeBlock:
		attrs = appendLanguageAttr(attrs, node.Info)
		r.cr(w)
		r.out(w, preTag)
		r.tag(w, codeTag[:len(codeTag)-1], attrs)
		escapeAllHTML(w, node.Literal)
		r.out(w, codeCloseTag)
		r.out(w, preCloseTag)
		if node.Parent.Type != Item {
			r.cr(w)
		}
	case blackfriday.Table:
		if entering {
			r.cr(w)
			r.out(w, tableTag)
		} else {
			r.out(w, tableCloseTag)
			r.cr(w)
		}
	case blackfriday.TableCell:
		openTag := tdTag
		closeTag := tdCloseTag
		if node.IsHeader {
			openTag = thTag
			closeTag = thCloseTag
		}
		if entering {
			align := cellAlignment(node.Align)
			if align != "" {
				attrs = append(attrs, fmt.Sprintf(`align="%s"`, align))
			}
			if node.Prev == nil {
				r.cr(w)
			}
			r.tag(w, openTag, attrs)
		} else {
			r.out(w, closeTag)
			r.cr(w)
		}
	case blackfriday.TableHead:
		if entering {
			r.cr(w)
			r.out(w, theadTag)
		} else {
			r.out(w, theadCloseTag)
			r.cr(w)
		}
	case blackfriday.TableBody:
		if entering {
			r.cr(w)
			r.out(w, tbodyTag)
			// XXX: this is to adhere to a rather silly test. Should fix test.
			if node.FirstChild == nil {
				r.cr(w)
			}
		} else {
			r.out(w, tbodyCloseTag)
			r.cr(w)
		}
	case blackfriday.TableRow:
		if entering {
			r.cr(w)
			r.out(w, trTag)
		} else {
			r.out(w, trCloseTag)
			r.cr(w)
		}
	// case blackfriday.Math:
	// 	r.out(w, mathTag)
	// 	escapeAllHTML(w, node.Literal)
	// 	r.out(w, mathCloseTag)
	default:
		panic("Unknown node type " + node.Type.String())
	}
	return blackfriday.GoToNext
}

func (r *Renderer) cr(w io.Writer) {
	w.Write(nlByte)
}

func (r *Renderer) out(w io.Writer, text []byte) {
	w.Write(text)
	r.lastOutputLen = len(text)
}

func (r *Renderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {

}

func (r *Renderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {

}
