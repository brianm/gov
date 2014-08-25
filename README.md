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
but lock dependency version for a given top level package. If <code>A -> B
-> C</code> then A needs to specify the versions of both B and C.

## Pass Through Commands ##

The normal go commands should pass through, except with fixed
dependencies, when using gov. That is <code>gov build</code> should
fix dependency versions, then just run the exact same command but
<code>s/gov/go/<code> on the command line. We might need a facility
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

# Working and Building #

Don't panic. We have a Makefile. The makefile manages the *project*
not the binary build. This way you can make a deb, rpm, man page, etc.

It *can* make a binary build, of course, but that is just a useful
thing, not general way of building.

If you want to work on this and have a workspace and everything set up
for you, my recommendation is to do this:

    $ make activate
    
This sets up the workspace, moves you into it in a subprocess, etc. To
stop working on the project just exit the sub-shell (C-d or
<code>exit</code>). 

You will need to modify both the Makefile in <code>master</code> and
<code>project</code> branches to use the proper package name and
binary name (the first two lines in the Makefile), but aside from that
things should Just Work.

# Versions, Release, Tags, and Branches 

We make releases, master is not always the best choice of branches to
use. When we make a release we'll merge the current release into the
<code>go1</code> branch. We'll also tag it, so you can get the exact
version you want.

If we make a backwards incompatible change, it will be a new project.
To put it differently, the <code>go1</code> branch will *always* be
backwards compatible.

# Working with this Project 

So, this project has (or will have shortly) stuff for buildings debs,
man pages, etc. This doesn't fit neatly into the <code>go
get</code>-able model (though it can). I also have beliefs about
workspace-per-project which are not widely held in the greater go
community. In the interest of accomodating everyone, the project is
<code>go get</code> friendly, but still has its own full workspace,
etc. How do we do this you ask? I am glad you asked!

You can checkout the <code>project</code> branch and run make to build
a workspace. You can then run <code>make activate</code> to enter the
workspace and do you stuff! It will drop you in a local checkout of
this project, inside the workspace.

Finally, as I don't hold with master always being perfectly stable,
there is a <code>go1</code> branch which is always the most current
release. <code>go get</code> prefers <code>go1</code> to
<code>master</code> so folks who want to rely on it as a library, or
install it via <code>go get</code> will get the most recent release,
rather than the most recent checkin to master.

If you hate all of this, ignore it. Pull requests, <code>go get</code>
etc all work as with any other Go project.
