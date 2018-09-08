package cibo

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Emulator struct {
	CPU         *CPU
	RAM         []byte
	BaseAddress int
	debugFlag bool
	// TODO: add device
}

func NewEmulator(beginAddress int, memSize int64, options ...interface{}) *Emulator {
	ram := make([]byte, memSize)
	debugFlag := false
	for _, option := range options {
    switch v := option.(type) {
			case bool:
        debugFlag = v
    }
	}
	emu := Emulator{nil, ram, beginAddress, debugFlag}
	emu.CPU = NewCPU(&emu)
	reg := &emu.CPU.X86registers
	reg.EIP = uint32(beginAddress)
	return &emu
}

func NewEmulatorWithLoadFile(beginAddress int, path string, options ...interface{}) *Emulator {
	log.SetFlags(0)

	filePath := checkPath(path)
	fileinfo, staterr := os.Stat(filePath)
	if staterr != nil {
		log.Fatalln(staterr)
	}
	memSize := fileinfo.Size()
	debugFlag := false
	for _, option := range options {
    switch v := option.(type) {
			case bool:
        debugFlag = v
    }
	}
	emu := NewEmulator(beginAddress, memSize, debugFlag)
	RAM := emu.RAM
	f, _ := os.Open(filePath)
	copySize, _ := io.ReadFull(f, RAM)
	if int64(copySize) != fileinfo.Size() {
		log.Fatalln("size not matched")
	}
	return emu
}

func (emu *Emulator) Run() {
	ramSize := len(emu.RAM)
	mappingEnd := emu.BaseAddress + ramSize
	cpu := emu.CPU
	mem := cpu.Memory
	reg := &cpu.X86registers
	if emu.debugFlag {
		for i := 0; i < int(ramSize); i++ {
			code := uint8(mem.GetCode8(0))
			log.Printf("EIP = 0x%X, Opcode = 0x%02X\n", reg.EIP, code)

			if cpu.InstTable[code] == nil {
				log.Fatalf("Not Implemented: 0x%x\n", code)
				break
			}

			cpu.InstTable[code]()

			if (reg.EIP <= uint32(emu.BaseAddress)) || (uint32(mappingEnd) <= reg.EIP) {
				log.Printf("No mapping area: 0x%X\n", reg.EIP)
				break
			}
		}
	} else {
		for i := 0; i < int(ramSize); i++ {
			code := uint8(mem.GetCode8(0))
			if cpu.InstTable[code] == nil {
				log.Fatalf("Not Implemented: 0x%x\n", code)
				break
			}

			cpu.InstTable[code]()

			if (reg.EIP <= uint32(emu.BaseAddress)) || (uint32(mappingEnd) <= reg.EIP) {
				log.Printf("No mapping area: 0x%X\n", reg.EIP)
				break
			}
		}
	}
}

func checkPath(filePath string) string {
	info, err := os.Stat(filePath)
	if err != nil {
		log.Fatalln("no binary file specified or found")
	}
	if info.IsDir() {
		files, err := ioutil.ReadDir(filePath)
		if err != nil {
			log.Fatalln("no binary file specified or found")
		}
		name := files[0].Name()
		return path.Join(filePath, name)

	} else {
		return filePath
	}
}
