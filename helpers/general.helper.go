package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/itzzsauravp/go-rem/internal/setup"
	"github.com/itzzsauravp/go-rem/types"
)

func LoadReminders() ([]types.Reminder, error) {
	data, err := os.ReadFile(setup.FILE_PATH)
	if err != nil {
		return nil, err
	}
	var reminders []types.Reminder

	if len(data) == 0 {
		return []types.Reminder{}, nil
	}
	err = json.Unmarshal(data, &reminders)
	return reminders, err
}

func SaveReminders(reminders []types.Reminder) error {
	data, err := json.MarshalIndent(reminders, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(setup.FILE_PATH, data, 0644)
}

func DisplayReminders(reminders []types.Reminder, query *types.Query) error {

	if len(reminders) == 0 {
		fmt.Fprintln(os.Stdout, "\nNo reminders found\n\nPlease check", ColorCyan("cue --help"), "or", ColorCyan("cue add --help"))
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tDuration\tIsRecursive\tIsActive")

	for _, reminder := range reminders {
		if query.IsActive != "" {
			activeFilter := query.IsActive == "true"
			if reminder.IsActive != activeFilter {
				continue
			}
		}
		if query.IsRecursive != "" {
			recursiveFilter := query.IsRecursive == "true"
			if reminder.IsRecursive != recursiveFilter {
				continue
			}
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%t\t%s\n",
			reminder.Rem_Id, reminder.Name, reminder.Duration.String(),
			reminder.IsRecursive, ColorIsActive(reminder.IsActive))
	}
	return w.Flush()
}
