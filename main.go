package main

import (
	"./cibo"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	log.SetFlags(0)

	filePath := getPath()
	if len(filePath) == 0 {
		log.Fatalln("no binary file specified or found")
	}
	fileinfo, staterr := os.Stat(filePath)
	if staterr != nil {
		log.Fatalln(staterr)
  }
	emu := cibo.NewEmulator(0x7c00, fileinfo.Size())
	RAM := emu.RAM
	f, _ := os.Open(filePath)
	copySize, _ := io.ReadFull(f, RAM)
	if int64(copySize) != fileinfo.Size() {
		log.Fatalln("size not matched")
	}

	cpu := emu.CPU
	reg := cpu.X86registers
	reg.Init()
	reg.Dump()
}


func getPath() string {
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
