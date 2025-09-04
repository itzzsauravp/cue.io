package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/itzzsauravp/cue.io/constants"
	"github.com/spf13/cobra"
)

var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Shows live logs for cue",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			fmt.Fprintf(os.Stderr, "This build currently only supports linux")
			return
		}
		journalCmd := exec.Command("journalctl", "--user", "-u", constants.ROOT_CMD, "-f")

		//  returns the child process standard output and errors to the parent go process ???
		journalCmd.Stdout = os.Stdout
		journalCmd.Stderr = os.Stderr

		err := journalCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run journalctl: %v\n", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(LogsCmd)
}
