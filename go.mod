module github.com/soypat/exercises

go 1.19
replace github.com/soypat/goldmark-latex => ../goldmark-latex
replace github.com/russross/blackfriday/v2 => ../blackfriday

require github.com/russross/blackfriday/v2 v2.1.0

require (
	github.com/soypat/goldmark-latex v0.1.0 // indirect
	github.com/yuin/goldmark v1.4.14 // indirect
)
