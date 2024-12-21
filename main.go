package main

import (
	"fmt"
	"github.com/mikekulinski/advent-of-code/2024/day2"
	"github.com/mikekulinski/advent-of-code/file_reader"
	"log"
)

func main() {
	lines := file_reader.ReadFile("2024/day2/input.txt")

	total, err := day2.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
