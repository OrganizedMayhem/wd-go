package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls <point>",
	Short: "Show files from given warp point (ls)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := getWarpPoint(args[0])
		if err != nil {
			return err
		}
		entries, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("list %s: %w", path, err)
		}
		for _, entry := range entries {
			fmt.Fprintln(cmd.OutOrStdout(), entry.Name())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
