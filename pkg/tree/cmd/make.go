package cmd

import (
	"log"
	"path/filepath"
	"strconv"
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
	flags.StringP("max-depth", "", "", "Max tree depth")

	// bind flags
	_ = viper.BindPFlag("tree.make.repopath", flags.Lookup("repository-path"))
	_ = viper.BindPFlag("tree.make.language", flags.Lookup("language"))
	_ = viper.BindPFlag("tree.make.output", flags.Lookup("output"))
	_ = viper.BindPFlag("tree.make.max_depth", flags.Lookup("max-depth"))

	// required
	_ = cmd.MarkFlagRequired("repository-path")
	_ = cmd.MarkFlagRequired("language")
	_ = cmd.MarkFlagRequired("output")
	_ = cmd.MarkFlagRequired("max-depth")

	return cmd
}

func makeTree(cmd *cobra.Command, args []string) error {
	createdDate := time.Now().String()
	language := viper.GetString("tree.make.language")
	repoRootPath := viper.GetString("tree.make.repopath")
	maxDepth, err := strconv.Atoi(viper.GetString("tree.make.max_depth"))
	if err != nil {
		return err
	}
	_, repoName := filepath.Split(repoRootPath)

	rootNode, err := tree.MakeNode(repoRootPath, repoName, 1, maxDepth, language)
	if err != nil {
		return err
	}
	nodeInfo := tree.NodeInfo{
		RootNode:       rootNode,
		RepositoryName: repoName,
		Language:       viper.GetString("tree.make.language"),
		CreatedDate:    createdDate,
	}

	// output to the .json
	outputName := viper.GetString("tree.make.output")
	utils.OutputJson(outputName, nodeInfo)
	log.Printf("Output to %v completed", outputName)
	return nil
}
