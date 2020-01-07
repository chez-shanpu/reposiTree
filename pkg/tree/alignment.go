package tree

import (
	"math"
)

func layerAlignmentDistanceTotal(sourceLayerRootNode *Node, targetLayerRootNode *Node) (sum float64) {
	sum = 0
	sNode := sourceLayerRootNode
	tNode := targetLayerRootNode
	sum += alignmentDistance(sNode, tNode)
	for sNode != nil || tNode != nil {
		if (sNode != nil && sNode.ChildNode != nil) || (tNode != nil && tNode.ChildNode != nil) {
			if sNode == nil {
				sum += layerAlignmentDistanceTotal(nil, tNode.ChildNode)
			} else if tNode == nil {
				sum += layerAlignmentDistanceTotal(sNode.ChildNode, nil)
			} else {
				sum += layerAlignmentDistanceTotal(sNode.ChildNode, tNode.ChildNode)
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
	sourceLayerLength := layerLength(sourceLayerRootNode)
	targetLayerLength := layerLength(targetLayerRootNode)

	// swap
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

func layerLength(leftmostNode *Node) (length int) {
	length = 0
	for node := leftmostNode; node != nil; node = node.NextNode {
		length++
	}
	return
}

// Always sourceLayer length < targetLayer length
func optNodesDiff(sourceLayerRootNode *Node, targetLayerRootNode *Node) float64 {
	res := 0.0

	if sourceLayerRootNode == nil && targetLayerRootNode == nil {
		return 0
	} else if sourceLayerRootNode == nil {
		for n := targetLayerRootNode; n != nil; n = n.NextNode {
			res += nodeDataSum(n)
		}
		return res
	}
	sNode := sourceLayerRootNode
	tNode := targetLayerRootNode
	sNodeNumSum := nodeNumSum(sourceLayerRootNode)
	tNodeNumSum := nodeNumSum(targetLayerRootNode)
	s := sNodeNumSum + tNodeNumSum
	t := s + 1

	g := Graph{
		NodeNum: t,
	}

	for i := 0; i < sNodeNumSum; i++ {

		for j := 0; j < tNodeNumSum; j++ {

			cost := nodeDataDiff(sNode, tNode)
			g.AddEdge(i, j+sNodeNumSum, 1, cost)
			tNode = tNode.NextNode
		}
		sNode = sNode.NextNode
		tNode = targetLayerRootNode
	}

	for i := 0; i < sNodeNumSum; i++ {
		g.AddEdge(s, i, 1, 0)
	}

	for j := 0; j < tNodeNumSum; j++ {
		g.AddEdge(j+sNodeNumSum, t, 1, 0)
	}

	res = MinCostFlow(&g, s, t, sNodeNumSum)

	n := targetLayerRootNode
	for j := 0; j < tNodeNumSum; j++ {
		e := g.Nodes[t].Edges[j]
		if e.Cap != 1 {
			res += nodeDataSum(n)
		}
		n = n.NextNode
	}
	return res
}

func nodeNumSum(node *Node) (cnt int) {
	cnt = 0
	for node != nil {
		cnt++
		node = node.NextNode
	}
	return
}

func nodeDataSum(node *Node) (sum float64) {
	sum = 0
	if node == nil {
		sum = 0
		return
	}
	for _, val := range node.Data {
		sum += val
	}
	return
}

func nodeDataDiff(sNode *Node, tNode *Node) float64 {
	res := 0.0
	for i := range sNode.Data {
		res += math.Abs(sNode.Data[i] - tNode.Data[i])
	}
	return res
}
