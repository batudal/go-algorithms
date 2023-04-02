package graphs

type WeightedAdjMatrix [][]int

func Bfs(graph WeightedAdjMatrix, source int, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}

	seen[source] = true
	q := make([]int, 1)
	q[0] = source

	for len(q) > 0 {
		curr := q[0]
		q = append(q[1:])
		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q = append(q, i)
		}
	}

	curr := needle
	out := make([]int, 0)

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) > 0 {
		out = append(out, source)

		// reverse list
		for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}
		return out
	}
	return nil
}
