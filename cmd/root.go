package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "diceware-cli",
		Short: "A generator of strong passwords using diceware passphrase.",
		Long: `diceware-cli let's you generate strong passwords based on easily memorable passphrases that are 
	also extremely resistant to attack.`,
		Version: GetDicewareCLIVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				fmt.Println("Use command 'generate' to start generating your strong passwords, or 'help' for instructions.")
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
