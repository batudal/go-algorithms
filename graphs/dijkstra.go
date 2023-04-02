package graphs

import (
	"math"
)

func DijkstraList(source int, sink int, arr WeightedAdjList) []int {
	seen := make([]bool, len(arr))
	prev := make([]int, len(arr))
	dists := make([]float64, len(arr))
	for i := 0; i < len(dists); i++ {
		dists[i] = math.Inf(1)
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]

			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + float64(edge.weight)
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	out := make([]int, 0)
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

func hasUnvisited(seen []bool, dists []float64) bool {
	for i := 0; i < len(seen); i++ {
		if !seen[i] && dists[i] < math.Inf(1) {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []float64) int {
	idx := -1
	lowest_distance := math.Inf(1)

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}
		if lowest_distance > dists[i] {
			lowest_distance = dists[i]
			idx = i
		}
	}
	return idx
}
