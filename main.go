package main

import (
	"./cibo"
)

func main() {
	console, _ := cibo.NewConsole()
	cpu := console.CPU
	r := cpu.Registers
	r.Init()
	r.Dump()
}
