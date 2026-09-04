package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <point>",
	Short: "Open the warp point in the default file explorer",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := getWarpPoint(args[0])
		if err != nil {
			return err
		}
		open := openDir(path)
		open.Stdout = cmd.OutOrStdout()
		open.Stderr = cmd.ErrOrStderr()
		if err := open.Run(); err != nil {
			return fmt.Errorf("open %s: %w", path, err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
