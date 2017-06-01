package main

import (
	"fmt"
	"reflect"
)

type Registers struct {
	// GPR
	EAX uint32
	EBX uint32
	ECX uint32
	EDX uint32
	ESP uint32
	EBP uint32
	ESI uint32
	EDI uint32
	// Instruction Register
	EIP uint32
	// Segment Register
	GS uint16
	FS uint16
	ES uint16
	DS uint16
	CS uint16
	SS uint16
	// FLAGS Register
	EFLAGS uint32
}

func (r *Registers) init() {
	v := reflect.ValueOf(r).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Uint16:
			f.Set(reflect.ValueOf(uint16(0)))
		case reflect.Uint32:
			if t.Field(i).Name == "EFLAGS" {
				// Reserved 1st bit, it's always 1 in EFLAGS.
				f.Set(reflect.ValueOf(uint32(2)))
			} else {
				f.Set(reflect.ValueOf(uint32(0)))
			}
		}
	}
}

func (r Registers) dump() {
	v := reflect.ValueOf(&r).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		registerName := t.Field(i).Name
		registerValue := v.Field(i).Interface()

		switch registerName {
		case "EFLAGS":
			fmt.Printf("%d: %s = %d (%032b)\n",
				i+1, registerName, registerValue, registerValue)
		default:
			fmt.Printf("%d: %s = %d\n",
				i+1, registerName, registerValue)
		}
	}
}

func main() {
	r := Registers{}
	r.init()
	r.dump()
}
