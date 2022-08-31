package exercises

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSmoke(t *testing.T) {
	// Smoke-test
	exercises, err := ParseDirExercises("testdata/guia6")
	if err != nil {
		t.Fatal(err)
	}
	if len(exercises) == 0 {
		t.Fatal("zero exercises found in directory")
	}
}

func TestGenerateExerciseGuide(t *testing.T) {
	const dir = "testdata/guia6"
	hdfp, err := os.Open(dir + "/000-header.md")
	if err != nil {
		t.Fatal(err)
	}
	mdHeader, err := io.ReadAll(hdfp)
	hdfp.Close()
	if err != nil {
		t.Fatal(err)
	}
	out, err := os.Create("testdata/guia6_test.md")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	_, err = out.Write(mdHeader)
	if err != nil {
		t.Fatal(err)
	}
	exercises, err := ParseDirExercises(dir)
	if err != nil {
		t.Fatal(err)
	}

	for i, ex := range exercises {
		fmt.Fprintf(out, "%d. %s\n\n", i, strings.ReplaceAll(ex.Problem, "\n", "\n\t"))
	}

	fmt.Fprintln(out, "# Solucionario (pares)")

	for i, ex := range exercises {
		if i%2 == 1 || true {
			fmt.Fprintf(out, "#### Solución al %d.\n```python\n%s\n```\n\n", i, ex.SolutionCode)
		}
	}
}
