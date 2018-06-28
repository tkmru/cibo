package main

import (
	"./cibo"
	"fmt"
)

func main() {
	console, _ := cibo.NewConsole()
	cpu := console.CPU
	// r := cpu.X64registers
	r := cpu.X86registers
	r.Init()
	r.SetCF()
	if r.IsCF() {
		fmt.Printf("set\n")
	}
	r.RemoveCF()
	if !r.IsCF() {
		fmt.Printf("no set\n")
	}
	r.Dump()

}
