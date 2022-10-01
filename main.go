package main

import (
	"os"

	"github.com/akastav/argocd-terraform-plugin/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
