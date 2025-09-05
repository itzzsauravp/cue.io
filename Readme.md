# CUE

**CUE** is a reminder CLI tool that helps you set reminders and get notifications on Linux and Windows via WSL **(not for native windows or macos yet)**.

---

```text
fnrr@archlnx ~/c/cue (master)> cue

cue.io is a small CLI project created by me for fun and to learn Golang + Cobra
Visit github.com/itzzsauravp/cue.io to view more details or contribute to it


 ██████╗██╗   ██╗███████╗
██╔════╝██║   ██║██╔════╝
██║     ██║   ██║█████╗
██║     ██║   ██║██╔══╝
╚██████╗╚██████╔╝███████╗
 ╚═════╝ ╚═════╝ ╚══════╝

██████╗ ██╗   ██╗
██╔══██╗╚██╗ ██╔╝
██████╔╝ ╚████╔╝
██╔══██╗  ╚██╔╝
██████╔╝   ██║
╚═════╝    ╚═╝

██╗████████╗███████╗███████╗███████╗ █████╗ ██╗   ██╗██████╗  █████╗ ██╗   ██╗██████╗
██║╚══██╔══╝╚══███╔╝╚══███╔╝██╔════╝██╔══██╗██║   ██║██╔══██╗██╔══██╗██║   ██║██╔══██╗
██║   ██║     ███╔╝   ███╔╝ ███████╗███████║██║   ██║██████╔╝███████║██║   ██║██████╔╝
██║   ██║    ███╔╝   ███╔╝  ╚════██║██╔══██║██║   ██║██╔══██╗██╔══██║╚██╗ ██╔╝██╔═══╝
██║   ██║   ███████╗███████╗███████║██║  ██║╚██████╔╝██║  ██║██║  ██║ ╚████╔╝ ██║
╚═╝   ╚═╝   ╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝  ╚═══╝  ╚═╝

```

```text
Usage:
  cue [flags]
  cue [command]

Available Commands:
  add             Add reminder
  completion      Generate the autocompletion script for the specified shell
  del             Delete reminder
  help            Help about any command
  install-service Install cue as a Linux systemd user service
  list            List reminder
  logs            Shows live logs for cue
  run             Start the reminder service
  serve           Run the daemon in background
  status          Stop the reminder service
  stop            Stop the reminder service

Flags:
  -h, --help   help for cue

Use "cue [command] --help" for more information about a command.
```

## Features

- Simple and minimal CLI interface.
- Cross-platform notifications (Linux & Windows via WSL).
- Recursive Reminders
- Auto-start as a systemd user service on Linux.

---

# Prerequisites

### General:

- Go 1.24.6

### For Linux:

### 1. **Install `libnotify`**<br/>

Depending on you distro use the following:

```bash
# Debian / Ubuntu
sudo apt update && sudo apt install libnotify-bin
```

```bash
# Fedora
sudo dnf update && sudo dnf install libnotify
```

```bash
# Arch
sudo pacman -Syu && sudo pacman -S libnotify
```

### For WSL:

### 1. **Install `wsl-notify-send.exe`**<br/>

Download the binary from [stuartleeks/wsl-notify-send](https://github.com/stuartleeks/wsl-notify-send/)

### 2. **Create a folder for the binary**<br/>

```bash
sudo mkdir -p "/mnt/c/Program Files/wsl-win"

# if in case this doesnot work try making one manually
# and save the downloaded binary under this folder
```

### 3. **Add the folder to your Windows `PATH`**

```bash
C:\Program Files\wsl-win
```

### 4. **Configure your shell**<br/>

- For Bash/Zsh: add this to your `~/.bashrc` or `~/.zshrc`

```bash
notify-send() { wsl-notify-send.exe --category $WSL_DISTRO_NAME "${@}"; }
```

```bash
# if bash
source ~/.bashrc

# if zsh
 source ~/.zshrc
```

- For Fish: add this to your `~/.config/fish/config.fish`

```bash
function notify-send
    wsl-notify-send.exe --category $WSL_DISTRO_NAME $argv
end
```

```bash
source ~/.config/fish/config.fish
```

<br>
<br>

# Installation

```bash
git clone https://github.com/itzzsauravp/cue.io.git
cd cue.io
go mod tidy
```

### Linux

```bash
make
# OR
make setup
# OR
make all
```

### WSL

```bash
make wsl
```

### To uninstall and revert all changes, simply run:

```bash
make clean
```

## Basic Usage

### Add a Reminder:

```bash
cue add -n <your reminders name> -d <duration>
```

Example:

```text
fnrr@archlnx ~/c/cue (master)> cue add -n "Remind me to write a Readme" -d 15m

Reminder added successfully
```

### List All Reminders:

```bash
cue list
```

Preview:

```text
fnrr@archlnx ~/c/cue (master)> cue list

ID  Name                                                             Duration  IsRecursive  IsActive
1   Welcome to CUE! Never forget a reminder again.                   1s        false        true
2   Time to stretch your legs! Take a quick 5-minute break.          5m0s      true         true
3   Drink water Staying hydrated is important for focus.             30m0s     true         true
4   Check your emails. Organize your inbox for better productivity.  1h0m0s    false        false
5   Evening reflection. Write down what you accomplished today.      2h0m0s    true         true
```

### Delete a Reminder:

```bash
cue del -i <reminder_id>
```

Preview:

```text
fnrr@archlnx ~/c/cue (master)> cue del -i 5

Reminder Deleted
```

## Applying Filters

You can apply some basic filters to list and del. `-r` and `-a` flags can be used to set the `IsRecursive` and `IsActive` field to true and false respectively

### Adding filters to the `list` and `del` commands

```bash
cue list -a <boolean_value> -r <boolean_value>
```

```text
fnrr@archlnx ~/c/cue (master)> cue list -r false

ID  Name                                                             Duration  IsRecursive  IsActive
1   Welcome to CUE! Never forget a reminder again.                   1s        false        true
4   Check your emails. Organize your inbox for better productivity.  1h0m0s    false        false
```

likewise you can use `-a` and `-r` independely or combined **_( same goes with `del` comamnd -> `cue del -a false` )_**

## Utility Commands:

These utility commands are just wrapper around the basic linux commands

### Start Service

```bash
cue run |> systemctl --user start cue
```

Note: This behavior is default after you run the `Makefile`. **_( No need to run this manually )_**

### Stop Service

```bash
cue stop |> systemctl --user stop cue
```

Note: Stops the service form running in the background and hence wont fire any notification

### Serve

```bash
cue serve
```

Note: Used as a test to see if the cue service is running as intended or not **_( Only for development purpose. No need to run this manually )_**

### Check Status

```bash
cue status |> systemctl --user status cue
```

Note: Check to see if the cue service is active or not.

### Cue Logs

```bash
cue logs |> journalctl --user -u cue -f
```

Note: Logs of cue service, when started, stopped, restarted etc.

## Future Improvements

- Integrate Calanders
- Crossplatform support accross all OS
- UX

## Feedback

Suggestions and constructive criticism are always welcome
