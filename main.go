import (
	"container/heap"
	"fmt"
)

type Edge struct {
	to   int
	cost int
}

type Graph struct {
	n   int
	adj [][]Edge
}

func Dijkstra(graph Graph, start int) []int {
	const INF = 1e9
	dist := make([]int, graph.n)
	for i := 0; i < graph.n; i++ {
		dist[i] = INF
	}
	dist[start] = 0
	q := make(PriorityQueue, 0)
	heap.Init(&q)
	heap.Push(&q, &Item{start, 0})
	for q.Len() > 0 {
		item := heap.Pop(&q).(*Item)
		v := item.value
		if dist[v] < item.priority {
			continue
		}
		for _, e := range graph.adj[v] {
			if newCost := dist[v] + e.cost; newCost < dist[e.to] {
				dist[e.to] = newCost
				heap.Push(&q, &Item{e.to, newCost})
			}
		}
	}
	return dist
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func main() {
	graph := Graph{
		n: 6,
		adj: [][]Edge{
			{{1, 2}, {2, 5}},
			{{0, 2}, {2, 4}, {3, 6}},
			{{0, 5}, {1, 4}, {3, 2}, {4, 3}},
			{{1, 6}, {2, 2}, {4, 7}},
			{{2, 3}, {3, 7}, {5, 1}},
			{{4, 1}},
		},
	}
	start := 0
	dist := Dijkstra(graph, start)
	for i, d := range dist {
		fmt.Printf("%d까지의 최단 거리: %d\n", i, d)
	}
}
