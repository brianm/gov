package main

import (
	"fmt"
	"github.com/brianm/gov/vcs"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("please pass path on command line")
	}
	path := os.Args[1]
	rs, err := vcs.FindRepos(path)
	if err != nil {
		log.Fatalf("error finding repos: %s", err)
		panic(err)
	}

	for _, r := range rs {
		clean, err := r.IsClean()
		if err != nil {
			log.Println(err)
		}
		rev, err := r.Rev()
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("%s\t%t\t%v\n", rev, clean, r)
	}
}
