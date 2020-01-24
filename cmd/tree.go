package cmd

import (
	treecmd "github.com/chez-shanpu/reposiTree/pkg/tree/cmd"
	"github.com/spf13/cobra"
)

func NewTreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tree",
		Short: "repository tree",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			return err
		},
	}

	cmd.AddCommand(treecmd.NewTreeMakeCmd())
	cmd.AddCommand(treecmd.NewTreeCompareCmd())
	cmd.AddCommand(treecmd.NewTreeRemakeCmd())

	return cmd
}
