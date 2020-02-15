package main

import (
	"diceware-cli/cmd"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./words")
	cmd.Execute(box)
}
