package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.10.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wd",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "wd version %s\n", version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
