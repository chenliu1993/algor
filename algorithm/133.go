package algorithm

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{}
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		cloneNode := &Node{node.Val, []*Node{}}
		visited[node] = cloneNode
		for _, val := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(val))
		}
		return cloneNode
	}
	return cg(node)
}
