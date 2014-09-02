package main

import (
	"flag"
	"fmt"
	"github.com/brianm/gov/vcs"
	"log"
	"os"
)

func main() {
	var path string
	var err error

	flag.Parse()

	if len(flag.Args()) < 1 {
		path, err = os.Getwd()
		if err != nil {
			log.Fatalf("Unable to get current directory: %s", err)
		}
	} else {
		path = flag.Arg(0)
	}

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
