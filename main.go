package main

import (
	"fmt"
	"github.com/brianm/gov/vcs"
)

func main() {
	rs, err := vcs.FindRepos("./example")
	if err != nil {
		panic(err)
	}

	for _, r := range rs {
		fmt.Printf("  %v\n", r)
	}
}
