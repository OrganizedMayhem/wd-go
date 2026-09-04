package cmd

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all stored warp points",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := newStore()
		if err != nil {
			return err
		}
		points, err := store.Load()
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Fprintln(cmd.OutOrStdout(), "No warp points yet. Add one with 'wd add'.")
			return nil
		}
		if err != nil {
			return err
		}
		for _, point := range points {
			fmt.Fprintf(cmd.OutOrStdout(), "%s:%s\n", point.Name, point.Path)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
