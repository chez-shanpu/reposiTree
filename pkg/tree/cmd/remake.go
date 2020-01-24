package cmd

import (
	"fmt"
	"github.com/chez-shanpu/reposiTree/pkg/tree"
	"github.com/chez-shanpu/reposiTree/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func NewTreeRemakeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remake",
		Short: "remaking a repository from a tree data",
		RunE:  remakeTree,
	}

	// flags
	flags := cmd.Flags()
	flags.StringP("treedata-path", "p", "", "A path to the tree data")
	flags.StringP("output", "o", "", "path to root dir that this program output to")

	// bind flags
	_ = viper.BindPFlag("tree.remake.treedata_path", flags.Lookup("treedata-path"))
	_ = viper.BindPFlag("tree.remake.output", flags.Lookup("output"))

	// required
	_ = cmd.MarkFlagRequired("treedata-path")
	_ = cmd.MarkFlagRequired("output")

	return cmd
}

func remakeTree(cmd *cobra.Command, args []string) error {
	treePath := viper.GetString("tree.remake.treedata_path")
	outPath := viper.GetString("tree.remake.output")

	var repo tree.NodeInfo
	err := utils.ReadJson(treePath, &repo)
	if err != nil {
		return err
	}
	rootPath := filepath.Join(outPath, repo.RepositoryName)
	err = os.Mkdir(rootPath, 0755)
	if err != nil {
		return err
	}

	fmt.Printf("Output %s's abstract dirs stared!", repo.RepositoryName)
	err = tree.OutputAbstractRepository([]*tree.Node{repo.RootNode}, rootPath)
	if err != nil {
		return err
	}

	fmt.Printf("Output %s's abstract dirs completed!", repo.RepositoryName)
	return nil
}
