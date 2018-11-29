package cibo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func ioIn8(address uint16) uint8 {
	fmt.Println("[cibo] asking for input:")
	switch address {
	case 0x03c7: // Palette Address(Read Mode)
		break
	case 0x03c9: // Palette Data
		break
	case 0x03cc: // Miscellaneous Output Register on VGA
		break
	case 0x03f8: // COM1
		var input []byte = make([]byte, 1)
		os.Stdin.Read(input)
		return uint8(input[0])
		break
	}
	return 0
}

func ioIn32(address uint16) uint32 {
	fmt.Println("[cibo] asking for input:")
	switch address {
	case 0x03c7: // Palette Address(Read Mode)
		break
	case 0x03c9: // Palette Data
		break
	case 0x03cc: // Miscellaneous Output Register on VGA
		break
	case 0x03f8: // COM1
		var input []byte = make([]byte, 4)
		os.Stdin.Read(input)
		var i uint32
		buf := bytes.NewReader(input)
		binary.Read(buf, binary.LittleEndian, &i)
		return i
		break
	}
	return 0
}

func ioOut8(address uint16, ascii uint8) {
	switch address {
	case 0x03c2: // Miscellaneous Output Register on VGA
		break
	case 0x03c8: // Palette Address(Write Mode)
		break
	case 0x03c9: // Palette Data
		break
	case 0x03f8: // COM1
		fmt.Println(string(ascii))
		break
	}
}

func ioOut32(address uint16, ascii uint32) {
	switch address {
	case 0x03c2: // Miscellaneous Output Register on VGA
		break
	case 0x03c8: // Palette Address(Write Mode)
		break
	case 0x03c9: // Palette Data
		break
	case 0x03f8: // COM1
		fmt.Println(string(ascii))
		break
	}
}
