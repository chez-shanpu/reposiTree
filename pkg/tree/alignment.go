package tree

import (
	"github.com/chez-shanpu/repo2tree/model"
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
	var dist float64 = 0

	sourceLayerLength := layerLength(sourceLayerRootNode)
	targetLayerLength := layerLength(targetLayerRootNode)
	layerLengthGap := math.Abs(float64(sourceLayerLength - targetLayerLength))
	if sourceLayerLength < targetLayerLength {
		tmpNode := sourceLayerRootNode
		sourceLayerRootNode = targetLayerRootNode
		targetLayerRootNode = tmpNode
		tmpLayerLength := sourceLayerLength
		sourceLayerLength = targetLayerLength
		targetLayerLength = tmpLayerLength
	}
	// TODO ここの比較の部分は最小費用流問題に落とし込んで，距離が最小になるノードの組み合わせを求めてから距離を算出する
	for lg := layerLengthGap; lg > 0; lg-- {
		dist += nodeDataSum(sourceLayerRootNode)
		sourceLayerRootNode = sourceLayerRootNode.NextNode
	}
	for remainLength := targetLayerLength; remainLength > 0; remainLength-- {
		dist += math.Abs(nodeDataSum(sourceLayerRootNode) - nodeDataSum(targetLayerRootNode))
		sourceLayerRootNode = sourceLayerRootNode.NextNode
		targetLayerRootNode = targetLayerRootNode.NextNode
	}
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

func nodeDataSum(node *Node) (sum float64) {
	sum = 0
	if node == nil {
		sum = 0
		return
	}
	for key, val := range node.Data {
		log.Printf("key is %d val is %f", key, val)
		sum += math.Pow(val, math.Pow(10, float64(key)))
		var hoge float64
		hoge=10000000
		log.Printf("sum is %v", sum)
		log.Printf("%f", hoge)
	}
	log.Printf("node %s's sum is %f",node.DirectoryName,sum)
	return
}
