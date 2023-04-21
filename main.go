package main

import (
	"fmt"
)

type Node struct {
	value    int
	children []*Node
	parent   *Node
}

func newNode(value int) *Node {
	return &Node{
		value:    value,
		children: make([]*Node, 0),
	}
}

func (n *Node) querySum() int {
	sum := n.value
	for _, child := range n.children {
		sum += child.querySum()
	}
	return sum
}

func (n *Node) updateValues(w int) {
	for n.parent != nil {
		n.value = n.parent.value
		n = n.parent
	}
	n.value = w
}

func buildTree(values []int, edges [][]int) *Node {
	nodes := make([]*Node, len(values)+1)
	for i, value := range values {
		nodes[i+1] = newNode(value)
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		nodes[u].children = append(nodes[u].children, nodes[v])
		nodes[v].parent = nodes[u]
	}

	return nodes[1]
}

func solution(values []int, edges [][]int, queries [][]int) []int {
	tree := buildTree(values, edges)
	result := make([]int, 0)

	for _, query := range queries {
		u, w := query[0], query[1]

		if w == -1 {
			result = append(result, tree.children[u-2].querySum())
		} else {
			if u == 1 {
				tree.updateValues(w)
			} else {
				tree.children[u-2].updateValues(w)
			}
		}
	}

	return result
}

func main() {
	values := []int{1, 10, 100, 1000, 10000}
	edges := [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}}
	queries := [][]int{{1, -1}, {2, -1}, {3, -1}, {4, -1}, {5, -1}, {4, 1000}, {1, -1}, {2, -1}, {3, -1}, {4, -1}, {5, -1}, {2, 1}, {1, -1}, {2, -1}, {3, -1}, {4, -1}, {5, -1}}

	result := solution(values, edges, queries)
	fmt.Println(result)
}
