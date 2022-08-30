package md2latex

import (
	"io"
)

var (
	nlByte           = []byte{'\n'}
	hardBreak        = []byte("\\\\")
	braceStart       = []byte{'{'}
	braceEnd         = []byte{'}'}
	boldStart        = []byte("\\textbf{")
	italicStart      = []byte("\\textit{")
	strikeStart      = []byte("\\sout{") // Using ulem package.
	hrefStart        = []byte("\\href{")
	codeStart        = []byte("\\texttt{")
	blockQuoteStart  = []byte("\\begin{lstlisting}[frame=none]")
	blockQuoteEnd    = []byte("\\end{lstlisting}")
	hruleCommand     = []byte("\\hrulefill")
	ulStart          = []byte("\\begin{itemize}")
	ulEnd            = []byte("\\end{itemize}")
	olStart          = []byte("\\begin{enumerate}")
	olEnd            = []byte("\\end{enumerate}")
	descriptionStart = []byte("\\begin{description}")
	descriptionEnd   = []byte("\\end{description}")
	itemCommand      = []byte("\\item~")
	tableStart       = []byte("\\begin{table}")
	tableEnd         = []byte("\\end{table}")
	headerTable      = [6][2][]byte{
		{[]byte("\\section{"), []byte("\\section*{")},
		{[]byte("\\subsection{"), []byte("\\subsection*{")},
		{[]byte("\\subsubsection{"), []byte("\\subsubsection*{")},
		{[]byte("\\paragraph{"), []byte("\\paragraph*{")},
		{[]byte("\\subparagraph{"), []byte("\\subparagraph*{")},
		{[]byte("\\textbf{"), []byte("\\textbf{")},
	}
)

var escapeTable = [256][]byte{
	'\\': []byte("\\textbackslash~"),
	'~':  []byte("\\textasciitilde~"),
	'^':  []byte("\\textasciicircum~"),
	'&':  []byte("\\&"),
	'%':  []byte("\\%"),
	'$':  []byte("\\$"),
	'#':  []byte("\\#"),
	'_':  []byte("\\_"),
	'{':  []byte("\\{"),
	'}':  []byte("\\}"),
}

func escapeLaTeX(w io.Writer, s []byte) {
	var start, end int
	for end < len(s) {
		escSeq := escapeTable[s[end]]
		if escSeq != nil {
			w.Write(s[start:end])
			w.Write(escSeq)
			start = end + 1
		}
		end++
	}
	if start < len(s) && end <= len(s) {
		w.Write(s[start:end])
	}
}

func escLink(w io.Writer, text []byte) {
	escapeLaTeX(w, text)
}
