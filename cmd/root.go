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
	generateCmd.Flags().StringVarP(&generateConfig.Lang, "lang", "l", "en", "password language\n available langs: en, pt")
	generateCmd.Flags().StringVar(&generateConfig.Separator, "separator", " ", "character that separates the words,\n to remove reparator use --separator=none")
	generateCmd.Flags().Int32VarP(&generateConfig.Size, "size", "s", 6, "the amount words the password will have")
	generateCmd.Flags().BoolVarP(&generateConfig.Pbcopy, "copy", "c", false, "pbcopy password")
	generateCmd.Flags().BoolVar(&generateConfig.Hide, "hide", false, "pbcopy and hide password. Password WON'T be printed out")

	rootCmd.AddCommand(generateCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
