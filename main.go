package main

import (
	"fmt"
	"github.com/mikekulinski/advent-of-code/day5"
	"github.com/mikekulinski/advent-of-code/file_reader"
	"log"
)

func main() {
	lines := file_reader.ReadFile("day5/input.txt")

	total, err := day5.Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
