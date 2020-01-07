package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/chez-shanpu/reposiTree/pkg/tree"
	"github.com/spf13/cobra"
)

func NewTreeCompareCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compare",
		Short: "compare trees based on alignment distance",
		RunE:  compareTree,
	}

	return cmd
}

func compareTree(cmd *cobra.Command, args []string) (err error) {
	var treeFilePaths []string
	var trees [2]*tree.NodeInfo

	treeFilePaths = args
	if len(treeFilePaths) != 2 {
		return errors.New("number of argument is wrong")
	}
	for key := range treeFilePaths {
		trees[key], err = readTreeFile(treeFilePaths[key])
		if err != nil {
			return err
		}
	}
	dist := tree.LayerAlignmentDistanceTotal(trees[0].RootNode, trees[1].RootNode)
	fmt.Printf("Alignment distance between %s and %s is %f",
		trees[0].RepositoryName, trees[1].RepositoryName, dist)
	return nil
}

func readTreeFile(filePath string) (*tree.NodeInfo, error) {
	var tree *tree.NodeInfo
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &tree)
	if err != nil {
		return nil, err
	}
	return tree, nil
}
