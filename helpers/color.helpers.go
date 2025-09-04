package helpers

import (
	"github.com/fatih/color"
)

var ColorGreen = color.New(color.FgGreen, color.Bold).SprintFunc()
var ColorRed = color.New(color.FgRed, color.Bold).SprintFunc()

var ColorYellow = color.New(color.FgYellow, color.Bold).SprintFunc()
var ColorCyan = color.New(color.FgCyan, color.Bold).SprintFunc()

func ColorIsRecursive(value bool) string {
	if value {
		return ColorCyan(value)
	}
	return ColorYellow(value)
}

func ColorIsActive(value bool) string {
	if value {
		return ColorGreen(value)
	}
	return ColorRed(value)
}
