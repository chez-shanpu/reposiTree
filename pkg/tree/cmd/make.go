package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/chez-shanpu/reposiTree/pkg/tree"
	"github.com/chez-shanpu/reposiTree/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewTreeMakeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make",
		Short: "Making a tree from a repository",
		RunE:  makeTree,
	}

	// flags
	flags := cmd.Flags()
	flags.StringP("repository-path", "p", "", "A path to the repository root")
	flags.StringP("language", "l", "", "Repository's main language")
	flags.StringP("output", "o", "", "output file name")

	// bind flags
	_ = viper.BindPFlag("tree.make.repopath", flags.Lookup("repository-path"))
	_ = viper.BindPFlag("tree.make.language", flags.Lookup("language"))
	_ = viper.BindPFlag("tree.make.output", flags.Lookup("output"))

	// required
	_ = cmd.MarkFlagRequired("repository-path")
	_ = cmd.MarkFlagRequired("language")
	_ = cmd.MarkFlagRequired("output")

	return cmd
}

func makeTree(cmd *cobra.Command, args []string) error {
	createdDate := time.Now().String()
	repoRootPath := viper.GetString("tree.make.repopath")
	_, repositoryName := filepath.Split(repoRootPath)

	rootNode, err := tree.MakeLayer([]string{repoRootPath}, 1, nil)
	if err != nil {
		return err
	}
	nodeInfo := tree.NodeInfo{
		RootNode:       rootNode,
		RepositoryName: repositoryName,
		Language:       viper.GetString("tree.make.language"),
		CreatedDate:    createdDate,
	}

	// output to the .json
	outputName := viper.GetString("tree.make.output")
	utils.OutputJson(outputName, nodeInfo)
	fmt.Printf("Output to %v completed", outputName)

	return nil
}
