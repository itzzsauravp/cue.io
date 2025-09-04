package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/itzzsauravp/go-rem/types"
)

func Notify(reminder types.Reminder) {
	fmt.Println(reminder)
	psCmd := fmt.Sprintf(`Import-Module BurntToast; New-BurntToastNotification -Text 'Cue Reminder','%s'`, reminder.Name)
	cmd := exec.Command(
		"powershell.exe",
		"-ExecutionPolicy",
		"Bypass",
		"-Command",
		psCmd,
	)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to send notification: %v", err)
	}
}
