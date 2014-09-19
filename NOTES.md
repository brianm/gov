# Project depends on its parent or child in repo

What is the correct behavior? It seems like it should not be allowed
to depend on a specific version of itself, so we should never include
the self repo in the list of dependencies.
