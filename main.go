package main

import (
	"fmt"
	"github.com/mikekulinski/advent-of-code/day10"
	"github.com/mikekulinski/advent-of-code/file_reader"
	"log"
)

func main() {
	lines := file_reader.ReadFile("day10/input.txt")

	total, err := day10.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
