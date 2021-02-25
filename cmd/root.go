package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version           = "1.3.0"
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

	generateCmd.Flags().StringVarP(&generateConfig.Lang, "lang", "l", "en", "password language\n available langs: en, pt")
	generateCmd.Flags().StringVar(&generateConfig.Separator, "separator", "/", "character that separates the words.\nuse --separator=none to remove reparator")
	generateCmd.Flags().Int32VarP(&generateConfig.Size, "size", "s", 6, "the amount words the password will have")
	generateCmd.Flags().BoolVarP(&generateConfig.Pbcopy, "copy", "c", false, "pbcopy password")
	generateCmd.Flags().BoolVar(&generateConfig.Hide, "hide", false, "pbcopy and hide password. Password WON'T be printed out")
	rootCmd.AddCommand(generateCmd)

	configCmd.Flags().BoolVarP(&customConfig.Add, "add", "a", false, "add new config")
	configCmd.Flags().BoolVarP(&customConfig.Lang, "lang", "l", false, "add new language")
	configCmd.Flags().StringVarP(&customConfig.Source, "source", "s", "", "dictionary source file")
	configCmd.Flags().StringVarP(&customConfig.Name, "name", "n", "", "language name")
	rootCmd.AddCommand(configCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
