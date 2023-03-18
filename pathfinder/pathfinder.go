package pathfinder

type Point struct {
	x int
	y int
}

func walk(maze []string, wall rune, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}
	row := []rune(maze[curr.y])
	if row[curr.x] == wall {
		return false
	}
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}
	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	*path = append(*path, curr)
	for i := range dir {
		if walk(maze, wall, Point{
			x: curr.x + dir[i][0],
			y: curr.y + dir[i][1],
		}, end, seen, path) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]

	return false
}

func Solve(maze []string, wall rune, start Point, end Point) []Point {
	seen := make([][]bool, 0)
	path := make([]Point, 0)
	walk(maze, wall, start, end, seen, &path)
	return path
}
