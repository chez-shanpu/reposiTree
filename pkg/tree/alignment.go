package tree

import (
	"math"
)

const SIGINIGICANT_DIGITS = 100000

func AlignmentDistance(sNodes, tNodes []*Node) (sum float64, err error) {
	if len(sNodes) > len(tNodes) {
		sNodes, tNodes = SwapNodeSlice(sNodes, tNodes)
	}

	if sNodes == nil && tNodes == nil {
		return 0, nil
	} else if sNodes == nil {
		for i, _ := range tNodes {
			sum += NodeDataDiff(nil, tNodes[i])
		}
	} else if tNodes == nil {
		for i, _ := range sNodes {
			sum += NodeDataDiff(sNodes[i], nil)
		}
	} else {
		var diff int
		diff, tNodes = optNodesDiff(sNodes, tNodes)
		sum += float64(diff) / float64(SIGINIGICANT_DIGITS)
	}

	for i, _ := range tNodes {
		var dist float64
		if i < len(sNodes) {
			dist, err = AlignmentDistance(sNodes[i].ChildNodes, tNodes[i].ChildNodes)
		} else {
			dist, err = AlignmentDistance(nil, tNodes[i].ChildNodes)
		}

		if err != nil {
			return 0, err
		} else {
			sum += dist
		}
	}
	return sum, nil
}

// Must sourceLayer length < targetLayer length
func optNodesDiff(sNodes, tNodes []*Node) (int, []*Node) {
	var res int

	sNodesLength := len(sNodes)
	tNodesLength := len(tNodes)
	s := tNodesLength * 2
	t := s + 1

	g := Graph{
		NodeNum: t + 1,
	}

	for i, _ := range tNodes {
		for j, _ := range tNodes {
			var d float64
			if i < sNodesLength {
				d = NodeDataDiff(sNodes[i], tNodes[j])
			} else {
				d = NodeDataDiff(nil, tNodes[j])
			}
			cost := int(math.Round(d * SIGINIGICANT_DIGITS))
			g.AddEdge(i, j+tNodesLength, 1, cost)
		}
	}

	for i, _ := range tNodes {
		g.AddEdge(s, i, 1, 0)
		g.AddEdge(i+tNodesLength, t, 1, 0)
	}

	res = MinCostFlow(&g, s, t, tNodesLength)
	tNewNodes := fixNodePointer(tNodes, &g)
	return res, tNewNodes
}

func fixNodePointer(tNodes []*Node, g *Graph) []*Node {
	var resNodes []*Node

	for i := range tNodes {
		for _, j := range g.Nodes[i].Edges {
			if j.ICap == 1 && j.Cap == 0 {
				resNodes = append(resNodes, tNodes[j.To-len(tNodes)])
				break
			}
		}
	}
	return resNodes
}
