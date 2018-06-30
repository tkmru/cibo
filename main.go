package main

import (
	"./cibo"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	log.SetFlags(0)

	binPath := getPaths()
	if len(binPath) == 0 {
		log.Fatalln("no binary file specified or found")
	}

	emu, _ := cibo.NewEmulator()
	cpu := emu.CPU
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


func getPaths() string {
	var arg string
	args := os.Args[1:]
	if len(args) == 1 {
		arg = args[0]
	} else {
		arg, _ = os.Getwd()
	}
	info, err := os.Stat(arg)
	if err != nil {
		return ""
	}
	if info.IsDir() {
		files, err := ioutil.ReadDir(arg)
		if err != nil {
			return ""
		}
		name := files[0].Name()
		return path.Join(arg, name)

	} else {
		return arg
	}
}
