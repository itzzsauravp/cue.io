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
		panic(err)
	}

	switch runtime.GOOS {
	case "linux":
		FILE_PATH = fmt.Sprintf("%s/.config/gorem/reminders.json", homeDir)
	case "darwin":
		FILE_PATH = fmt.Sprintf("%s/Library/Application Support/gorem/reminders.json", homeDir)
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			appData = homeDir
		}
		FILE_PATH = fmt.Sprintf("%s\\gorem\\reminders.json", appData)
	default:
		FILE_PATH = fmt.Sprintf("%s/reminders.json", homeDir)
	}
}
