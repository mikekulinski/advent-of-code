package day8

import (
	"fmt"
	"regexp"
	"strings"
)

/*
--- Day 8: Haunted Wasteland ---
You're still riding a camel across Desert Island when you spot a sandstorm quickly approaching. When you turn to warn
the Elf, she disappears before your eyes! To be fair, she had just finished warning you about ghosts a few minutes ago.

One of the camel's pouches is labeled "maps" - sure enough, it's full of documents (your puzzle input) about how to
navigate the desert. At least, you're pretty sure that's what they are; one of the documents contains a list of
left/right instructions, and the rest of the documents seem to describe some kind of network of labeled nodes.

It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel
follow the same instructions, you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you
have to follow the left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

# RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
Starting with AAA, you need to look up the next element based on the next left/right instruction in your input.
In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose
the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.

Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the whole sequence of
instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example, here is a situation that takes
6 steps to reach ZZZ:

# LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?
*/
var (
	// expr is a regular expression that matches the 3 groups of letters that we'd like to capture.
	expr = regexp.MustCompile(`([A-Za-z0-9]+)\s*=\s*\(([A-Za-z0-9]+),\s*([A-Za-z0-9]+)\)`)
)

func Part1(lines []string) (int, error) {
	directions, graph := parseInput(lines)
	numSteps := traverseGraph("AAA", shouldStopZZZ, directions, graph)
	return numSteps, nil
}

func parseInput(lines []string) ([]string, map[string][]string) {
	directions := strings.Split(lines[0], "")

	graph := map[string][]string{}
	for _, line := range lines[2:] {
		letterGroups := expr.FindStringSubmatch(line)
		// letterGroups[0] is the whole match, and the rest are the individual capture groups
		// which represent each of the groups of letters.
		graph[letterGroups[1]] = []string{letterGroups[2], letterGroups[3]}
	}
	return directions, graph
}

func traverseGraph(startingNode string, shouldStop func(string) bool, directions []string, graph map[string][]string) int {
	currentNode := startingNode
	i := 0
	for {
		// Continuously pick the next direction, looping over the directions if we run out.
		direction := directions[i%len(directions)]
		if direction == "L" {
			currentNode = graph[currentNode][0]
		} else if direction == "R" {
			currentNode = graph[currentNode][1]
		} else {
			panic(fmt.Sprintf("invalid direction: %s", direction))
		}

		i++
		if shouldStop(currentNode) {
			break
		}
	}
	return i
}

func shouldStopZZZ(currentNode string) bool {
	return currentNode == "ZZZ"
}

/*
--- Part Two ---
The sandstorm is upon you and you aren't any closer to escaping the wasteland. You had the camel follow the
instructions, but you've barely left your starting position. It's going to take significantly more steps to escape!

What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws of spacetime?
Only one way to find out.

After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names
ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that
ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.

For example:

LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right
instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat
this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z,
they act like any other node and you continue as normal.) In this example, you would proceed as follows:

Step 0: You are at 11A and 22A.
Step 1: You choose all of the left paths, leading you to 11B and 22B.
Step 2: You choose all of the right paths, leading you to 11Z and 22C.
Step 3: You choose all of the left paths, leading you to 11B and 22Z.
Step 4: You choose all of the right paths, leading you to 11Z and 22B.
Step 5: You choose all of the left paths, leading you to 11B and 22C.
Step 6: You choose all of the right paths, leading you to 11Z and 22Z.
So, in this example, you end up entirely on nodes that end in Z after 6 steps.

Simultaneously start on every node that ends with A. How many steps does it take before you're only on nodes that
end with Z?
*/

func Part2(lines []string) (int, error) {
	directions, graph := parseInput(lines)
	numSteps := traverseGraphLikeAGhost(directions, graph)
	return numSteps, nil
}

func traverseGraphLikeAGhost(directions []string, graph map[string][]string) int {
	startingNodes := findStartingNodes(graph)

	var steps []int
	for _, startingNode := range startingNodes {
		numSteps := traverseGraph(startingNode, shouldStopZ, directions, graph)
		steps = append(steps, numSteps)
	}
	return leastCommonMultiple(steps)
}

func findStartingNodes(graph map[string][]string) []string {
	startingNodes := []string{}
	for node := range graph {
		if strings.HasSuffix(node, "A") {
			startingNodes = append(startingNodes, node)
		}
	}
	return startingNodes
}

func shouldStopZ(currentNode string) bool {
	return strings.HasSuffix(currentNode, "Z")
}

func leastCommonMultiple(steps []int) int {
	lcm := steps[0]
	for _, step := range steps[1:] {
		lcm = lcm * step / greatestCommonDivisor(lcm, step)
	}
	return lcm
}

// Find the greatest common divisor of two numbers using Euclid's algorithm.
func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
