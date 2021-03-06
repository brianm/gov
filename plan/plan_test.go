package plan

import (
	"go/build"
	"path/filepath"
	"testing"
)

var example *build.Package

func init() {
	var err error
	example, err = build.Import("github.com/brianm/gov/example", ".", 0)
	if err != nil {
		panic(err)
	}
}

func TestCreatePlanExcludesSelfRepo(t *testing.T) {
	plan, err := CreatePlanFor(example.Dir)
	if err != nil {
		t.Fatal(err)
	}

	repos := plan.DependentRepos
	if len(repos) != 1 {
		t.Errorf("expected one dependency, found %d", len(repos))
		for _, r := range repos {
			t.Logf("found repo: %s", r)
		}
	}
}

func TestTargetRepoFound(t *testing.T) {
	plan, err := CreatePlanFor(example.Dir)
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Join(plan.TargetRepo.Root(), "example") != example.Dir {
		t.Fatalf("example repo not found correctly!")
	}
}
