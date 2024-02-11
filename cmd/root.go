package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.diceware-cli.yaml)")
}

var (
	cfgFile string

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

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("No configuration file found: ", err.Error())
				return
			}
			fmt.Println("Error reading config file: ", err.Error())
		}

	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".diceware-cli")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("No configuration file found: ", err.Error())
				return
			}
			fmt.Println("Error reading config file: ", err.Error())
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
