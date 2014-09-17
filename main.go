package main

import (
	"fmt"
	"github.com/brianm/gov/vcs"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "gov"
	app.Usage = "gov <command>"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:      "report",
			ShortName: "r",
			Usage:     "Report on repos used",
			Action:    report,
		},
	}

	app.Run(os.Args)
}

func report(ctx *cli.Context) {
	var path string
	var err error

	if len(ctx.Args()) < 1 {
		path, err = os.Getwd()
		if err != nil {
			log.Fatalf("Unable to get current directory: %s", err)
		}
	} else {
		path = ctx.Args()[0]
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

		state := "dirty"
		if clean {
			state = "clean"
		}
		fmt.Printf("%s\t%s\t%v\n", rev, state, r)
	}
}
