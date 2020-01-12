package tree

func LayerAlignmentDistanceTotal(sourceLayerRootNode *Node, targetLayerRootNode *Node) (sum float64) {
	sum = 0
	sNode := sourceLayerRootNode
	tNode := targetLayerRootNode

	sum += alignmentDistance(sNode, tNode)

	for sNode != nil || tNode != nil {
		if (sNode != nil && sNode.ChildNode != nil) || (tNode != nil && tNode.ChildNode != nil) {
			if sNode == nil {
				sum += LayerAlignmentDistanceTotal(nil, tNode.ChildNode)
			} else if tNode == nil {
				sum += LayerAlignmentDistanceTotal(sNode.ChildNode, nil)
			} else {
				sum += LayerAlignmentDistanceTotal(sNode.ChildNode, tNode.ChildNode)
			}
		}
		if sNode != nil {
			sNode = sNode.NextNode
		}
		if tNode != nil {
			tNode = tNode.NextNode
		}
	}
	return
}

// Calculate the total alignment distance for that layer
func alignmentDistance(sourceLayerRootNode *Node, targetLayerRootNode *Node) float64 {
	sourceLayerLength := sourceLayerRootNode.LayerLength()
	targetLayerLength := targetLayerRootNode.LayerLength()

	// swap
	// make sourceLayer length < targetLayer length
	if sourceLayerLength > targetLayerLength {
		tmpNode := sourceLayerRootNode
		sourceLayerRootNode = targetLayerRootNode
		targetLayerRootNode = tmpNode
		tmpLayerLength := sourceLayerLength
		sourceLayerLength = targetLayerLength
		targetLayerLength = tmpLayerLength
	}

	dist := optNodesDiff(sourceLayerRootNode, targetLayerRootNode)
	return dist
}

// Must sourceLayer length < targetLayer length
func optNodesDiff(sourceLayerRootNode *Node, targetLayerRootNode *Node) float64 {
	res := 0.0

	if sourceLayerRootNode == nil && targetLayerRootNode == nil {
		return 0
	} else if sourceLayerRootNode == nil {
		for n := targetLayerRootNode; n != nil; n = n.NextNode {
			res += n.NodeDataSum()
		}
		return res
	}
	sNodeNumSum := sourceLayerRootNode.NodeNumSum()
	tNodeNumSum := targetLayerRootNode.NodeNumSum()
	s := sNodeNumSum + tNodeNumSum
	t := s + 1

	g := Graph{
		NodeNum: t + 1,
	}

	sNode := sourceLayerRootNode
	tNode := targetLayerRootNode
	for i := 0; i < tNodeNumSum; i++ {
		for j := 0; j < tNodeNumSum; j++ {
			cost := NodeDataDiff(sNode, tNode)
			g.AddEdge(i, j+sNodeNumSum, 1, cost)
			tNode = tNode.NextNode
		}
		if sNode != nil {
			sNode = sNode.NextNode
		}
		tNode = targetLayerRootNode
	}

	for i := 0; i < tNodeNumSum; i++ {
		g.AddEdge(s, i, 1, 0)
		g.AddEdge(i+sNodeNumSum, t, 1, 0)
	}

	res = MinCostFlow(&g, s, t, tNodeNumSum)

	*targetLayerRootNode = fixNodePointer(sourceLayerRootNode, targetLayerRootNode, &g)
	return res
}

func fixNodePointer(sNode, tNode *Node, g *Graph) Node {
	var resRootNode Node
	var tmpNodePointer *Node

	for i := 0; i < tNode.LayerLength(); i++ {
		for _, j := range g.Nodes[i].Edges {
			if j.ICap == 1 && j.Cap == 0 {
				if resRootNode == *new(Node) {
					resRootNode = tNode.GetNode(j.To - tNode.LayerLength())
					tmpNodePointer = &resRootNode
				} else {
					tmpNode := tNode.GetNode(j.To - tNode.LayerLength())
					tmpNodePointer.NextNode = &tmpNode
					tmpNodePointer = &tmpNode
					tmpNodePointer.NextNode = nil
				}
				break
			}
		}
	}
	return resRootNode
}
