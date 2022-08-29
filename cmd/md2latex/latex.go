//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
//
// LaTeX rendering backend
//
//

package main

import (
	"bytes"

	"github.com/russross/blackfriday"
)

// Override overrides certain Renderer methods.
type Override struct {
	blackfriday.Renderer
}

func (options *Override) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	marker := out.Len()

	switch level {
	case 1:
		out.WriteString("\n\\section*{")
	case 2:
		out.WriteString("\n\\subsection*{")
	case 3:
		out.WriteString("\n\\subsubsection*{")
	case 4:
		out.WriteString("\n\\paragraph{")
	case 5:
		out.WriteString("\n\\subparagraph{")
	case 6:
		out.WriteString("\n\\textbf{")
	}
	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString("}\n")
}
