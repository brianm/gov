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
	app.Usage = "manage golang build dependencies"
	app.EnableBashCompletion = true
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "report",
			ShortName: "r",
			Usage:     "Report on repos used",
			Action:    report,
		},
		{
			Name:   "bash-autocomplete",
			Usage:  "Use as 'eval \"$(gov bash-autocomplete)\"' to set up bash autocompletion",
			Action: autocomplete,
		},
	}

	app.Run(os.Args)
}

func autocomplete(*cli.Context) {
	fmt.Printf("%s", `
_gov_bash_autocomplete() {
     local cur prev opts base
     COMPREPLY=()
     cur="${COMP_WORDS[COMP_CWORD]}"
     prev="${COMP_WORDS[COMP_CWORD-1]}"
     opts=$( ${COMP_WORDS[@]:0:COMP_CWORD} --generate-bash-completion )
     # add -f to compgen to get filenames as well
     COMPREPLY=( $(compgen -f -W "${opts}" -- ${cur}) )
     return 0
} 
complete -F _gov_bash_autocomplete gov

`)
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
		clean := r.IsClean()
		rev := r.Rev()

		state := "dirty"
		if clean {
			state = "clean"
		}
		fmt.Printf("%s\t%s\t%v\n", rev, state, r)
	}
}
