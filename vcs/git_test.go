// This test uses github.com/brianm/govdep* repos
// to be known quantities. If you much about with those
// repos it *will* break the test. Sorry.
package vcs

import (
	_ "bitbucket.org/xnio/govdep2"
	_ "github.com/brianm/govdep1"
	"go/build"
	"os"
	"path/filepath"
	"testing"
)

var govdep1, govdep2 *build.Package

func init() {
	var err error
	govdep1, err = build.Import("github.com/brianm/govdep1", ".", 0)
	if err != nil {
		panic(err)
	}

	govdep2, err = build.Import("bitbucket.org/xnio/govdep2", ".", 0)
	if err != nil {
		panic(err)
	}
}

func TestSetup(t *testing.T) {
	if govdep1 == nil {
		t.Errorf("govdep1 is nil")
	}
	if govdep2 == nil {
		t.Errorf("govdep2 is nil")
	}
}

func TestGovDep1IsClean(t *testing.T) {
	g := GitRepo{govdep1.Dir}
	clean, err := g.IsClean()
	if err != nil {
		t.Errorf("Unable to look for clean state of %s: %s", govdep1.Dir, err)
	}
	if !clean {
		t.Fatalf("%s should have been clean!", govdep1.Dir)
	}
}

func TestMakeGovDep1Dirty(t *testing.T) {
	tmp := filepath.Join(govdep1.Dir, "TestMakeGovDep1Dirty")
	_, err := os.Create(tmp)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmp)

	g := GitRepo{govdep1.Dir}
	clean, err := g.IsClean()
	if err != nil {
		t.Errorf("Unable to look for clean state of %s: %s", govdep1.Dir, err)
	}
	if clean {
		t.Fatalf("%s should have been dirty!", govdep1.Dir)
	}
}
