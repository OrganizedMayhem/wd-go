package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var addcdCmd = &cobra.Command{
	Use:   "addcd <path> [point]",
	Short: "Adds a path to your warp points",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		point := filepath.Base(path)
		if len(args) > 1 {
			point = args[1]
		}

		store, err := newStore()
		if err != nil {
			return err
		}
		if err := store.Put(point, path); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Added warp point '%s' to '%s'\n", point, path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addcdCmd)
}
