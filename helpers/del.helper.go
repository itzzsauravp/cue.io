package helpers

import (
	"fmt"
	"os"

	"github.com/itzzsauravp/go-rem/internal/setup"
	"github.com/itzzsauravp/go-rem/types"
)

func DeleteAllReminders() {
	file, err := os.OpenFile(setup.FILE_PATH, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening the file @ `%s`: %v", setup.FILE_PATH, err)
	}
	defer file.Close()
	_, err = file.WriteString("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while clearing the contents of file @ `%s`: %v", setup.FILE_PATH, err)
	}
}

func DeleteReminder(query types.QueryDel) {
	var newReminders []types.Reminder

	reminders, err := LoadReminders()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading the reminders file: %v", err)
	}

	for _, reminder := range reminders {
		if query.IsActive != "" {
			activeFilter := query.IsActive == "true"
			if reminder.IsActive == activeFilter {
				continue
			}
		}
		if query.IsRecursive != "" {
			recursiveFilter := query.IsRecursive == "true"
			if reminder.IsRecursive == recursiveFilter {
				continue
			}
		}
		if query.ID != 0 {
			if reminder.Rem_Id == query.ID {
				continue
			}
		}
		newReminders = append(newReminders, reminder)
	}

	err = SaveReminders(newReminders)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while writing to the reminders file: %v", err)
	}
	fmt.Fprint(os.Stdout, "Reminder Deleted")
}
