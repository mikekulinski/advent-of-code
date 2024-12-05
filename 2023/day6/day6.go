package day6

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
--- Day 6: Wait For It ---
The ferry quickly brings you across Island Island. After asking around, you discover that there is indeed normally a
large pile of sand somewhere near here, but you don't see anything besides lots of water and the small island where the
ferry has docked.

As you try to figure out what to do next, you notice a poster on a wall near the ferry dock. "Boat races! Open to the
public! Grand prize is an all-expenses-paid trip to Desert Island!" That must be where the sand comes from! Best of all,
the boat races are starting in just a few minutes.

You manage to sign up as a competitor in the boat races just in time. The organizer explains that it's not really a
traditional race - instead, you will get a fixed amount of time during which your boat has to travel as far as it can,
and you win if your boat goes the farthest.

As part of signing up, you get a sheet of paper (your puzzle input) that lists the time allowed for each race and also
the best distance ever recorded in that race. To guarantee you win the grand prize, you need to make sure you go farther
in each race than the current record holder.

The organizer brings you over to the area where the boat races are held. The boats are much smaller than you expected
- they're actually toy boats, each with a big button on top. Holding down the button charges the boat, and releasing
the button allows the boat to move. Boats move faster if their button was held longer, but time spent holding the
button counts against the total race time. You can only hold the button at the start of the race, and boats don't move
until the button is released.

For example:

Time:      7  15   30
Distance:  9  40  200
This document describes three races:

The first race lasts 7 milliseconds. The record distance in this race is 9 millimeters.
The second race lasts 15 milliseconds. The record distance in this race is 40 millimeters.
The third race lasts 30 milliseconds. The record distance in this race is 200 millimeters.
Your toy boat has a starting speed of zero millimeters per millisecond. For each whole millisecond you spend
at the beginning of the race holding down the button, the boat's speed increases by one millimeter per millisecond.

So, because the first race lasts 7 milliseconds, you only have a few options:

Don't hold the button at all (that is, hold it for 0 milliseconds) at the start of the race. The boat won't move;
it will have traveled 0 millimeters by the end of the race.
Hold the button for 1 millisecond at the start of the race. Then, the boat will travel at a speed of 1 millimeter
per millisecond for 6 milliseconds, reaching a total distance traveled of 6 millimeters.
Hold the button for 2 milliseconds, giving the boat a speed of 2 millimeters per millisecond. It will then get 5 milliseconds to move,
reaching a total distance of 10 millimeters.
Hold the button for 3 milliseconds. After its remaining 4 milliseconds of travel time, the boat will have gone 12 millimeters.
Hold the button for 4 milliseconds. After its remaining 3 milliseconds of travel time, the boat will have gone 12 millimeters.
Hold the button for 5 milliseconds, causing the boat to travel a total of 10 millimeters.
Hold the button for 6 milliseconds, causing the boat to travel a total of 6 millimeters.
Hold the button for 7 milliseconds. That's the entire duration of the race. You never let go of the button.
The boat can't move until you let go of the button. Please make sure you let go of the button so the boat gets to move.
0 millimeters.
Since the current record for this race is 9 millimeters, there are actually 4 different ways you could win:
you could hold the button for 2, 3, 4, or 5 milliseconds at the start of the race.

In the second race, you could hold the button for at least 4 milliseconds and at most 11 milliseconds and beat the record,
a total of 8 different ways to win.

In the third race, you could hold the button for at least 11 milliseconds and no more than 19 milliseconds and still
beat the record, a total of 9 ways you could win.

To see how much margin of error you have, determine the number of ways you can beat the record in each race;
in this example, if you multiply these values together, you get 288 (4 * 8 * 9).

Determine the number of ways you could beat the record in each race. What do you get if you multiply these
numbers together?
*/

func Part1(lines []string) (int, error) {
	times := parseLine(lines[0])
	distanceRecords := parseLine(lines[1])
	fmt.Println("Times: ", times)
	fmt.Println("Distance Records: ", distanceRecords)
	permutations := totalPermutations(times, distanceRecords)
	return permutations, nil
}

func parseLine(line string) []int {
	split := strings.Split(line, ":")
	numsStr := strings.Fields(split[1])
	var nums []int
	for _, numStr := range numsStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func totalPermutations(maxTimes []int, distanceRecords []int) int {
	permutations := 1
	for i := range maxTimes {
		maxTime := float64(maxTimes[i])
		distanceRecord := float64(distanceRecords[i])
		numWays := numWaysToWin(maxTime, distanceRecord)
		permutations *= numWays
	}
	return permutations
}

func numWaysToWin(maxTime float64, distanceRecord float64) int {
	lower, upper := winningBoundaries(maxTime, distanceRecord)
	fmt.Println("Lower: ", lower, " Upper: ", upper)
	numWays := int(math.Abs(lower-upper) + 1)
	return numWays
}

func winningBoundaries(maxTime float64, distanceRecord float64) (float64, float64) {
	zero1, zero2 := CalculateZeros(maxTime, distanceRecord)
	minZero := min(zero1, zero2)
	maxZero := max(zero1, zero2)
	return nextInteger(minZero), integerBefore(maxZero)
}

// We can figure out at what point we can break the distance record for the allowed time
// by modeling the problem as a quadratic equation and finding the zeros.
// `0 = ax^2 + bx + c` translates to `0 = t_hold^2 - t_total * t_hold + x_record`
// We can solve this by using the quadratic formula.
func CalculateZeros(maxTime float64, distanceRecord float64) (float64, float64) {
	root := math.Sqrt(math.Pow(maxTime, 2) - 4*distanceRecord)
	first := (maxTime + root) / 2
	second := (maxTime - root) / 2
	return first, second
}

func nextInteger(num float64) float64 {
	if math.Ceil(num) == num {
		return num + 1
	}
	return math.Ceil(num)
}

func integerBefore(num float64) float64 {
	if math.Floor(num) == num {
		return num - 1
	}
	return math.Floor(num)
}

/*
--- Part Two ---
As the race is about to start, you realize the piece of paper with race times and record distances you got earlier
actually just has very bad kerning. There's really only one race - ignore the spaces between the numbers on each line.

So, the example from before:

Time:      7  15   30
Distance:  9  40  200
...now instead means this:

Time:      71530
Distance:  940200
Now, you have to figure out how many ways there are to win this single race. In this example, the race lasts for
71530 milliseconds and the record distance you need to beat is 940200 millimeters. You could hold the button
anywhere from 14 to 71516 milliseconds and beat the record, a total of 71503 ways!

How many ways can you beat the record in this one much longer race?
*/
func Part2(lines []string) (int, error) {
	time := parseLinePart2(lines[0])
	distanceRecord := parseLinePart2(lines[1])
	fmt.Println("Times: ", time)
	fmt.Println("Distance Records: ", distanceRecord)
	numWays := numWaysToWin(float64(time), float64(distanceRecord))
	return numWays, nil
}

func parseLinePart2(line string) int {
	split := strings.Split(line, ":")
	numsStr := strings.Fields(split[1])
	combinedNumStr := ""
	for _, numStr := range numsStr {
		combinedNumStr += numStr
	}
	num, err := strconv.Atoi(combinedNumStr)
	if err != nil {
		panic(err)
	}
	return num
}
