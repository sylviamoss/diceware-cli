package cmd

import (
	"diceware-cli/diceware"

	"github.com/spf13/cobra"
)

var (
	lang string
	size int32

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generates a diceware password with custom configuration.",
		Long: `Generates strong passwords based on easily memorable passwords that are 
	also extremely resistant to attack.`,
		Run: func(cmd *cobra.Command, args []string) {
			diceware.Generate(lang, size)
		},
	}
)
