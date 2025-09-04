package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the reminder service",
	Run: func(cmd *cobra.Command, args []string) {
		switch runtime.GOOS {
		case "linux":
			exec.Command("systemctl", "--user", "stop", "cue").Run()
		default:
			fmt.Println("Unsupported OS: please kill manually")
		}
	},
}

func init() {
	RootCmd.AddCommand(StopCmd)
}
