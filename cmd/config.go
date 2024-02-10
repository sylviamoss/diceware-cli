package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sylviamoss/diceware-cli/diceware"
)

func init() {
	configCmd.Flags().BoolVar(&customConfig.AddLang, "add-lang", false, "add new config language")
	configCmd.Flags().StringVar(&customConfig.Source, "source", "", "dictionary source file")
	configCmd.Flags().StringVar(&customConfig.Name, "name", "", "language name")
	rootCmd.AddCommand(configCmd)
}

var (
	customConfig diceware.CustomConfig

	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Adds new language dictionary.",
		Long:  `Adds new language dictionary.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := customConfig.Configure(); err != nil {
				errorMsg := fmt.Sprintf("Ops...something went wrong: %s", err.Error())
				return errors.New(errorMsg)
			}
			return nil
		},
	}
)
