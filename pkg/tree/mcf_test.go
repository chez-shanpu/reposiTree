package tree_test

import (
	"github.com/chez-shanpu/reposiTree/pkg/tree"
	"log"
	"testing"
)

func TestMinCostFlow(t *testing.T) {
	expectRes := -178.0

	costs := [5][10]float64{{8, 9, 12, 16, 21, 2, 12, 15, 20, 3},
		{7, 12, 5, 14, 13, 10, 16, 12, 18, 11},
		{19, 2, 14, 7, 16, 5, 8, 7, 1, 4},
		{3, 7, 8, 2, 9, 1, 4, 3, 1, 2},
		{20, 27, 33, 39, 26, 21, 15, 30, 28, 19}}
	sNodeNum := 5
	tNodeNum := 10
	sNode := sNodeNum + tNodeNum
	tNode := sNode + 1

	g := tree.Graph{
		NodeNum: sNodeNum + tNodeNum + 2,
		Nodes:   [tree.MaxV]tree.McfNode{},
	}

	for i := 0; i < sNodeNum; i++ {
		for j := 0; j < tNodeNum; j++ {
			g.AddEdge(i, j+sNodeNum, 1, -costs[i][j])
		}
	}

	for i := 0; i < sNodeNum; i++ {
		g.AddEdge(sNode, i, 2, 0)
	}

	for j := 0; j < tNodeNum; j++ {
		g.AddEdge(j+sNodeNum, tNode, 1, 0)
	}

	res := tree.MinCostFlow(&g, sNode, tNode, tNodeNum)
	//TODO debug
	log.Printf("graph: %v", g)
	if res != expectRes {
		t.Errorf("Return: %v Expected: %v", res, expectRes)
	}
}
