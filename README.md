<h1 align="center">fsdoublestar</h1>

<p  align="center">
 <a href="https://github.com/forensicanalysis/fsdoublestar/actions"><img src="https://github.com/forensicanalysis/fsdoublestar/workflows/CI/badge.svg" alt="build" /></a>
 <a href="https://codecov.io/gh/forensicanalysis/fsdoublestar"><img src="https://codecov.io/gh/forensicanalysis/fsdoublestar/branch/master/graph/badge.svg" alt="coverage" /></a>
 <a href="https://goreportcard.com/report/github.com/forensicanalysis/fsdoublestar"><img src="https://goreportcard.com/badge/github.com/forensicanalysis/fsdoublestar" alt="report" /></a>
 <a href="https://godocs.io/github.com/forensicanalysis/fsdoublestar"><img src="https://godocs.io/github.com/forensicanalysis/fsdoublestar?status.svg" alt="doc" /></a>
</p>

This package implements recursive directory globbing via `**` for Go's [io/fs](https://tip.golang.org/pkg/io/fs).

## Example

``` golang
package main

import (
	"fmt"
	"os"

	"github.com/forensicanalysis/fsdoublestar"
)

func main() {
	// get file system for this repository
	wd, _ := os.Getwd()
	fsys := os.DirFS(wd)

	// get all yml files
	matches, _ := fsdoublestar.Glob(fsys, "**/*.yml")

	// print matches
	fmt.Println(matches)
	// Output: [.github/workflows/ci.yml .github/.golangci.yml]
}
```

## Acknowledgement

This repository is based on [Bob Matcuk's](https://github.com/bmatcuk) great [doublestar](https://github.com/bmatcuk/doublestar) package.
