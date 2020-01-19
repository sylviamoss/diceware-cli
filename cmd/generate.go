package cmd

import (
	"diceware-cli/diceware"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	lang   string
	size   int32
	pbcopy bool
	hide   bool

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generates a diceware password with custom configuration.",
		Long: `Generates strong passwords based on easily memorable passwords that are 
	also extremely resistant to attack.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := diceware.Generate(lang, size, pbcopy, hide)
			if err != nil {
				fmt.Printf("Ops...something went wrong: %s", err.Error())
			}
		},
	}
)
