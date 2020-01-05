package cmd

import (
	"github.com/chez-shanpu/reposiTree/pkg/tree"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(NewTreeCmd())
}

func NewTreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tree",
		Short: "repository tree",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			return err
		},
	}

	cmd.AddCommand(tree.NewTreeMakeCmd())
	cmd.AddCommand(tree.NewTreeCompareCmd())

	return cmd
}
