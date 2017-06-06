package main

import (
	"github.com/tkmru/cibo/cibo"
)

func main() {
	c := cibo.CPU{}
	r := c.Registers
	r.Init()
	r.Dump()
}
