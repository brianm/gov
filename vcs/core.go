package vcs

import (
	"fmt"
	"go/build"
	"os"
)

type Repo struct {
	Root string
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
			repo, err := findRepo(val)
			if err != nil {
				return nil, err
			}
			repos[repo.Root] = repo
		}
	}

	rs := make([]Repo, 0, 0)
	for _, val := range repos {
		rs = append(rs, val)
	}

	return rs, nil
}

type VcsType int

const (
	None VcsType = iota
	Git
	Unknown
)

func (v VcsType) String() string {
	switch v {
	case None:
		return "None"
	case Git:
		return "Git"
	case Unknown:
		return "Unknown"
	default:
		return "really unknown!"
	}
}

func findVcsType(dir string) (VcsType, error) {

	ctx := build.Default

	if ctx.IsDir(fmt.Sprintf("%s/.git")) {

	}
	// git test
	git_dir := fmt.Sprintf("%s/.git", dir)
	fmt.Printf("testing %s\n", git_dir)
	fi, err := os.Stat(git_dir)
	if err != nil {
		return Unknown, err
	}
	if fi.IsDir() {
		return Git, nil
	}

	// recur!

	return Unknown, nil
}

func findRepo(pkg *build.Package) (Repo, error) {

	dir := pkg.Dir

	vctype, err := findVcsType(dir)
	if err != nil {
		return Repo{dir}, err
	}
	fmt.Printf("%v\n", vctype)

	//if strings.HasPrefix(pkg.ImportPath, "github.com/brianm/a") {
	// special case right now until I get vcs detection
	// this just allows the multiple import de-dupe to be
	// tested
	//return Repo{"github.com/brianm/a"}, nil
	//}
	return Repo{dir}, nil
}
