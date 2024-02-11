package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sylviamoss/diceware-cli/diceware"
)

func init() {
	generateCmd.Flags().StringVar(&dicewareConfig.Lang, "lang", "en", "password language\n available langs: en, pt")
	generateCmd.Flags().StringVar(&dicewareConfig.Separator, "separator", "/", "character that separates the words.\nuse --separator=none to remove reparator")
	generateCmd.Flags().Int32Var(&dicewareConfig.Size, "size", 6, "the amount words the password will have")
	generateCmd.Flags().BoolVar(&dicewareConfig.Pbcopy, "copy", false, "pbcopy password")
	generateCmd.Flags().BoolVar(&dicewareConfig.Hide, "hide", false, "pbcopy and hide password. You WON'T see the password")
	generateCmd.Flags().BoolVar(&dicewareConfig.Lower, "lower", false, "remove capitalized first letters")
	generateCmd.Flags().BoolVar(&dicewareConfig.RemoveNumber, "remove-number", false, "removes the random number we add by default")

	// Configure viper to read from the config file, if set
	viper.BindPFlag("generate.lang", generateCmd.Flags().Lookup("lang"))
	viper.BindPFlag("generate.separator", generateCmd.Flags().Lookup("separator"))
	viper.BindPFlag("generate.size", generateCmd.Flags().Lookup("size"))
	viper.BindPFlag("generate.copy", generateCmd.Flags().Lookup("copy"))
	viper.BindPFlag("generate.hide", generateCmd.Flags().Lookup("hide"))
	viper.BindPFlag("generate.lower", generateCmd.Flags().Lookup("lower"))
	viper.BindPFlag("generate.remove-number", generateCmd.Flags().Lookup("remove-number"))

	rootCmd.AddCommand(generateCmd)
}

var (
	dicewareConfig diceware.Config

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generates a diceware password with custom configuration.",
		Long: `Generates strong passwords based on easily memorable passwords that are 
	also extremely resistant to attack.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read the from config file, otherwise default value will be used
			dicewareConfig = diceware.Config{
				Lang:         viper.GetString("generate.lang"),
				Separator:    viper.GetString("generate.separator"),
				Size:         viper.GetInt32("generate.size"),
				Pbcopy:       viper.GetBool("generate.copy"),
				Hide:         viper.GetBool("generate.hide"),
				Lower:        viper.GetBool("generate.lower"),
				RemoveNumber: viper.GetBool("generate.remove-number"),
			}

			if err := dicewareConfig.Generate(); err != nil {
				errorMsg := fmt.Sprintf("Ops...something went wrong: %s", err.Error())
				return errors.New(errorMsg)
			}
			return nil
		},
	}
)
