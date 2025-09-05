package utils

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"

	"github.com/gen2brain/beeep"
	"github.com/itzzsauravp/cue.io/constants"
	"github.com/itzzsauravp/cue.io/helpers"
	"github.com/itzzsauravp/cue.io/types"
)

func Notify(reminder types.Reminder) {

	var username string
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get user home directory: %v", err)
		return
	}
	user, err := user.Current()
	username = user.Username
	if err != nil {
		username = "User"
	}
	switch runtime.GOOS {
	case "linux":
		if helpers.IsWSL() {
			fmt.Println("\nThis is running on WSL")
			cmd := exec.Command(
				"/mnt/c/Program Files/wsl-win/wsl-notify-send.exe",
				"--icon", constants.NOTIFICATION_ICON_PATH_WIN,
				"--category", fmt.Sprintf("Hey %s", username),
				"--appId", constants.APP_ID,
				"--expire-time", "-1",
				reminder.Name,
			)
			fmt.Println(cmd)
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "There was an error executing the command: %v", err)
			}
		} else {
			fmt.Println("This is running on LINUX")
			imagePath := fmt.Sprintf("%s%s", home, constants.NOTIFICATION_ICON_PATH)
			err = beeep.Notify(fmt.Sprintf("Hey %s\n", username), reminder.Name, imagePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to send notification: %v", err)
			}

		}
	case "windows":
		exec.Command("notify-send", "Hello There").Run()
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
