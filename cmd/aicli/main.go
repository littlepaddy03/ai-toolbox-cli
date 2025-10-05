package main

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var configDir string

func main() {
	root := &cobra.Command{Use: "at", Short: "AI toolbox CLI"}
	root.PersistentFlags().StringVar(&configDir, "config-dir", "",
		"config directory (defaults to ./configs relative to executable)")
	// Resolve config dir default = ./config next to executable
	if configDir == "" {
		exe, _ := os.Executable()
		base := filepath.Dir(exe)
		configDir = filepath.Join(base, "config")
	}
}
