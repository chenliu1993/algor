
type graph struct {
	nodes int
	adj   map[string][]string
	trace []string
}

func NewGraph(tickets [][]string) *graph {
	g := &graph{}
	g.trace = []string{}
	g.adj = make(map[string][]string)

	for i := 0; i < len(tickets); i++ {
		if _, ok := g.adj[tickets[i][0]]; !ok {
			g.adj[tickets[i][0]] = []string{}
		}
		g.adj[tickets[i][0]] = append(g.adj[tickets[i][0]], tickets[i][1])
	}

	g.nodes = len(g.adj)
	for k, _ := range g.adj {
		sort.Strings(g.adj[k])
	}

	return g

}

func (g *graph) dfs(node string) {
	for {
		if _, ok := g.adj[node]; !ok || len(g.adj[node]) == 0 {
			break
		}
		current := g.adj[node][0]
		g.adj[node] = g.adj[node][1:]
		g.dfs(current)
	}
	g.trace = append(g.trace, node)
}

func findItinerary(tickets [][]string) []string {
	g := NewGraph(tickets)
	g.dfs("JFK")

	for i := 0; i < len(g.trace)/2; i++ {
		g.trace[i], g.trace[len(g.trace)-1-i] = g.trace[len(g.trace)-1-i], g.trace[i]
	}
	return g.trace
}

