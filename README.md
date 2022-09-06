# exercises
Exercise management for PDF creation


## Demo
Builds "Guia 6 - Funciones" for IPC class @ UdeSA.

Requires [Go 1.19+](https://go.dev/doc/install) to build $\LaTeX$ `testdata/guia6_test.tex` file.

```shell
go test .
```
Optional: To also build to pdf have `latexmk` installed. 
If `latexmk` is not installed no pdf will be generated.