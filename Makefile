BINARY_NAME = cue
BUILD_DIR = .
INSTALL_DIR = /usr/local/bin
CONFIG_DIR = $(HOME)/.config/cue
FILE = $(CONFIG_DIR)/reminders.json

.PHONY = all build install setup clean

all: setup

build:
	go build -o $(BINARY_NAME) $(BUILD_DIR)/main.go

install: build
	sudo ln -sf $(PWD)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to $(INSTALL_DIR)"

# init: 
# 	mkdir -p $(CONFIG_DIR)
# 	touch $(FILE)
# 	@echo "Initialized file at $(CONFIG_DIR)"

setup: build install
	@echo "Setup completed! You can now run `$(BINARY_NAME)` from anywhere."

clean: 
	rm -rf $(BINARY_NAME)
	sudo rm -rf /usr/local/bin/$(BINARY_NAME)
	rm -rf $(CONFIG_DIR)
	systemctl --user stop cue
	systemctl --user disable cue
	sudo rm $(HOME)/.config/systemd/user/$(BINARY_NAME).service
	systemctl --user daemon-reload
