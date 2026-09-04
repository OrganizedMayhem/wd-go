package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [point]",
	Short: "Print path to given warp point",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			path, err := getWarpPoint(args[0])
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), path)
			return nil
		}

		store, err := newStore()
		if err != nil {
			return err
		}
		points, err := store.Load()
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Fprintln(cmd.OutOrStdout(), "No warp points for current directory.")
			return nil
		}
		if err != nil {
			return err
		}

		currentPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		var names []string
		for _, point := range points {
			if point.Path == currentPath {
				names = append(names, point.Name)
			}
		}
		if len(names) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No warp points for current directory.")
			return nil
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Warp points for current directory: %s\n", strings.Join(names, ", "))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
