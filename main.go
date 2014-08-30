package main

import (
	"fmt"
	"github.com/brianm/gov/vcs"
	"log"
)

func main() {
	rs, err := vcs.FindRepos("./example")
	if err != nil {
		log.Fatalf("error finding repos: %s", err)
		panic(err)
	}

	for _, r := range rs {
		fmt.Printf("  %v\n", r)
	}
}
