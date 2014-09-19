package plan

import (
	"fmt"
	"github.com/brianm/gov/vcs"
	"log"
)

type Plan struct {
	TargetRepo     vcs.Repo
	DependentRepos []vcs.Repo
}

func CreatePlanFor(path string) (Plan, error) {
	plan := Plan{}

	tr, err := vcs.FindRepoForPath(path)
	if err != nil {
		return plan, fmt.Errorf("target not under source control: %s", err)
	}
	plan.TargetRepo = tr

	repos, err := vcs.FindRepos(path)
	if err != nil {
		return plan, fmt.Errorf("no repos for %s: %s", path, err)
	}

	plan.DependentRepos = make([]vcs.Repo, 0)
	for _, r := range repos {
		log.Printf("p=%s\tr=%s", path, r.Root())
		if tr.Root() != r.Root() {
			// exclude self repo
			plan.DependentRepos = append(plan.DependentRepos, r)
		}
	}

	return plan, nil
}