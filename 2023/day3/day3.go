package day3

import (
	"strconv"
)

/*
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source,
but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone!
The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one.
If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers
and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally,
is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right)
and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/

type PartLoc struct {
	// The row that this part is on.
	Row int
	// The starting x coordinate of this part, inclusive.
	StartX int
	// The ending x coordinate of this part, exclusive.
	EndX int
}

func Part1(input []string) (int, error) {
	grid := createGrid(input)
	parts := findParts(grid)
	validParts := searchGridForValidParts(grid, parts)
	return sumPartNumbers(validParts)
}

func createGrid(input []string) [][]string {
	grid := [][]string{}
	for _, line := range input {
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}
	return grid
}

// findParts searches the grid for collections of numbers that represent a part.
// For each part found, we will return the coordinates of the first digit in that part.
func findParts(grid [][]string) map[PartLoc]string {
	parts := map[PartLoc]string{}
	for i, row := range grid {
		window := ""
		for j, char := range row {
			_, err := strconv.Atoi(char)
			// If the character is a number, add it to the window.
			if err == nil {
				window += char
			} else if window != "" {
				// Otherwise, put the existing part in the map and reset the window.
				coords := PartLoc{
					Row:    i,
					StartX: j - len(window),
					EndX:   j,
				}
				parts[coords] = window
				window = ""
			}
		}

		// We might have a part at the end of the row. So if the window isn't empty, add it to the map.
		if window != "" {
			coords := PartLoc{
				Row:    i,
				StartX: len(row) - len(window),
				EndX:   len(row),
			}
			parts[coords] = window
		}
	}
	return parts
}

// searchGridForValidParts searches the grid for parts that are touching a symbol.
func searchGridForValidParts(grid [][]string, parts map[PartLoc]string) []string {
	validParts := []string{}
	for coords, part := range parts {
		for i := coords.StartX; i < coords.EndX; i++ {
			if isTouching(grid, coords.Row, i, isSymbol) {
				validParts = append(validParts, part)
				break
			}
		}
	}
	return validParts
}

// isTouching checks the surrounding 8 cells and checks if that cell satisfies the provided function.
// If it does, then return true. Otherwise, return false.
func isTouching(grid [][]string, row, col int, fn func(char string) bool) bool {
	// Check the row above.
	if row > 0 {
		if fn(grid[row-1][col]) {
			return true
		}
		if col > 0 && fn(grid[row-1][col-1]) {
			return true
		}
		if col < len(grid[row])-1 && fn(grid[row-1][col+1]) {
			return true
		}
	}
	// Check the row below.
	if row < len(grid)-1 {
		if fn(grid[row+1][col]) {
			return true
		}
		if col > 0 && fn(grid[row+1][col-1]) {
			return true
		}
		if col < len(grid[row])-1 && fn(grid[row+1][col+1]) {
			return true
		}
	}
	// Check the column to the left.
	if col > 0 && fn(grid[row][col-1]) {
		return true
	}
	// Check the column to the right.
	if col < len(grid[row])-1 && fn(grid[row][col+1]) {
		return true
	}
	return false
}

func isSymbol(char string) bool {
	symbols := []string{"#", "$", "%", "&", "*", "+", "-", "/", "=", "@"}
	for _, symbol := range symbols {
		if char == symbol {
			return true
		}
	}
	return false
}

func sumPartNumbers(parts []string) (int, error) {
	sum := 0
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return 0, err
		}
		sum += num
	}
	return sum, nil
}

/*
--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in
the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone
labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding
a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station.
You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is
adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out
which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35,
so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617
is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?
*/
func Part2(input []string) (int, error) {
	grid := createGrid(input)
	parts := findParts(grid)
	partsTouchingGears := findPartsTouchingGears(grid, parts)
	validGears := filterOutInvalidGears(partsTouchingGears)
	gearRatios, err := calculateGearRatios(validGears)
	if err != nil {
		return 0, err
	}
	return sumGearRatios(gearRatios), nil
}

