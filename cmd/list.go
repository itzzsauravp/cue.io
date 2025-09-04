package cmd

import (
	"fmt"
	"os"

	"github.com/itzzsauravp/cue.io/helpers"
	"github.com/itzzsauravp/cue.io/types"
	"github.com/spf13/cobra"
)

var list_active string
var list_recursive string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List reminder",
	Long:  "The command will list all your reminders active, inactive, recursive, non-recursive",
	Run: func(cmd *cobra.Command, args []string) {
		reminders, err := helpers.LoadReminders()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading the data: %v \n", err)
			return
		}

		query := &types.Query{
			IsActive:    list_active,
			IsRecursive: list_recursive,
		}

		helpers.DisplayReminders(reminders, query)
	},
}

func init() {

	listCmd.Flags().StringVarP(&list_active, "active", "a", "", "Lists active or inactive reminders (-a true || -a false)")
	listCmd.Flags().StringVarP(&list_recursive, "recursive", "r", "", "Lists recursive or non-recursive reminders (-r true || -r false)")

	RootCmd.AddCommand(listCmd)
}
