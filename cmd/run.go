package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/itzzsauravp/go-rem/helpers"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the reminder service",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			fmt.Fprint(os.Stderr, "This install-service currently only supports Linux.")
			return
		}
		exec.Command("systemctl", "--user", "start", "cue").Run()
		fmt.Printf("\n%s\n", fmt.Sprintf("Executed command %s...\n", helpers.ColorGreen("systemctl --user start cue")))

	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
}
