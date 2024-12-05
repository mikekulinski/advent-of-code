package day11

import (
	"math"
)

/*
--- Day 11: Cosmic Expansion ---
You continue following signs for "Hot Springs" and eventually come across an observatory. The Elf within turns out
to be a researcher studying cosmic expansion using the giant telescope here.

He doesn't know anything about the missing machine parts; he's only visiting for this research project. However,
he confirms that the hot springs are the next-closest area likely to have people; he'll even take you straight there
once he's done with today's observation analysis.

Maybe you can help him with the analysis to speed things up?

The researcher has collected a bunch of data and compiled the data into a single giant image (your puzzle input).
The image includes empty space (.) and galaxies (#). For example:

...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
The researcher is trying to figure out the sum of the lengths of the shortest path between every pair of galaxies.
However, there's a catch: the universe expanded in the time it took the light from those galaxies to reach
the observatory.

Due to something involving gravitational effects, only some space expands. In fact, the result is that any rows
or columns that contain no galaxies should all actually be twice as big.

In the above example, three columns and two rows contain no galaxies:

   v  v  v
 ...#......
 .......#..
 #.........
>..........<
 ......#...
 .#........
 .........#
>..........<
 .......#..
 #...#.....
   ^  ^  ^
These rows and columns need to be twice as big; the result of cosmic expansion therefore looks like this:

....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......
Equipped with this expanded universe, the shortest path between every pair of galaxies can be found. It can help
to assign every galaxy a unique number:

....1........
.........2...
3............
.............
.............
........4....
.5...........
............6
.............
.............
.........7...
8....9.......
In these 9 galaxies, there are 36 pairs. Only count each pair once; order within the pair doesn't matter. For each pair,
find any shortest path between the two galaxies using only steps that move up, down, left, or right exactly one . or #
at a time. (The shortest path between two galaxies is allowed to pass through another galaxy.)

For example, here is one of the shortest paths between galaxies 5 and 9:

....1........
.........2...
3............
.............
.............
........4....
.5...........
.##.........6
..##.........
...##........
....##...7...
8....9.......
This path has length 9 because it takes a minimum of nine steps to get from galaxy 5 to galaxy 9 (the eight locations
marked # plus the step onto galaxy 9 itself). Here are some other example shortest path lengths:

Between galaxy 1 and galaxy 7: 15
Between galaxy 3 and galaxy 6: 17
Between galaxy 8 and galaxy 9: 5
In this example, after expanding the universe, the sum of the shortest path between all 36 pairs of galaxies is 374.

Expand the universe, then find the length of the shortest path between every pair of galaxies. What is the sum of these
lengths?
*/

type Coord struct {
	R int
	C int
}

func Part1(lines []string) (int, error) {
	grid := parseInput(lines)
	galaxies := findGalaxyLocations(grid)
	emptyRows := findEmptyRows(grid)
	emptyCols := findEmptyColumns(grid)
	distance := distanceBetweenAllGalaxies(galaxies, emptyRows, emptyCols)
	return distance, nil
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

func findGalaxyLocations(grid [][]string) []Coord {
	var galaxies []Coord
	for r := range grid {
		for c := range grid {
			if grid[r][c] == "#" {
				galaxies = append(galaxies, Coord{r, c})
			}
		}
	}
	return galaxies
}

func findEmptyRows(grid [][]string) []int {
	var emptyRows []int
	for r := range grid {
		isEmpty := true
		for c := range grid[r] {
			if grid[r][c] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, r)
		}
	}
	return emptyRows
}

func findEmptyColumns(grid [][]string) []int {
	var emptyColumns []int
	for c := 0; c < len(grid[0]); c++ {
		isEmpty := true
		for r := range grid[c] {
			if grid[r][c] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyColumns = append(emptyColumns, c)
		}
	}
	return emptyColumns
}

func distanceBetweenAllGalaxies(galaxies []Coord, emptyRows []int, emptyCols []int) int {
	total := 0
	for i := 0; i < len(galaxies); i++ {
		g := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			h := galaxies[j]
			total += manhattanDistWithSpaceWarp(g, h, emptyRows, emptyCols)
		}
	}
	return total
}

// manhattanDistWithSpaceWarp will calculate the manhattan distance between the two points, while
// also taking into account that empty rows and columns grow by 2x due to cosmic expansion.
func manhattanDistWithSpaceWarp(g, h Coord, emptyRows []int, emptyCols []int) int {
	distR := int(math.Abs(float64(g.R - h.R)))
	distC := int(math.Abs(float64(g.C - h.C)))
	return distR + distC + howMuchEmptySpace(g.R, h.R, emptyRows) + howMuchEmptySpace(g.C, h.C, emptyCols)
}

func howMuchEmptySpace(p1, p2 int, empties []int) int {
	total := 0
	for _, empty := range empties {
		if (p1 < empty && p2 > empty) ||
			(p1 > empty && p2 < empty) {
			total++
		}
	}
	return total
}

/*
--- Part Two ---
The galaxies are much older (and thus much farther apart) than the researcher initially estimated.

Now, instead of the expansion you did before, make each empty row or column one million times larger.
That is, each empty row should be replaced with 1000000 empty rows, and each empty column should be replaced
with 1000000 empty columns.

(In the example above, if each empty row or column were merely 10 times larger, the sum of the shortest paths
between every pair of galaxies would be 1030. If each empty row or column were merely 100 times larger, the sum
of the shortest paths between every pair of galaxies would be 8410. However, your universe will need to expand
far beyond these values.)

Starting with the same initial image, expand the universe according to these new rules, then find the length of
the shortest path between every pair of galaxies. What is the sum of these lengths?
*/

func Part2(lines []string) (int, error) {
	grid := parseInput(lines)
	galaxies := findGalaxyLocations(grid)
	emptyRows := findEmptyRows(grid)
	emptyCols := findEmptyColumns(grid)
	distance := distanceBetweenAllGalaxiesWithMegaSpaceWarp(galaxies, emptyRows, emptyCols)
	return distance, nil
}

// manhattanDistWithMegaSpaceWarp will calculate the manhattan distance between the two points, while
// also taking into account that empty rows and columns grow by 2x due to cosmic expansion.
func manhattanDistWithMegaSpaceWarp(g, h Coord, emptyRows []int, emptyCols []int) int {
	distR := int(math.Abs(float64(g.R - h.R)))
	distC := int(math.Abs(float64(g.C - h.C)))
	return distR + distC + 999999*howMuchEmptySpace(g.R, h.R, emptyRows) + 999999*howMuchEmptySpace(g.C, h.C, emptyCols)
}

func distanceBetweenAllGalaxiesWithMegaSpaceWarp(galaxies []Coord, emptyRows []int, emptyCols []int) int {
	total := 0
	for i := 0; i < len(galaxies); i++ {
		g := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			h := galaxies[j]
			total += manhattanDistWithMegaSpaceWarp(g, h, emptyRows, emptyCols)
		}
	}
	return total
}
