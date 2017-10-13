package main

import (
	"./cibo"
)

func main() {
	c := cibo.CPU{}
	r := c.Registers
	r.Init()
	r.Dump()
}
