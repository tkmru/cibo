package main

import (
	"os"
	"github.com/tkmru/cibo/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
