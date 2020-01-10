package tree

import (
	"io/ioutil"
	"math"
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

func (n *Node) NodeNumSum() int {
	cnt := 0
	for node := n; node != nil; node = node.NextNode {
		cnt++
	}
	return cnt
}

func (n *Node) NodeDataSum() float64 {
	sum := 0.0
	if n == nil {
		return sum
	}
	for _, val := range n.Data {
		sum += val
	}
	return sum
}

func (n *Node) LayerLength() int {
	length := 0
	for node := n; node != nil; node = node.NextNode {
		length++
	}
	return length
}

func (n *Node) GetNode(index int) *Node {
	node := n
	for i := 0; i < index; i++ {
		if node != nil {
			node = node.NextNode
		} else {
			return nil
		}
	}
	return node
}

func NodeDataDiff(sNode *Node, tNode *Node) float64 {
	res := 0.0
	if sNode == nil {
		res = tNode.NodeDataSum()
	} else if tNode == nil {
		res = sNode.NodeDataSum()
	} else {
		for i := range sNode.Data {
			res += math.Abs(sNode.Data[i] - tNode.Data[i])
		}
	}
	return res
}

func MakeLayer(dirPaths []string, depth int, parentNode *Node, language string) (*Node, error) {
	var rightmostNode *Node
	var leftmostNode *Node

	for _, dirPath := range dirPaths {
		targetNode, err := makeNode(dirPath, depth, parentNode, language)
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

func makeNode(dirPath string, depth int, parentNode *Node, language string) (*Node, error) {
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
			nodeDataIndex, err = FileClassifier(file.Name(), language)
			if err != nil {
				return nil, err
			}
			if node.Data[nodeDataIndex] == 0 {
				node.Data[nodeDataIndex] = 1
			}
		}
	}
	for key := range node.Data {
		if parentNode == nil {
			node.Data[key] = node.Data[key] / float64(depth)
		} else {
			node.Data[key] = (node.Data[key] + parentNode.Data[key]) / float64(depth)
		}
	}
	if subDirPaths != nil {
		node.ChildNode, err = MakeLayer(subDirPaths, depth+1, &node, language)
		if err != nil {
			return nil, err
		}
	}
	return &node, nil
}
