package cmd

import (
	"fmt"
	"os"

	"github.com/itzzsauravp/go-rem/helpers"
	"github.com/itzzsauravp/go-rem/utils"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the daemon in background",
	Run: func(cmd *cobra.Command, args []string) {
		reminders, err := helpers.LoadReminders()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while loading reminders: %v", err)
		}
		utils.StartReminder(reminders)

		// somehow this keeps running the main go routine forever (idk how)
		select {}
	},
}

func init() {
	RootCmd.AddCommand(ServeCmd)
}
