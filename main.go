package main

import (
	"github.com/tkmru/cibo/cibo"
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
	memSize := fileinfo.Size()
	beginAddress := 0x7c00
	emu := cibo.NewEmulator(beginAddress, memSize)
	RAM := emu.RAM
	cpu := emu.CPU
	reg := &cpu.X86registers
	f, _ := os.Open(filePath)
	copySize, _ := io.ReadFull(f, RAM)
	if int64(copySize) != fileinfo.Size() {
		log.Fatalln("size not matched")
	}

	reg.Init()
	emu.Run()
	log.Println("==================== registers ====================")
	reg.Dump()
}

func getPath() string {
	var arg string
	args := os.Args[1:]
	if len(args) == 1 {
		arg = args[0]
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
