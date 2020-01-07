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
	// TODO このNodeNumはnilじゃないNodesのインデックスの最大値が入る（つまりNodeの最大値-1） あとで名前変えたい
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

// sにはネットワークの始点となるNodeの番号，tにはネットワークの終点となるNodeの番号がそれぞれ入る
func MinCostFlow(g *Graph, s int, t int, inif int) float64 {
	var preNode [MaxV]int
	var preEdge [MaxV]int

	// distにはNode間の距離（エッジのコストに基づいて算出される）が入る 初期値はINF
	dist := make([]float64, MaxV)

	res := 0.0
	// fには最終的に入る終点へのエッジの本数が入る
	f := inif

	for f > 0 {
		dist, _ := fill(dist, INF, 0, g.NodeNum)
		dist[s] = 0
		for true {
			update := false
			// それぞれのNodeについて処理を行う
			for node := 0; node < g.NodeNum; node++ {
				if dist[node] == INF {
					continue
				}
				for i, _ := range g.Nodes[node].Edges {
					e := &(g.Nodes[node].Edges[i])
					// そのエッジを通ることができ，かつ現在わかってるその行先Nodeへの距離よりも現在いるNodeからそこへ向かうコスト（距離）が小さい場合
					if e.Cap > 0 && dist[e.To] > dist[node]+e.Cost {
						dist[e.To] = dist[node] + e.Cost
						// そのNodeへの経路においてそのNodeの一個前は"node"である
						preNode[e.To] = node
						// そのNodeへの経路においてそのNodeへのエッジは"i"である
						// ここiでいいか疑問に思ったけど"node"と組み合わせれば一意に定められるか
						preEdge[e.To] = i
						update = true
					}
				}
			}
			// いまのグラフ内でupdateするものがなければループを抜ける
			if update == false {
				break
			}
		}

		// t（終点）に到達するNodeがなければ距離は0として返す
		if dist[t] == INF {
			return 0
		}

		d := f

		// 終点から始点に向かって辿っていく
		for node := t; node != s; node = preNode[node] {
			// dにはdと注目nodeへのそのときの最適エッジの容量のうちの最小値が入る
			d = int(math.Min(float64(d), float64(g.Nodes[preNode[node]].Edges[preEdge[node]].Cap)))
		}
		// fからdを引く（dだけ経路が確定したため？）
		f -= d

		// resに確定したtへの距離を加算する
		res += dist[t] * float64(d)

		// 終点から始点に向かって辿っていく
		// 始点から終点方向への容量をdだけ減らして，終点から始点方向への容量をdだけ増やす
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
		return nil, fmt.Errorf("Error")
	}
	for i := start; i <= end; i++ {
		slice[i] = val
	}
	return slice, nil
}
