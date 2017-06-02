package main

import (
	"github.com/tkmru/cibo/cibo"
)

func main() {
	r := cibo.Registers{}
	r.Init()
	r.Dump()
}
