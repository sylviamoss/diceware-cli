package cmd

import (
	"diceware-cli/diceware"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	customConfig diceware.CustomConfig

	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Adds new language dictionary.",
		Long:  `Adds new language dictionary.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := customConfig.Configure()
			if err != nil {
				fmt.Printf(err.Error())
			}
		},
	}
)
