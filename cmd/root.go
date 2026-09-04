// Package cmd implements the command-line interface for wd.
package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "wd",
	Short:         "Warp to custom directories in terminal",
	Long:          `wd (warp directory) is a tool that lets you jump to custom directories in the terminal`,
	Args:          cobra.ArbitraryArgs,
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Fprintf(cmd.OutOrStdout(), "wd version %s\n", version)
			return nil
		}
		if len(args) == 0 {
			return cmd.Help()
		}

		path, err := getWarpPoint(args[0])
		if err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), path)
		return nil
	},
}

func getWarpPoint(name string) (string, error) {
	store, err := newStore()
	if err != nil {
		return "", err
	}
	path, err := store.Get(name)
	if errors.Is(err, fs.ErrNotExist) {
		return "", errors.New("no warp points yet")
	}
	return path, err
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "print version")
}
