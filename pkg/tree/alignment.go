package tree

import (
	"log"
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

	log.Printf("Alignment Distance between sourceLayerRootNode and targetLayerRootNode is %f", dist)
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
	sNodeSum := nodeNumSum(sourceLayerRootNode)
	tNodeSum := nodeNumSum(targetLayerRootNode)
	s := sNodeSum + tNodeSum
	t := s + 1

	g := Graph{
		NodeNum: t,
	}

	for i := 0; i < sNodeSum; i++ {
		for j := 0; j < tNodeSum; j++ {
			cost := math.Abs(nodeDataSum(sNode) - nodeDataSum(tNode))
			g.AddEdge(i, j+sNodeSum, 1, cost)
			if tNode != nil {
				tNode = tNode.NextNode
			}
		}
		if sNode != nil {
			sNode = sNode.NextNode
		}
	}

	for i := 0; i < sNodeSum; i++ {
		g.AddEdge(s, i, 1, 0)
	}

	for j := 0; j < tNodeSum; j++ {
		g.AddEdge(j+sNodeSum, t, 1, 0)
	}

	res = MinCostFlow(&g, s, t, tNodeSum)

	n := targetLayerRootNode
	for j := 0; j < tNodeSum; j++ {
		e := g.Nodes[t].Edges[j]
		if e.Cap != 1 {
			res += nodeDataSum(n)
		}
		n = n.NextNode
	}

	return res
}

func nodeNumSum(node *Node) (cnt int) {
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
