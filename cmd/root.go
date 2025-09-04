package cmd

import (
	"github.com/itzzsauravp/go-rem/constants"
	"github.com/itzzsauravp/go-rem/helpers"
	"github.com/itzzsauravp/go-rem/internal/setup"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   constants.ROOT_CMD,
	Short: "REM, Your personal reminder in CLI",
	Long:  `REM is a small project created my "@itzzsauravp" to keep track of your reminders through CLI`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		helpers.PreRun(setup.FILE_PATH)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() error {
	return RootCmd.Execute()
}
