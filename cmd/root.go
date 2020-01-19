package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diceware",
	Short: "A generator of diceware passwords.",
	Long: `diceware-cli let's you generate strong passwords based on easily memorable passwords that are 
	also extremely resistant to attack.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use command generate to start generating your strong passwords!")
	},
}

func Execute() {
	generateCmd.Flags().StringVarP(&lang, "lang", "l", "en", "password language")
	generateCmd.Flags().Int32VarP(&size, "size", "s", 6, "the amount words the password will have")
	generateCmd.Flags().BoolVarP(&pbcopy, "copy", "c", false, "pbcopy password")
	generateCmd.Flags().BoolVar(&hide, "hide", false, "pbcopy and hide password. Password WON'T be printed out")

	rootCmd.AddCommand(generateCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
