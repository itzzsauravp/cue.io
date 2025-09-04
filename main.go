package main

import (
	"log"

	"github.com/itzzsauravp/cue.io/cmd"
)

func main() {

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
