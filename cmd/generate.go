package cmd

import (
	"diceware-cli/diceware"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	generateConfig diceware.GenerateConfig

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generates a diceware password with custom configuration.",
		Long: `Generates strong passwords based on easily memorable passwords that are 
	also extremely resistant to attack.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := generateConfig.Generate(); err != nil {
				fmt.Printf("Ops...something went wrong: %s", err.Error())
			}
		},
	}
)
