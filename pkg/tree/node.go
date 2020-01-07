package tree

import (
	"io/ioutil"
	"path/filepath"
)

type Node struct {
	Data          [9]float64 `json:"data"`
	DirectoryName string     `json:"directory_name"`
	NextNode      *Node      `json:"next_node"`
	ChildNode     *Node      `json:"child_node"`
}

type NodeInfo struct {
	RootNode       *Node  `json:"root_node"`
	RepositoryName string `json:"repository_name"`
	Language       string `json:"language"`
	CreatedDate    string `json:"created_date"`
}

func MakeLayer(dirPaths []string, depth int, parentNode *Node) (*Node, error) {
	var rightmostNode *Node
	var leftmostNode *Node

	for _, dirPath := range dirPaths {
		targetNode, err := makeNode(dirPath, depth, parentNode)
		if err != nil {
			return nil, err
		}

		if leftmostNode == nil {
			leftmostNode = targetNode
		} else {
			rightmostNode.NextNode = targetNode
		}
		rightmostNode = targetNode
	}
	return leftmostNode, nil
}

func makeNode(dirPath string, depth int, parentNode *Node) (*Node, error) {
	var node Node
	var nodeDataIndex int
	var subDirPaths []string

	node.DirectoryName = dirPath
	node.Data = [9]float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			subDirPaths = append(subDirPaths, filepath.Join(dirPath, file.Name()))
		} else {
			nodeDataIndex = FileClassifier(file.Name())
			if node.Data[nodeDataIndex] == 0 {
				node.Data[nodeDataIndex] = 1
			}
		}
	}
	for key := range node.Data {
		if parentNode == nil {
			node.Data[key] = node.Data[key] / float64(depth)
		} else {
			node.Data[key] = node.Data[key]/float64(depth) + parentNode.Data[key]
		}
	}
	if subDirPaths != nil {
		node.ChildNode, err = MakeLayer(subDirPaths, depth+1, &node)
		if err != nil {
			return nil, err
		}
	}
	return &node, nil
}
