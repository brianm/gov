// This test uses github.com/brianm/govdep* repos
// to be known quantities. If you much about with those
// repos it *will* break the test. Sorry.
package vcs

import (
	_ "github.com/brianm/govdep1"
	"go/build"
	"testing"
)

func TestFindRepos(t *testing.T) {
	example, err := build.Import("github.com/brianm/gov/example", ".", 0)
	if err != nil {
		t.Error(err)
	}

	repos, err := FindRepos(example.Dir)
	if len(repos) != 1 {
		t.Errorf("expected one dependency, found %d", len(repos))
		for _, r := range repos {
			t.Logf("found repo: %s", r)
		}
	}
}
