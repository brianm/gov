# Gov #

Project and dependency management experiment for Go. After all, if the
problem is a lack of standard project automation tooling, clearly the
answer is yet another one!

# Design #

## Work With Go Get ##

Whether it is good or evil <code>go get</code> is a fact of Go life.
Play nicely with it!

## Support One GOPATH to Rule Them All ##

As the ecosystem expands, <code>go get</code> will start acting like
maven -- downloading the internet. The default repo caching behavior
is "one GOPATH to rule them all" which is damned handy, and the
default way Go likes to work. This means you have a local cache of
everything, you just lack control over it!

## Require Total Tree Versioning ##

If we are in a one-GOPATH world, and we want specific version of
dependencies, we should fix dependencies for the thing being built.
This means that we don't do recursive dependency version resolution,
but lock dependency version for a given top level package. If <code>A
-> B -> C</code> then A needs to specify the versions of both B and C.

## Pass Through Commands ##

The normal go commands should pass through, except with fixed
dependencies, when using gov. That is <code>gov build</code> should
fix dependency versions, then just run the exact same command but
<code>s/gov/go/</code> on the command line. We might need a facility
for gov-only flags, but they will then be --gov-foo or such, clearly
namespace.

## Be Strict ##

This is go after all, if a dependency is used but not specified,
break the build. This includes transitive dependencies. At the same
time, make it easy to fix the build. People will work ahead of their
frozen dependencies, so make it trivial to recur through the build and
find what has been added, then add it.

## Release is First Class ##

Releases are things. When <code>gov release</code> is run, it handles
tagging, merge to <code>go1</code> release 1.0.0 and later, and
creating a <code>source-package</code> which has the full closure of
dependencies so that we can easily build source debs or rpms which are
correct (all dependencies vendored).

Furthermore, we should support building rpm or deb bundles, probably
via a Makefile dropped in the the source package, so that the deb or
rpm is built exactly from the source bundle.

# Usage #

```
gov <task> [options]

tasks:
    sync     ensure correct version of all dependencies
    get      "go get ./..." + "gov sync"
    build    "gov get" + "go build"
    release  The Magic Starts Here
    check    Fast verification of dependency goodness
```
