# cibo

**CURRENTLY UNDER DEVELOPMENT**

```cibo``` is a tool I've made to learn Golang and make somes experiments with CPU.

## Installation
The ```go get``` command compile the binary and place it in your ```$GOPATH/bin``` directory.

```
go get github.com/tkmru/cibo
```

## How to use
```cibo``` can be used from command.

```
$ ./cibo help
cibo - x86 CPU emulator

Usage:
  cibo [flags]
  cibo [command]

Available Commands:
  help        Help about any command
  version     Print the version number

Flags:
  -b, --bit int   bit mode (default 32)
      --debug     debug mode
  -h, --help      help for cibo

Use "cibo [command] --help" for more information about a command.
```

Also, ```cibo``` can be used api like unicorn.

```
package main

import (
	"github.com/tkmru/cibo/core"
)

func main() {
	beginAddress := 0x7c00
	emu := cibo.NewEmulator(32, beginAddress, 29, true)
	cpu := emu.CPU
	reg := &cpu.X86registers

	emu.RAM = []byte{0xb8, 0x01, 0x00, 0x00, 0x00, 0x3d, 0x02, 0x00, 0x00, 0x00, 0x75, 0x05, 0xe9, 0xef, 0x83, 0xff,
		0xff, 0xb8, 0x02, 0x00, 0x00, 0x00, 0x3d, 0x02, 0x00, 0x00, 0x00, 0x74, 0xef}
	/*
	    mov eax, 0x1
	    cmp eax, 0x2
	    jnz not_equal

	  equal:
	    jmp 0

	  not_equal:
	    mov eax, 0x2
	    cmp eax, 0x2
	    jz equal
	*/

	reg.Init()
	emu.Run()
}
```

## License
The MIT License
