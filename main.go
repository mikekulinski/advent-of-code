package main

import (
	"fmt"
	"github.com/mikekulinski/advent-of-code/day4"
	"github.com/mikekulinski/advent-of-code/file_reader"
	"log"
)

func main() {
	lines := file_reader.ReadFile("day4/input.txt")

	total, err := day4.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
