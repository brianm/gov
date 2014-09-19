package main

import (
	"github.com/brianm/gov/example/child"
	"github.com/brianm/govdep1"
	"github.com/brianm/govdep1/sub1"
)

func main() {
	println(govdep1.Name())
	println(sub1.Name())
	println(child.Name())
}
