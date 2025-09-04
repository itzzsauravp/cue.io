package helpers

import (
	"fmt"
	"os"

	"github.com/itzzsauravp/go-rem/types"
)

func AddReminder(rem *types.Reminder) {

	reminders, err := LoadReminders()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while loading the reminders: %v", err)
	}

	reminders = append(reminders, *rem)
	err = SaveReminders(reminders)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while saving the reminders: %v", err)
	}
	fmt.Println("Reminder added successfully")
}
