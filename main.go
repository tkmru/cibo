package main

import (
	"github.com/tkmru/cibo/core"
	"os"
)

func main() {
	var path string
	args := os.Args[1:]
	if len(args) == 1 {
		path = args[0]
	}

	beginAddress := 0x7c00
	emu := cibo.NewEmulatorWithLoadFile(beginAddress, path)
	cpu := emu.CPU
	reg := &cpu.X86registers

	reg.Init()
	emu.Run()
	reg.Dump()
}
