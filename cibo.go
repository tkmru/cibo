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

func (r Registers) dump() {
	s := reflect.ValueOf(&r).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s = %v\n",
			i, typeOfT.Field(i).Name, f.Interface())
	}
}

func main() {
	r := Registers{}
	r.dump()
}
