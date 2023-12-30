package main

import (
	"fmt"
	"github.com/mikekulinski/advent-of-code/day11"
	"github.com/mikekulinski/advent-of-code/file_reader"
	"log"
)

func main() {
	lines := file_reader.ReadFile("day11/input.txt")

	total, err := day11.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
