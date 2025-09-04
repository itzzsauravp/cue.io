package utils

import (
	"fmt"
	"time"

	"github.com/itzzsauravp/go-rem/types"
)

func StartReminder(reminders []types.Reminder) {
	fmt.Println("\nReminder has been started...")
	for _, rem := range reminders {
		if !rem.IsActive {
			continue
		}
		r := rem
		if !rem.IsRecursive {
			go func() {
				time.Sleep(rem.Duration)
				Notify(r)
			}()
		} else {
			go func() {
				ticker := time.NewTicker(rem.Duration)
				defer ticker.Stop()
				for {
					<-ticker.C
					Notify(r)
				}
			}()
		}
	}

}
