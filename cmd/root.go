package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// set these value as go build -ldflags option
	// Version is version number which automatically set on build. `git describe --tags`
	Version string
	// Revision is git commit hash which automatically set `git rev-parse --short HEAD` on build.
	Revision string
)

func Execute() {
	GoVersion := runtime.Version()
	Compiler := runtime.Compiler
	RootCmd := &cobra.Command{
		Use:   "repotr",
		Short: "Convert repository to tree structure and analyze",
		Version: fmt.Sprintf("reposiTree Version: %s (Revision: %s / GoVersion: %s / Compiler: %s)\n",
			Version, Revision, GoVersion, Compiler),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			return err
		},
	}

	RootCmd.AddCommand(NewTreeCmd())
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
