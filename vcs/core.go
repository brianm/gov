package vcs

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
)

type Repo interface {
	Root() string
	IsClean() (bool, error)
}

type GitRepo struct {
	root string
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

func findAllImports(pkg *build.Package, seen map[string]*build.Package) error {
	_, ok := seen[pkg.Dir]
	if ok {
		return nil
	}

	seen[pkg.Dir] = pkg
	for _, imp := range pkg.Imports {
		if imp == "C" {
			continue
		}
		cp, err := build.Import(imp, pkg.Dir, 0)
		if err != nil {
			return err
		}
		err = findAllImports(cp, seen)
		if err != nil {
			return err
		}
	}
	return nil
}

func FindRepos(from string) ([]Repo, error) {
	pkg, err := build.ImportDir(from, 0)
	if err != nil {
		return nil, fmt.Errorf("unable to load package at %s: %s", from, err)
	}

	imports := make(map[string]*build.Package)
	err = findAllImports(pkg, imports)
	if err != nil {
		return nil, err
	}
	delete(imports, pkg.Dir)

	repos := make(map[string]Repo)
	for _, val := range imports {
		if !val.Goroot {
			//fmt.Printf("pkg\t%s -> %v\n", val.ImportPath, val.Dir)
			repo, err := FindRepo(val)
			if err != nil {
				return nil, err
			}
			repos[repo.Root()] = repo
		}
	}

	rs := make([]Repo, 0, 0)
	for _, val := range repos {
		rs = append(rs, val)
	}

	return rs, nil
}

func isGit(dir string) bool {
	git_dir := fmt.Sprintf("%s/.git", dir)
	fi, err := os.Stat(git_dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if fi.IsDir() {
		return true
	}
	return false

}

func findRepoForPath(dir string) (Repo, error) {
	if filepath.Clean(dir) == "/" {
		return nil, fmt.Errorf("No repo found")
	}
	if isGit(dir) {
		return GitRepo{dir}, nil
	}
	parent, _ := filepath.Split(dir)
	return findRepoForPath(parent)
}

func FindRepo(pkg *build.Package) (Repo, error) {
	return findRepoForPath(pkg.Dir)
}
