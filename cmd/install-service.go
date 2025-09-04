package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/itzzsauravp/go-rem/helpers"
	"github.com/spf13/cobra"
)

var InstallServiceCmd = &cobra.Command{
	Use:   "install-service",
	Short: "Install cue as a Linux systemd user service",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			fmt.Fprint(os.Stderr, "This install-service currently only supports Linux.")
			return
		}

		// havent checked is this thing for symlinks
		exe, err := os.Executable()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting executable path: %v\n", err)
			return
		}
		servicePath := filepath.Join(os.Getenv("HOME"), ".config/systemd/user/cue.service")
		os.MkdirAll(filepath.Dir(servicePath), 0755)
		fmt.Println("This is the absolute path for the binary:", exe)
		fmt.Println("This is the path for the service:", servicePath)

		unit := fmt.Sprintf(`[Unit]
Description=cue Reminder Daemon

[Service]
ExecStart=%s serve
Restart=always

[Install]
WantedBy=default.target
`, exe)
		fmt.Println("This is the content of the service file", unit)
		err = os.WriteFile(servicePath, []byte(unit), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write service file: %v\n", err)
			return
		}

		exec.Command("systemctl", "--user", "daemon-reload").Run()
		exec.Command("systemctl", "--user", "enable", "cue").Run()

		fmt.Printf("Service installed! You can now run %s to start it and %s to stop it.\n", helpers.ColorCyan(`cue run`), helpers.ColorRed(`cue stop`))

	},
}

func init() {
	RootCmd.AddCommand(InstallServiceCmd)
}
