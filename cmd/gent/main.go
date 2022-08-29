package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dir := flag.String("d", ".", "Exercise directory.")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage("need command")
		os.Exit(1)
	}

	switch args[0] {
	case "insert":
		if len(args) != 2 {
			usage("need filename argument for insert")
			os.Exit(1)
		}
		commandInsert(*dir, args[1])
	case "collapse":
		commandCollapse(*dir)
	case "list":
		list(*dir)
	default:
		usage("unknown command " + args[0])
		os.Exit(1)
	}
	log.Println("finished succesfully")
}

func list(dir string) {
	exercises, err := parseDirExercises(dir)
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(byNumber(exercises))
	for i, ex := range exercises {
		fmt.Printf("%d:\t%v\n", i, ex.filename())
	}
}

func commandCollapse(dir string) {
	exercises, err := parseDirExercises(dir)
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(byNumber(exercises))
	for i, ex := range exercises {
		expectedNum := i + 1
		if ex.num == expectedNum {
			continue
		}
		renamed := e{num: i + 1, name: ex.name}
		if _, err := os.Stat(renamed.filename()); err == nil {
			log.Fatalf("can't rename %q to already existing %q", ex.filename(), renamed.filename())
		}
		err := os.Rename(ex.filename(), renamed.filename())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("collapse: renaming %q -> %q\n", ex.filename(), renamed.filename())
	}
}

func commandInsert(dir, name string) {
	newExercise, err := parseExerciseFilename(name)
	if err != nil {
		log.Fatal("parsing argument exercise name: ", err)
	}
	if _, err := os.Stat(name); err == nil {
		log.Fatalf("exercise %s already exists", name)
	}
	log.Printf("inserting exercise %q", newExercise.filename())
	oldExercises, err := parseDirExercises(dir)
	if err != nil {
		log.Fatal(err)
	}
	for i := range oldExercises {
		if oldExercises[i].name == newExercise.name {
			log.Fatalf("exercise %q already exists", newExercise.name)
		}
	}
	// Modify old exercise names and add new exercise.
	for _, exercise := range oldExercises {
		if exercise.num < newExercise.num {
			continue // Look for conflicting exercise.
		}
		if exercise.num != newExercise.num {
			break // No need to rename exercise, no conflict found.
		}
		renamedExercise := e{
			num:  exercise.num + 1.,
			name: exercise.name,
		}
		err := os.Rename(exercise.filename(), renamedExercise.filename())
		if err != nil {
			log.Fatal("")
		}
	}
	fp, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	fp.Close()
}

var excluded = []string{}

func isExcluded(s string) bool {
	for _, ex := range excluded {
		if ex == s {
			return true
		}
	}
	return false
}

func parseDirExercises(dir string) ([]e, error) {
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}
	var found []e
	for _, entry := range entries {
		filename := entry.Name()
		if entry.IsDir() || isExcluded(filename) {
			continue
		}
		exercise, err := parseExerciseFilename(filename)
		if err != nil {
			return nil, errors.New("parsing existing exercise directory: " + err.Error())
		}
		newName := (e).filename(e{num: exercise.num + 1, name: exercise.name})
		_, err = os.Stat(newName)
		if err == nil {
			fmt.Errorf("conflict rename: exercise rename %q already exists", newName)
		}
		found = append(found, exercise)
	}
	return found, nil
}

type e struct {
	num  int
	name string
}

func (en e) filename() string {
	return fmt.Sprintf("%03d-%s", en.num, en.name)
}

func parseExerciseFilename(filename string) (e, error) {
	if len(filename) < 8 {
		return e{}, errors.New(filename + " exercise filename name must be more than 8 characters long")
	}
	if !strings.HasSuffix(filename, ".py") {
		return e{}, errors.New(filename + " is not a python file")
	}
	n, err := strconv.Atoi(filename[0:3])
	if err != nil {
		return e{}, err
	}
	en := e{num: n, name: filename[4:]}
	if en.filename() != filename {
		return e{}, fmt.Errorf("generated directory name %q does not match argument name %q", en.filename(), filename)
	}
	return en, nil
}

func usage(msg string) {
	fmt.Fprintf(os.Stderr, "grand exercise numeration tool (gent).\n")
	fmt.Fprintf(os.Stderr, "usage:\n\tgent <command> [argument]\n")
	fmt.Fprintf(os.Stderr, "Supported commands:\n")
	fmt.Fprintf(os.Stderr, "\tinsert [name of exercise directory]\n")
	fmt.Fprintf(os.Stderr, "inserts an exercise directory and renames existing exercise numbers to avoid number collisions\n")
	fmt.Fprintf(os.Stderr, "\tcollapse\n")
	fmt.Fprintf(os.Stderr, "renames all exercise numbers so that exercise numbers increase by 1\n")
	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
}

type byNumber []e

func (a byNumber) Len() int           { return len(a) }
func (a byNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byNumber) Less(i, j int) bool { return a[i].num < a[j].num }
