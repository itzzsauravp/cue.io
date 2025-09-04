package main

import (
	"log"

	"github.com/itzzsauravp/go-rem/cmd"
)

func main() {

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
