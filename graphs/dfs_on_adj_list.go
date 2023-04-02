package graphs

type WeightedAdjList [][]GraphEdge

type GraphEdge struct {
	to     int
	weight int
}

func Dfs(graph WeightedAdjList, source int, needle int) []int {
	seen := make([]bool, len(graph))
	path := make([]int, len(graph))

	walk(graph, source, needle, seen, path)
	return path
}

func walk(graph WeightedAdjList, curr int, needle int, seen []bool, path []int) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true
	path = append(path, curr)

	if curr == needle {
		return true
	}

	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]

		if walk(graph, edge.to, needle, seen, path) {
			return true
		}
	}
	path = path[1:]
	return false
}
