package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sylviamoss/diceware-cli/diceware"
)

func init() {
	cobra.OnInitialize(initConfig)

	configCmd.Flags().StringVarP(&config.file, "file", "f", "", "config file (default is $HOME/.diceware-cli.yaml)")

	// Deprecated flags
	configCmd.Flags().BoolVar(&config.AddLang, "add-lang", false, "add new config language")
	configCmd.Flags().MarkDeprecated("add-lang", "please use the equivalent flag in the 'dictionary' command instead")

	configCmd.Flags().StringVar(&config.Source, "source", "", "dictionary source file")
	configCmd.Flags().MarkDeprecated("source", "please use the equivalent flag in the 'dictionary' command instead")

	configCmd.Flags().StringVar(&config.Name, "name", "", "language name")
	configCmd.Flags().MarkDeprecated("name", "please use the equivalent flag in the 'dictionary' command instead")

	configCmd.AddCommand(generateConfigCmd)
	rootCmd.AddCommand(configCmd)

}

type Config struct {
	diceware.Dictionary
	file string
}

var (
	config Config

	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configures the diceware-cli config file, used to override defaults of flags",
		Long:  `Configures the diceware-cli settings, such as overriding the default of the generate command flags.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.Configure(); err != nil {
				errorMsg := fmt.Sprintf("Ops...something went wrong: %s", err.Error())
				return errors.New(errorMsg)
			}
			return nil
		},
	}

	generateConfigCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generates the content of a diceware-cli yaml config file",
		Long: `Generates a diceware-cli yaml config content to be used to override the default values of command flags.

The default location of the file is $HOME/.diceware-cli.yaml, but you can override it using 'diceware-cli --config=/path/to/config.yaml'`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`# diceware-cli config file yaml content
# You can customize the default values of the flags by setting them in this file.
generate:
  lang: en
  separator: /
  size: 6
  copy: false
  hide: false
  lower: false
  remove-number: false`)
		},
	}
)

func initConfig() {
	if config.file != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config.file)

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
