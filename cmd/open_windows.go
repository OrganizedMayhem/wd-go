//go:build windows

package cmd

import "os/exec"

func openDir(path string) *exec.Cmd {
	return exec.Command("cmd", "/c", "start", "", path)
}
