package cmd

import (
	"fmt"

	"github.com/itzzsauravp/go-rem/helpers"
	"github.com/itzzsauravp/go-rem/types"
	"github.com/spf13/cobra"
)

var del_active string
var del_recursive string
var del_id int

var DeleteCmd = &cobra.Command{
	Use:       "del",
	Short:     "Delete reminder",
	Long:      "This command will delete reminders when paired with flags like -i, -r, -a",
	ValidArgs: []cobra.Completion{"all"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 && args[0] == "all" {
			fmt.Println("Deleting all your reminders")
			helpers.DeleteAllReminders()
		} else {
			query := types.QueryDel{
				ID:          del_id,
				IsActive:    del_active,
				IsRecursive: del_recursive,
			}
			helpers.DeleteReminder(query)
		}
	},
}

func init() {

	DeleteCmd.Flags().IntVarP(&del_id, "id", "i", 0, "Deletes a reminders of id i")
	DeleteCmd.Flags().StringVarP(&del_active, "active", "a", "", "Deletes all the active reminders")
	DeleteCmd.Flags().StringVarP(&del_recursive, "recursive", "r", "", "Deletes all the recursive reminders")

	RootCmd.AddCommand(DeleteCmd)
}
