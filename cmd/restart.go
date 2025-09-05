package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/itzzsauravp/cue.io/helpers"
	"github.com/spf13/cobra"
)

var RestartCmd = &cobra.Command{
	Use:    "restart",
	Hidden: true,
	Short:  "Restarts the reminder service",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			fmt.Fprint(os.Stderr, "This install-service currently only supports Linux.")
			return
		}
		helpers.RestartDaemon()
		fmt.Println("cue service restarted")

	},
}

func init() {
	RootCmd.AddCommand(RestartCmd)
}
