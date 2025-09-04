package setup

import (
	"fmt"
	"os"
	"runtime"
)

const ()

var FILE_PATH string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while finding the home directory: %v", err)
	}
	if runtime.GOOS == "linux" {
		FILE_PATH = fmt.Sprintf("%s/.config/cue/reminders.json", homeDir)
		return
	}
	fmt.Fprintf(os.Stderr, "This build currently only supports linux: %v", err)
}
