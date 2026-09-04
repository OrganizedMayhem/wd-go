package cmd

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove points warping to nonexistent directories",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := newStore()
		if err != nil {
			return err
		}
		if err := store.Clean(); errors.Is(err, fs.ErrNotExist) {
			fmt.Fprintln(cmd.OutOrStdout(), "No warp points to clean.")
			return nil
		} else if err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Cleaned warp points.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
