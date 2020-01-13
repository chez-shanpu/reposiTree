package tree

import (
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
)

type Node struct {
	Vector        [MAX_FILETYPE]float64 `json:"data"`
	DirectoryName string                `json:"directory_name"`
	ChildNodes    []*Node               `json:"child_nodes"`
}

type NodeInfo struct {
	RootNode       *Node  `json:"root_node"`
	RepositoryName string `json:"repository_name"`
	Language       string `json:"language"`
	CreatedDate    string `json:"created_date"`
}

func SwapNodeSlice(sNodes, tNodes []*Node) ([]*Node, []*Node) {
	tmp := sNodes
	sNodes = tNodes
	tNodes = tmp
	return sNodes, tNodes
}

func NodeDataDiff(sNode *Node, tNode *Node) (res float64) {
	var sum float64

	if sNode == nil && tNode == nil {
		return 0
	} else if sNode == nil {
		for i := range tNode.Vector {
			sum += math.Pow(tNode.Vector[i], 2)
		}
	} else if tNode == nil {
		for i := range sNode.Vector {
			sum += math.Pow(sNode.Vector[i], 2)
		}
	} else {
		for i := range sNode.Vector {
			sum += math.Pow(sNode.Vector[i]-tNode.Vector[i], 2)
		}
	}
	res = math.Sqrt(sum)
	res = math.Round(res*SIGINIGICANT_DIGITS) / SIGINIGICANT_DIGITS
	return res
}

func MakeNode(dirPath string, dirName string, depth int, language string, pNode *Node) (*Node, error) {
	var subDirs []os.FileInfo

	n := new(Node)
	n.DirectoryName = dirName
	n.Vector = [MAX_FILETYPE]float64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			subDirs = append(subDirs, file)
		} else {
			nodeDataIndex, err := FileClassifier(file.Name(), language)
			if err != nil {
				return nil, err
			}
			if n.Vector[nodeDataIndex] == 0 {
				n.Vector[nodeDataIndex] = 1
			}
		}
	}
	for key := range n.Vector {
		if pNode == nil {
			n.Vector[key] = n.Vector[key] / float64(depth)
		} else {
			n.Vector[key] = (n.Vector[key] + n.Vector[key]) / float64(depth)
		}
	}

	for _, subDir := range subDirs {
		childNode, err := MakeNode(filepath.Join(dirPath, subDir.Name()), filepath.Join(dirName, subDir.Name()), depth+1, language, n)
		if err != nil {
			return nil, err
		}
		n.ChildNodes = append(n.ChildNodes, childNode)
	}
	return n, nil
}
