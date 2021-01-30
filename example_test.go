package fsdoublestar_test

import (
	"fmt"
	"os"

	"github.com/forensicanalysis/fsdoublestar"
)

func Example() {
	// get file system for this repository
	wd, _ := os.Getwd()
	fsys := os.DirFS(wd)

	// get all yml files
	matches, _ := fsdoublestar.Glob(fsys, "**/*.yml")

	// print matches
	fmt.Println(matches)
	// Output: [.github/workflows/ci.yml .github/.golangci.yml]
}