type GearLoc struct {
	// The row that this gear is on.
	Row int
	// The column that this gear is on.
	Col int
}

func findPartsTouchingGears(grid [][]string, parts map[PartLoc]string) map[GearLoc][]string {
	gearsMap := map[GearLoc][]string{}
	for coords, part := range parts {
		// For each part, get a set of all of the gears that are touching it.
		gearsTouching := findGearsTouching(grid, coords.Row, coords.StartX)
		for i := coords.StartX; i < coords.EndX; i++ {
			touching := findGearsTouching(grid, coords.Row, i)
			gearsTouching = mergeMaps(gearsTouching, touching)
		}

		// For each of those gears, add that to a map of each gear and the parts that are touching it.
		for gear := range gearsTouching {
			gearsMap[gear] = append(gearsMap[gear], part)
		}
	}
	return gearsMap
}

func findGearsTouching(grid [][]string, row, col int) map[GearLoc]struct{} {
	loc := GearLoc{}
	gearsTouching := map[GearLoc]struct{}{}
	// Check the row above.
	if row > 0 {
		loc = GearLoc{Row: row - 1, Col: col}
		if isGear(grid, loc) {
			gearsTouching[loc] = struct{}{}
		}
		loc = GearLoc{Row: row - 1, Col: col - 1}
		if col > 0 && isGear(grid, loc) {
			gearsTouching[loc] = struct{}{}
		}
		loc = GearLoc{Row: row - 1, Col: col + 1}
		if col < len(grid[row])-1 && isGear(grid, loc) {
			gearsTouching[loc] = struct{}{}
		}
	}
	// Check the row below.
	if row < len(grid)-1 {
		loc = GearLoc{Row: row + 1, Col: col}
		if isGear(grid, loc) {
			gearsTouching[loc] = struct{}{}
		}
		loc = GearLoc{Row: row + 1, Col: col - 1}
		if col > 0 && isGear(grid, loc) {
			gearsTouching[loc] = struct{}{}
		}
		loc = GearLoc{Row: row + 1, Col: col + 1}
		if col < len(grid[row])-1 && isGear(grid, loc) {
			gearsTouching[loc] = struct{}{}
		}
	}
	// Check the column to the left.
	loc = GearLoc{Row: row, Col: col - 1}
	if col > 0 && isGear(grid, loc) {
		gearsTouching[loc] = struct{}{}
	}
	// Check the column to the right.
	loc = GearLoc{Row: row, Col: col + 1}
	if col < len(grid[row])-1 && isGear(grid, loc) {
		gearsTouching[loc] = struct{}{}
	}
	return gearsTouching
}

func isGear(grid [][]string, gearLoc GearLoc) bool {
	char := grid[gearLoc.Row][gearLoc.Col]
	return char == "*"
}

// mergeMaps adds all of the values from m2 into m1 and returns m1.
func mergeMaps(m1, m2 map[GearLoc]struct{}) map[GearLoc]struct{} {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

// filterOutInvalidGears removes any gears that aren't touching exactly two parts.
func filterOutInvalidGears(gears map[GearLoc][]string) map[GearLoc][]string {
	validGears := map[GearLoc][]string{}
	for gear, parts := range gears {
		if len(parts) == 2 {
			validGears[gear] = parts
		}
	}
	return validGears
}

// calculateGearRatios takes the two parts that are a touching each gear,
// and calculates the gear ratio by multiplying the two parts together.
func calculateGearRatios(gears map[GearLoc][]string) (map[GearLoc]int, error) {
	gearRatios := map[GearLoc]int{}
	for gear, parts := range gears {
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		gearRatios[gear] = num1 * num2
	}
	return gearRatios, nil
}

func sumGearRatios(gearRatios map[GearLoc]int) int {
	sum := 0
	for _, ratio := range gearRatios {
		sum += ratio
	}
	return sum
}
