import (
	"container/list"
)

type graph struct {
	nodes    int
	adj      []*list.List
	zeros    *list.List
	indegree []int
}

func NewGraph(numCourses int, prerequisites [][]int) *graph {
	g := &graph{
		nodes: numCourses,
	}
	g.adj = make([]*list.List, g.nodes)
	for i := 0; i < g.nodes; i++ {
		g.adj[i] = list.New()
	}
	g.indegree = make([]int, g.nodes)
	for i := 0; i < g.nodes; i++ {
		g.indegree[i] = 0
	}

	for i := 0; i < len(prerequisites); i++ {
		g.adj[prerequisites[i][1]].PushBack(prerequisites[i][0])
		g.indegree[prerequisites[i][0]] = g.indegree[prerequisites[i][0]] + 1
	}

	g.zeros = list.New()
	for i := 0; i < g.nodes; i++ {
		if g.indegree[i] == 0 {
			g.zeros.PushBack(i)
		}
	}
	return g
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	count := 0
	result := []int{}

	g := NewGraph(numCourses, prerequisites)
	for {
		if g.zeros.Len() == 0 {
			break
		}
		node := g.zeros.Front().Value.(int)
		count = count + 1
		g.zeros.Remove(g.zeros.Front())

		result = append(result, node)

		for e := g.adj[node].Front(); e != nil; e = e.Next() {
			val := e.Value.(int)
			g.indegree[val] = g.indegree[val] - 1
			if g.indegree[val] == 0 {
				g.zeros.PushBack(val)
			}
		}
	}
	if count < g.nodes {
		return []int{}
	}
	return result
}





