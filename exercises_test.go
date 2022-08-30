package exercises

import "testing"

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
