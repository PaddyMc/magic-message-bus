package main

import (
	"os"

	"github.com/allinbits/magic-message-bus/cmd/magic-message-busd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
