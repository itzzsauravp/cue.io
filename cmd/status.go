package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Stop the reminder service",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			fmt.Fprint(os.Stderr, "Error while stopping cue.service")
			return
		}
		b, err := exec.Command("systemctl", "--user", "status", "cue").Output()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error while reading status for cue.service")
			return
		}
		fmt.Println(string(b))
	},
}

func init() {
	RootCmd.AddCommand(StatusCmd)
}
