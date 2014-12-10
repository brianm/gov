package vcs

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
)

type Dependency interface {
	Root() string
	IsClean() bool
	Rev() string
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

func FindRepos(from string) ([]Dependency, error) {
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

	repos := make(map[string]Dependency)
	for _, val := range imports {
		if !val.Goroot {
			repo, err := FindRepo(val)
			if err != nil {
				return nil, err
			}
			repos[repo.Root()] = repo
		}
	}

	rs := make([]Dependency, 0, 0)
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

func FindRepoForPath(dir string) (Dependency, error) {
	dir = filepath.Clean(dir)
	if dir == "/" {
		return nil, fmt.Errorf("No repo found")
	}
	if isGit(dir) {
		return CreateGitRepo(dir)
	}
	parent, _ := filepath.Split(dir)
	return FindRepoForPath(parent)
}

func FindRepo(pkg *build.Package) (Dependency, error) {
	return FindRepoForPath(pkg.Dir)
}
