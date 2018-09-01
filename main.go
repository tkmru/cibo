package main

import (
	"github.com/tkmru/cibo/cmd"
	"os"
)

func main() {
	cmd.Execute(os.Args[1:])
}
