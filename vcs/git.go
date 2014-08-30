package vcs

import (
	"fmt"
	"os/exec"
)

type GitRepo struct {
	root string
}

func (g GitRepo) String() string {
	return fmt.Sprintf("git %s", g.Root())
}

func (g GitRepo) Root() string {
	return g.root
}

func (g GitRepo) IsClean() (bool, error) {
	// test -z "$(git status --porcelain)"
	git, err := exec.LookPath("git")
	if err != nil {
		return false, fmt.Errorf("unable to locate git utility: %s", err)
	}
	git_status := exec.Command(git, "status", "--porcelain")
	git_status.Dir = g.root
	out, err := git_status.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("Unable to run git status: %s", err)
	}
	return len(out) == 0, nil
}

func (g GitRepo) Rev() (string, error) {
	return "master", nil
}
