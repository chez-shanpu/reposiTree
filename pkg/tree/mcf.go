package tree

import (
	"fmt"
	"math"
)

// Maximum number of network nodes
const MaxV = 10000
const INF = 10000000

type Edge struct {
	Rev  int
	From int
	To   int
	Cap  int
	ICap int
	Cost float64
}

type Graph struct {
	NodeNum int
	Nodes   [MaxV]McfNode
}

type McfNode struct {
	Edges []Edge
}

func (g *Graph) ReEdge(e *Edge) *Edge {
	if e.From != e.To {
		return &g.Nodes[e.To].Edges[e.Rev]
	} else {
		return &g.Nodes[e.To].Edges[e.Rev+1]
	}
}

func (g *Graph) AddEdge(from, to, cap int, cost float64) {
	g.Nodes[from].Edges =
		append(
			g.Nodes[from].Edges,
			Edge{
				Rev:  len(g.Nodes[to].Edges),
				From: from,
				To:   to,
				Cap:  cap,
				ICap: cap,
				Cost: cost,
			})
	g.Nodes[to].Edges =
		append(
			g.Nodes[to].Edges,
			Edge{
				Rev:  len(g.Nodes[from].Edges) - 1,
				From: to,
				To:   from,
				Cap:  0,
				ICap: 0,
				Cost: -cost,
			})
}

func MinCostFlow(g *Graph, s int, t int, inif int) float64 {
	var preNode [MaxV]int
	var preEdge [MaxV]int

	dist := make([]float64, MaxV)
	res := 0.0
	f := inif

	for f > 0 {
		dist, _ = fill(dist, INF, 0, g.NodeNum)
		dist[s] = 0
		for true {
			update := false
			for node := 0; node < g.NodeNum; node++ {
				if dist[node] == INF {
					continue
				}
				for i, _ := range g.Nodes[node].Edges {
					e := &(g.Nodes[node].Edges[i])
					if (e.Cap > 0) && (dist[e.To] > dist[node]+e.Cost) {
						dist[e.To] = dist[node] + e.Cost
						preNode[e.To] = node
						preEdge[e.To] = i
						update = true
					}
				}
			}
			if update == false {
				break
			}
		}
		if dist[t] == INF {
			return 0
		}
		d := f
		for node := t; node != s; node = preNode[node] {
			d = int(math.Min(float64(d), float64(g.Nodes[preNode[node]].Edges[preEdge[node]].Cap)))
		}
		f -= d
		res += dist[t] * float64(d)
		for node := t; node != s; node = preNode[node] {
			e := &(g.Nodes[preNode[node]].Edges[preEdge[node]])
			re := g.ReEdge(e)
			e.Cap -= d
			re.Cap += d
		}
	}
	return res
}

func fill(slice []float64, val float64, start, end int) ([]float64, error) {
	if len(slice) < start || len(slice) < end {
		return nil, fmt.Errorf("error")
	}
	for i := start; i < end; i++ {
		slice[i] = val
	}
	return slice, nil
}
