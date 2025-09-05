package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the reminder service",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			fmt.Fprint(os.Stderr, "This install-service currently only supports Linux.")
			return
		}
		exec.Command("systemctl", "--user", "start", "cue").Run()
		fmt.Println("cue is now running in the backgroud...")

	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
}
