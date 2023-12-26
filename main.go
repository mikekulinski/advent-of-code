package main

import (
	"fmt"
	"github.com/mikekulinski/advent-of-code/day7"
	"github.com/mikekulinski/advent-of-code/file_reader"
	"log"
)

func main() {
	lines := file_reader.ReadFile("day7/input.txt")

	total, err := day7.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
