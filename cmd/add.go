package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [point]",
	Short: "Adds the current working directory to your warp points",
	Long: `Adds the current working directory to your warp points.
If no point is specified, the current directory's name will be used.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}

		point := filepath.Base(path)
		if len(args) > 0 {
			point = args[0]
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
	rootCmd.AddCommand(addCmd)
}
