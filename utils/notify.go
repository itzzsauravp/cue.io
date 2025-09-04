package utils

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/gen2brain/beeep"
	"github.com/itzzsauravp/go-rem/constants"
	"github.com/itzzsauravp/go-rem/types"
)

func Notify(reminder types.Reminder) {

	if runtime.GOOS == "linux" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get user home directory: %v", err)
			return
		}
		user, _ := user.Current()
		var username string = "user"
		if user.Name != "" {
			username = user.Name
		}
		imagePath := fmt.Sprintf("%s/.config/cue/assets/%s", home, constants.NOTIFICATION_ICON)
		fmt.Println(imagePath)
		fmt.Println(reminder)
		err = beeep.Notify(fmt.Sprintf("Hey %s\nCue Reminder", username), reminder.Name, imagePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to send notification: %v", err)
		}
	}

	// this is for windows wip
	// psCmd := fmt.Sprintf(`Import-Module BurntToast; New-BurntToastNotification -Text 'Cue Reminder','%s'`, reminder.Name)
	// cmd := exec.Command(
	// 	"powershell.exe",
	// 	"-ExecutionPolicy",
	// 	"Bypass",
	// 	"-Command",
	// 	psCmd,
	// )
	// if err := cmd.Run(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "failed to send notification: %v", err)
	// }

}
