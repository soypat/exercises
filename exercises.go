package exercises

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Exercise struct {
	// Problem description that guides the user towards the solution.
	// Expected to be human readable text formatted in markdown.
	Problem string
	// An executable solution as put forward by exercise editor.
	SolutionCode string
}

type Filename struct {
	dir  string
	Num  int
	Name string
}

func ParseDirExercises(dir string) ([]Exercise, error) {
	found, err := ParseDir(dir)
	if err != nil {
		return nil, err
	}
	var exercises []Exercise
	for _, filename := range found {
		fp, err := os.Open(filename.Path())
		if err != nil {
			return nil, err
		}
		b, err := io.ReadAll(fp)
		if err != nil {
			return nil, err
		}
		exercise, err := ParsePythonExercise(string(b))
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

// ParseDir parses filenames and sorts according to number.
func ParseDir(dir string) ([]Filename, error) {
	if dir == "" {
		dir = "."
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var found []Filename
	for _, entry := range entries {
		filename := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(filename, ".py") {
			continue
		}
		path := filepath.Join(dir, filename)
		exercise, err := ParseFilename(path)
		if err != nil {
			continue
			return nil, errors.New("parsing directory \"" + dir + "\": " + err.Error())
		}
		found = append(found, exercise)
	}
	sort.Sort(ByFilenameNumber(found))

	return found, nil
}

type ByFilenameNumber []Filename

func (a ByFilenameNumber) Len() int           { return len(a) }
func (a ByFilenameNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFilenameNumber) Less(i, j int) bool { return a[i].Num < a[j].Num }

func (f Filename) Path() string {
	return filepath.Join(f.dir, fmt.Sprintf("%03d-%s", f.Num, f.Name))
}

func ParseFilename(path string) (Filename, error) {
	dir, filename := filepath.Split(path)
	if len(filename) < 8 {
		return Filename{}, errors.New(path + " exercise filename name must be more than 8 characters long")
	}
	n, err := strconv.Atoi(filename[0:3])
	if err != nil {
		return Filename{}, err
	}
	en := Filename{Num: n, Name: filename[4:], dir: dir}
	if en.Path() != path {
		return Filename{}, fmt.Errorf("generated path %q does not match argument path %q", en.Path(), path)
	}
	return en, nil
}

func ParsePythonExercise(text string) (Exercise, error) {
	if len(text) < 10 {
		return Exercise{}, errors.New("too short exercise")
	}

	const docstringSep = "\n\"\"\"\n"
	var afterDesc string
	if text[:4] == "\"\"\"\n" {
		afterDesc = text[4:]
	} else {
		_, afterDesc, _ = strings.Cut(text, docstringSep)
	}
	if afterDesc == "" {
		return Exercise{}, errors.New("no description found")
	}
	problem, solution, ok := strings.Cut(afterDesc, docstringSep)
	if !ok {
		return Exercise{}, errors.New("problem docstring not found")
	}
	return Exercise{Problem: problem, SolutionCode: strings.TrimSpace(solution)}, nil
}
