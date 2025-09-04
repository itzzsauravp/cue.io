package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/itzzsauravp/go-rem/helpers"
	"github.com/itzzsauravp/go-rem/types"
	"github.com/spf13/cobra"
)

var name string
var duration string
var add_recursive bool
var add_active bool

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add reminder",
	Long:  "This command paired with flags like -n, -t, -r, -a is used to create a reminder form the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		dur, err := time.ParseDuration(duration)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse the duration (Ex: 10m, 20s, 5h) :%v\n", err)
		}
		extReminders, _ := helpers.LoadReminders()
		rem := types.Reminder{
			Rem_Id:      len(extReminders) + 1,
			Name:        name,
			Duration:    dur,
			IsRecursive: add_recursive,
			IsActive:    add_active,
		}
		helpers.AddReminder(&rem)
	},
}

func init() {

	AddCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the reminder*")
	AddCmd.Flags().StringVarP(&duration, "duration", "d", "", "Duration of the reminder*")
	AddCmd.Flags().BoolVarP(&add_recursive, "recursive", "r", false, "Is reminder recursive default=false, optional")
	AddCmd.Flags().BoolVarP(&add_active, "active", "a", true, "Is reminder active. defualt=true, optional")

	AddCmd.MarkFlagRequired("name")
	AddCmd.MarkFlagRequired("duration")
	RootCmd.AddCommand(AddCmd)
}
