package main

import (
	"fmt"
	"github.com/mikekulinski/advent-of-code/day1"
	"github.com/mikekulinski/advent-of-code/file_reader"
)

func main() {
	lines := file_reader.ReadFile("day1/input.txt")

	total, err := day1.Part2(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println(total)
}
