package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	Version           = "1.3.3"
	VersionPrerelease = "dev"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "diceware-cli version",
		Run: func(cmd *cobra.Command, args []string) {
			version := GetDicewareCLIVersion()
			fmt.Printf("diceware-cli version %s\n", version)
		},
	}
)

func GetDicewareCLIVersion() string {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "%s", Version)
	if VersionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", VersionPrerelease)
	}
	return versionString.String()
}
