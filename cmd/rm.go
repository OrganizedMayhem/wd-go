package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm <point>",
	Short: "Removes the given warp point",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := newStore()
		if err != nil {
			return err
		}
		if err := store.Remove(args[0]); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Removed warp point '%s'\n", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
