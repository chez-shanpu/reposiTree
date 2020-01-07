package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "repotr",
	Short: "Convert repository to tree structure and analyze",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		return err
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
