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
	fmt.Println("User:", username)
	fmt.Println("User:", runtime.GOOS)
	switch runtime.GOOS {
	case "linux":
		if helpers.IsWSL() {
			cmd := exec.Command("notify-send", `"Hello There"`)
			fmt.Println(cmd.String())
			cmd.Env = append(cmd.Env, os.Environ()...)
			cmd.Run()
		} else {
			imagePath := fmt.Sprintf("%s/.config/cue/assets/%s", home, constants.NOTIFICATION_ICON)
			err = beeep.Notify(fmt.Sprintf("Hey %s\nCue Reminder", username), reminder.Name, imagePath)
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
