package cmd

import (
	"fmt"
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

func init() {
	RootCmd.AddCommand(newVersionCmd())
}

func newVersionCmd() *cobra.Command {
	GoVersion := runtime.Version()
	Compiler := runtime.Compiler

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("reposiTree Version: %s (Revision: %s / GoVersion: %s / Compiler: %s)\n",
				Version, Revision, GoVersion, Compiler)
		},
	}

	return cmd
}
