package day10

import (
	"fmt"
	"math"
)

/*
--- Day 10: Pipe Maze ---
You use the hang glider to ride the hot air from Desert Island all the way up to the floating metal island.
This island is surprisingly cold and there definitely aren't any thermals to glide on, so you leave your hang
glider behind.

You wander around for a while, but you don't find any people or animals. However, you do occasionally find signposts
labeled "Hot Springs" pointing in a seemingly consistent direction; maybe you can find someone at the hot springs and
ask them where the desert-machine parts are made.

The landscape here is alien; even the flowers and trees are made of metal. As you stop to admire some metal grass,
you notice something metallic scurry away in your peripheral vision and jump into a big pipe! It didn't look like
any animal you've ever seen; if you want a better look, you'll need to get ahead of it.

Scanning the area, you discover that the entire field you're standing on is densely packed with pipes; it was hard
to tell at first because they're the same metallic silver color as the "ground". You make a quick sketch of all of
the surface pipes you can see (your puzzle input).

The pipes are arranged in a two-dimensional grid of tiles:

| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape
the pipe has.
Based on the acoustics of the animal's scurrying, you're confident the pipe that contains the animal is one large,
continuous loop.

For example, here is a square loop of pipe:

.....
.F-7.
.|.|.
.L-J.
.....
If the animal had entered this loop in the northwest corner, the sketch would instead look like this:

.....
.S-7.
.|.|.
.L-J.
.....
In the above diagram, the S tile is still a 90-degree F bend: you can tell because of how the adjacent pipes
connect to it.

Unfortunately, there are also many pipes that aren't connected to the loop! This sketch shows the same loop as above:

-L|F7
7S-7|
L|7||
-L-J|
L|-JF
In the above diagram, you can still figure out which pipes form the main loop: they're the ones connected to S,
pipes those pipes connect to, pipes those pipes connect to, and so on. Every pipe in the main loop connects to its
two neighbors (including S, which will have exactly two pipes connecting to it, and which is assumed to connect back
to those two pipes).

Here is a sketch that contains a slightly more complex main loop:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...
Here's the same example sketch with the extra, non-main-loop pipe tiles also shown:

7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
If you want to get out ahead of the animal, you should find the tile in the loop that is farthest from the starting
position. Because the animal is in the pipe, it doesn't make sense to measure this by direct distance. Instead,
you need to find the tile that would take the longest number of steps along the loop to reach from the starting point -
regardless of which way around the loop the animal went.

In the first example with the square loop:

.....
.S-7.
.|.|.
.L-J.
.....
You can count the distance each tile in the loop is from the starting point like this:

.....
.012.
.1.3.
.234.
.....
In this example, the farthest point from the start is 4 steps away.

Here's the more complex loop again:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...
Here are the distances for each tile on that loop:

..45.
.236.
01.78
14567
23...
Find the single giant loop starting at S. How many steps along the loop does it take to get from the starting
position to the point farthest from the starting position?
*/

type Coord struct {
	R int
	C int
}

func Part1(lines []string) (int, error) {
	grid := parseInput(lines)
	graph, startNode := createGraph(grid)
	furthestDistance := bfs(graph, startNode)
	fmt.Println(furthestDistance)
	return furthestDistance, nil
}

func parseInput(lines []string) [][]string {
	var grid [][]string
	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}
	return grid
}

func createGraph(grid [][]string) (map[Coord][]Coord, Coord) {
	startNode := Coord{}
	graph := map[Coord][]Coord{}
	for r, row := range grid {
		for c, col := range row {
			if col == "S" {
				startNode = Coord{r, c}
			}
			neighbors := getNeighbors(r, c, grid)
			if len(neighbors) > 0 {
				graph[Coord{r, c}] = neighbors
			}
		}
	}
	return graph, startNode
}

func getNeighbors(r, c int, grid [][]string) []Coord {
	val := grid[r][c]

	var neighbors []Coord
	switch val {
	case "S":
		return getStartingNeighbors(r, c, grid)
	case "|":
		return []Coord{
			{r - 1, c},
			{r + 1, c},
		}
	case "-":
		return []Coord{
			{r, c - 1},
			{r, c + 1},
		}
	case "L":
		return []Coord{
			{r - 1, c},
			{r, c + 1},
		}
	case "J":
		return []Coord{
			{r - 1, c},
			{r, c - 1},
		}
	case "7":
		return []Coord{
			{r, c - 1},
			{r + 1, c},
		}
	case "F":
		return []Coord{
			{r + 1, c},
			{r, c + 1},
		}
	case ".":
		// Do nothing since ground doesn't have any connections.
	}

	// Filter out invalid neighbors.
	var validNeighbors []Coord
	for _, neighbor := range neighbors {
		if isValid(neighbor.R, neighbor.C, grid) {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}
	return validNeighbors
}

func isValid(r, c int, grid [][]string) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r])
}

