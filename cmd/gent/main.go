package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/soypat/exercises"
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
	exes, err := exercises.ParseDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for i, ex := range exes {
		fmt.Printf("%d:\t%v\n", i, ex.Path())
	}
}

func commandCollapse(dir string) {
	exes, err := exercises.ParseDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for i, ex := range exes {
		expectedNum := i + 1
		if ex.Num == expectedNum {
			continue
		}
		renamed := exercises.Filename{Num: i + 1, Name: ex.Name}
		if _, err := os.Stat(renamed.Path()); err == nil {
			log.Fatalf("can't rename %q to already existing %q", ex.Path(), renamed.Path())
		}
		err := os.Rename(ex.Path(), renamed.Path())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("collapse: renaming %q -> %q\n", ex.Path(), renamed.Path())
	}
}

func commandInsert(dir, name string) {
	newExercise, err := exercises.ParseFilename(filepath.Join(dir, name))
	if err != nil {
		log.Fatal("parsing argument exercise name: ", err)
	}
	if _, err := os.Stat(newExercise.Path()); err == nil {
		log.Fatalf("exercise %s already exists", name)
	}
	log.Printf("inserting exercise %q", newExercise.Path())
	oldExercises, err := exercises.ParseDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for i := range oldExercises {
		if oldExercises[i].Name == newExercise.Name {
			log.Fatalf("exercise %q already exists", newExercise.Name)
		}
	}
	// Modify old exercise names and add new exercise.
	for _, exercise := range oldExercises {
		if exercise.Num < newExercise.Num {
			continue // Look for conflicting exercise.
		}
		if exercise.Num != newExercise.Num {
			break // No need to rename exercise, no conflict found.
		}
		renamedExercise := exercises.Filename{
			Num:  exercise.Num + 1.,
			Name: exercise.Name,
		}
		err := os.Rename(exercise.Path(), renamedExercise.Path())
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
