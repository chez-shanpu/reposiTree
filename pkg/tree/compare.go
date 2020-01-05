package tree

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
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
	var trees [2]*NodeInfo

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
	dist := layerAlignmentDistanceTotal(trees[0].RootNode, trees[1].RootNode)
	fmt.Printf("Alignment distance between %s and %s is %f",
		trees[0].RepositoryName, trees[1].RepositoryName, dist)
	return nil
}

func readTreeFile(filePath string) (*NodeInfo, error) {
	var tree *NodeInfo
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, tree)
	if err != nil {
		return nil, err
	}
	return tree, nil
}