func getStartingNeighbors(r, c int, grid [][]string) []Coord {
	potentialNeighbors := []Coord{
		{r, c + 1},
		{r, c - 1},
		{r + 1, c},
		{r - 1, c},
	}
	// For each neighbor of ours, add the ones if we are one of their neighbors.
	coord := Coord{R: r, C: c}
	var actualNeighbors []Coord
	for _, n := range potentialNeighbors {
		// Skip if this neighbor isn't valid.
		if !isValid(n.R, n.C, grid) {
			continue
		}
		neighborsNeighbors := getNeighbors(n.R, n.C, grid)
		// Check all of our potential neighbors, if the starting location is one of their neighbors,
		// then add it as our neighbor.
		for _, nn := range neighborsNeighbors {
			if nn == coord {
				actualNeighbors = append(actualNeighbors, n)
				break
			}
		}
	}
	return actualNeighbors
}

func bfs(graph map[Coord][]Coord, startNode Coord) int {
	// visited is a mapping of the nodes that we've seen.
	visited := map[Coord]bool{startNode: true}
	distance := 0
	queue := []Coord{startNode}
	for len(queue) > 0 {
		var nextQueue []Coord
		for _, node := range queue {
			visited[node] = true
			for _, neighbor := range graph[node] {
				if _, ok := visited[neighbor]; !ok {
					nextQueue = append(nextQueue, neighbor)
				}
			}
		}
		distance++
		queue = nextQueue
	}
	return distance - 1
}

/*
--- Part Two ---
You quickly reach the farthest point of the loop, but the animal never emerges. Maybe its nest is within the
area enclosed by the loop?

To determine whether it's even worth taking the time to search for such a nest, you should calculate how many
tiles are contained within the loop. For example:

...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
The above loop encloses merely four tiles - the two pairs of . in the southwest and southeast (marked I below).
The middle . tiles (marked O below) are not in the loop. Here is the same loop again with those regions marked:

...........
.S-------7.
.|F-----7|.
.||OOOOO||.
.||OOOOO||.
.|L-7OF-J|.
.|II|O|II|.
.L--JOL--J.
.....O.....
In fact, there doesn't even need to be a full tile path to the outside for tiles to count as outside the loop -
squeezing between pipes is also allowed! Here, I is still within the loop and O is still outside the loop:

..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........
In both of the above examples, 4 tiles are enclosed by the loop.

Here's a larger example:

.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
The above sketch has many random bits of ground, some of which are in the loop (I) and some of which are outside
it (O):

OF----7F7F7F7F-7OOOO
O|F--7||||||||FJOOOO
O||OFJ||||||||L7OOOO
FJL7L7LJLJ||LJIL-7OO
L--JOL7IIILJS7F-7L7O
OOOOF-JIIF7FJ|L7L7L7
OOOOL7IF7||L7|IL7L7|
OOOOO|FJLJ|FJ|F7|OLJ
OOOOFJL-7O||O||||OOO
OOOOL---JOLJOLJLJOOO
In this larger example, 8 tiles are enclosed by the loop.

Any tile that isn't part of the main loop can count as being enclosed by the loop. Here's another example with many
bits of junk pipe lying around that aren't connected to the main loop at all:

FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
Here are just the tiles that are enclosed by the loop marked with I:

FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJIF7FJ-
L---JF-JLJIIIIFJLJJ7
|F|F-JF---7IIIL7L|7|
|FFJF7L7F-JF7IIL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
In this last example, 10 tiles are enclosed by the loop.

Figure out whether you have time to search for the nest by calculating the area within the loop. How many tiles
are enclosed by the loop?
*/

func Part2(lines []string) (int, error) {
	grid := parseInput(lines)
	graph, startNode := createGraph(grid)
	nodes := dfs(graph, startNode)
	fmt.Println(nodes)
	area := shoelaceFormula(nodes)
	interiorPoints := picksTheorem(area, nodes)
	return interiorPoints, nil
}

func dfs(graph map[Coord][]Coord, startNode Coord) []Coord {
	// visited is a mapping of the nodes that we've seen and the distance from the start node.
	visited := map[Coord]bool{}
	stack := []Coord{startNode}
	path := []Coord{}
	for len(stack) > 0 {
		// Take the node off the top of the stack.
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[node]; !ok {
			visited[node] = true
			path = append(path, node)
			for _, neighbor := range graph[node] {
				stack = append(stack, neighbor)
			}
		}
	}
	return path
}

func shoelaceFormula(nodes []Coord) int {
	sum1 := nodes[len(nodes)-1].C * nodes[0].R
	sum2 := nodes[0].C * nodes[len(nodes)-1].R
	for i := 0; i < len(nodes)-1; i++ {
		sum1 += nodes[i].C * nodes[i+1].R
		sum2 += nodes[i+1].C * nodes[i].R
	}
	return int(math.Abs(float64(sum1-sum2) / 2))
}

func picksTheorem(area int, nodes []Coord) int {
	// area = interiorPoints + boundaryPoints/2 - 1
	return area - len(nodes)/2 + 1
}
