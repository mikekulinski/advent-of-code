package main

import (
	"fmt"
	"github.com/mikekulinski/advent-of-code/day9"
	"github.com/mikekulinski/advent-of-code/file_reader"
	"log"
)

func main() {
	lines := file_reader.ReadFile("day9/input.txt")

	total, err := day9.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
