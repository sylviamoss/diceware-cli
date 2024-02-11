package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sylviamoss/diceware-cli/diceware"
)

func init() {
	dictionaryCmd.Flags().BoolVar(&dictionary.AddLang, "add-lang", false, "add new config language")
	dictionaryCmd.Flags().StringVar(&dictionary.Source, "source", "", "dictionary source file")
	dictionaryCmd.Flags().StringVar(&dictionary.Name, "name", "", "language name")
	rootCmd.AddCommand(dictionaryCmd)
}

var (
	dictionary diceware.Dictionary

	dictionaryCmd = &cobra.Command{
		Use:   "dictionary",
		Short: "Manages the diceware words dictionary.",
		Long:  `Manages the diceware words dictionary. Allows adding a new language dictionary.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := dictionary.Configure(); err != nil {
				errorMsg := fmt.Sprintf("Ops...something went wrong: %s", err.Error())
				return errors.New(errorMsg)
			}
			return nil
		},
	}
)
