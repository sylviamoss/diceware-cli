package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version           = "1.3.3"
	VersionPrerelease = "dev"
	rootCmd           = &cobra.Command{
		Use:   "diceware-cli",
		Short: "A generator of strong passwords using diceware passphrase.",
		Long: `diceware-cli let's you generate strong passwords based on easily memorable passphrases that are 
	also extremely resistant to attack.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				fmt.Println("Use command 'generate' to start generating your strong passwords, or 'help' for instructions.")
			}
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "diceware-cli version",
		Run: func(cmd *cobra.Command, args []string) {
			var versionString bytes.Buffer
			fmt.Fprintf(&versionString, "%s", Version)
			if VersionPrerelease != "" {
				fmt.Fprintf(&versionString, "-%s", VersionPrerelease)
			}
			fmt.Printf("diceware-cli %s\n", versionString.String())
		},
	}
)

func Execute() {
	rootCmd.AddCommand(versionCmd)

	generateCmd.Flags().StringVar(&dicewareConfig.Lang, "lang", "en", "password language\n available langs: en, pt")
	generateCmd.Flags().StringVar(&dicewareConfig.Separator, "separator", "/", "character that separates the words.\nuse --separator=none to remove reparator")
	generateCmd.Flags().Int32Var(&dicewareConfig.Size, "size", 6, "the amount words the password will have")
	generateCmd.Flags().BoolVar(&dicewareConfig.Pbcopy, "copy", false, "pbcopy password")
	generateCmd.Flags().BoolVar(&dicewareConfig.Hide, "hide", false, "pbcopy and hide password. You WON'T see the password")
	generateCmd.Flags().BoolVar(&dicewareConfig.Lower, "lower", false, "remove capitalized first letters")
	generateCmd.Flags().BoolVar(&dicewareConfig.RemoveNumber, "remove-number", false, "removes the random number we add by default")
	rootCmd.AddCommand(generateCmd)

	configCmd.Flags().BoolVar(&customConfig.AddLang, "add-lang", false, "add new config language")
	configCmd.Flags().StringVar(&customConfig.Source, "source", "", "dictionary source file")
	configCmd.Flags().StringVar(&customConfig.Name, "name", "", "language name")
	rootCmd.AddCommand(configCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
