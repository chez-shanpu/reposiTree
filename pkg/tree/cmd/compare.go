package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"

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
	sRootNode := []*tree.Node{trees[0].RootNode}
	tRootNode := []*tree.Node{trees[1].RootNode}
	dist, err := tree.AlignmentDistance(sRootNode, tRootNode)
	if err != nil {
		return err
	}
	dist = math.Round(dist*tree.SIGINIGICANT_DIGITS) / tree.SIGINIGICANT_DIGITS
	fmt.Print(dist)
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
