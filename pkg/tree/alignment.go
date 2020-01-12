package tree

func AlignmentDistance(sNodes, tNodes []*Node) (sum float64, err error) {
	if len(sNodes) > len(tNodes) {
		sNodes, tNodes = SwapNodeSlice(sNodes, tNodes)
	}

	res, tNodes := optNodesDiff(sNodes, tNodes)
	sum += res

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
func optNodesDiff(sNodes, tNodes []*Node) (float64, []*Node) {
	var res float64
	if sNodes == nil && tNodes == nil {
		return 0, tNodes
	} else if sNodes == nil {
		for i, _ := range tNodes {
			res += NodeDataDiff(nil, tNodes[i])
		}
		return res, tNodes
	} else if tNodes == nil {
		for i, _ := range sNodes {
			res += NodeDataDiff(sNodes[i], nil)
		}
		return res, tNodes
	}

	sNodesLength := len(sNodes)
	tNodesLength := len(tNodes)
	s := tNodesLength * 2
	t := s + 1

	g := Graph{
		NodeNum: t + 1,
	}

	for i, _ := range tNodes {
		for j, _ := range tNodes {
			var cost float64
			if i < sNodesLength {
				cost = NodeDataDiff(sNodes[i], tNodes[j])
			} else {
				cost = NodeDataDiff(nil, tNodes[j])
			}
			g.AddEdge(i, j+tNodesLength, 1, cost)
		}
	}

	for i, _ := range tNodes {
		g.AddEdge(s, i, 1, 0)
		g.AddEdge(i+tNodesLength, t, 1, 0)
	}

	res = MinCostFlow(&g, s, t, tNodesLength)

	tNewNodes := fixNodePointer(sNodes, tNodes, &g)
	return res, tNewNodes
}

func fixNodePointer(sNodes, tNodes []*Node, g *Graph) []*Node {
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
