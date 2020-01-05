package tree

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chez-shanpu/reposiTree/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

func init() {
	cmd.NewTreeCmd().AddCommand(NewTreeCompareCmd())
}

func NewTreeCompareCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compare",
		Short: "compare trees based on alignment distance",
		RunE:  compareTree,
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
