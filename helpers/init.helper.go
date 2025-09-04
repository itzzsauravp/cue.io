package helpers

import (
	"os"
	"path/filepath"

	"github.com/itzzsauravp/cue.io/internal/setup"
)

func PreRun(filePath string) {
	_, err := os.Stat(setup.FILE_PATH)
	if os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(filePath), 0755)
		if err != nil {
			panic(err)
		}

		_, err = os.Create(filePath)
		if err != nil {
			panic(err)
		}
	}
}
