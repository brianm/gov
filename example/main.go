package main

import (
	"github.com/brianm/govdep1"
	"github.com/brianm/govdep1/sub1"
	"github.com/brianm/govdep2"
	"github.com/brianm/govdep2/sub2"
)

func main() {
	println(govdep1.Name())
	println(govdep2.Name())
	println(sub1.Name())
	println(sub2.Name())
}
