package main

import (
	"bitbucket.org/xnio/govdep2"
	"bitbucket.org/xnio/govdep2/sub2"
	"github.com/brianm/govdep1"
	"github.com/brianm/govdep1/sub1"
)

func main() {
	println(govdep1.Name())
	println(govdep2.Name())
	println(sub1.Name())
	println(sub2.Name())
}
