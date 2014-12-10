package vcs

import (
	"fmt"
	"os/exec"
	"strings"
)

type gitRepo struct {
	root  string
	clean bool
	rev   string
}

func CreateGitRepo(path string) (Repo, error) {
	r := gitRepo{}

	// IsClean
	// test -z "$(git status --porcelain)"
	git, err := exec.LookPath("git")
	if err != nil {
		return r, fmt.Errorf("unable to locate git utility: %s", err)
	}
	git_status := exec.Command(git, "status", "--porcelain")
	git_status.Dir = path
	out, err := git_status.CombinedOutput()
	if err != nil {
		return r, fmt.Errorf("Unable to run git status: %s", err)
	}
	r.clean = (len(out) == 0)

	// rev

	find_rev_sh := "git symbolic-ref -q --short HEAD || git describe --tags"
	rev := exec.Command("/bin/sh", "-c", find_rev_sh)
	rev.Dir = path
	out, err = rev.CombinedOutput()
	if err != nil {
		return r, fmt.Errorf("Unable to find branch or tag: %s", err)
	}
	r.rev = strings.TrimSpace(string(out))

	r.root = path

	return r, nil
}

func (g gitRepo) String() string {
	return fmt.Sprintf("git %s", g.Root())
}

func (g gitRepo) Root() string {
	return g.root
}

func (g gitRepo) IsClean() bool {
	return g.clean
}

func (g gitRepo) Rev() string {
	return g.rev
}
