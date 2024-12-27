package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

const (
	mulPrefix = "mul("
	doToken   = "do()"
	dontToken = "don't()"
)

var re2 = regexp.MustCompile(`([0-9]{1,3})`)

func Part1(input []string) (int, error) {
	line := ""
	for _, l := range input {
		line += l
	}

	expressions := parse(line)
	fmt.Println(expressions)

	return evaluate(expressions)
}

func parse(line string) []string {
	expressions := []string{}
	temp := ""
	i := 0
	for i < len(line) {
		// Check if the next 4 characters match the prefix we're looking for.
		if i+len(mulPrefix)-1 < len(line) && line[i:i+4] == mulPrefix {
			temp = "mul("
			isValid := true
			// Check if the next 3 characters are digits.
			for n := 0; n < 4; n++ {
				char := line[i+len(temp)]
				if unicode.IsDigit(rune(char)) {
					temp += string(char)
				} else if char == ',' {
					temp += string(char)
					break
				} else {
					isValid = false
					break
				}
			}

			if !isValid {
				i++
				continue
			}

			// Check if the next 3 characters are digits.
			for n := 0; n < 4; n++ {
				char := line[i+len(temp)]
				if unicode.IsDigit(rune(char)) {
					temp += string(char)
				} else if line[i+len(temp)] == ')' {
					temp += string(char)
					break
				} else {
					isValid = false
					break
				}
			}

			if !isValid {
				i++
				continue
			}
			expressions = append(expressions, temp)
			i += len(temp)
		} else {
			i++
		}
	}

	return expressions
}

func evaluate(expressions []string) (int, error) {
	sumOfProducts := 0
	for _, exp := range expressions {
		nums := re2.FindAllString(exp, -1)
		first, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		second, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		sumOfProducts += first * second
	}

	return sumOfProducts, nil
}

// TODO: Figure out if things are mostly safe. What if we took out 1 bad level?
func Part2(input []string) (int, error) {
	line := ""
	for _, l := range input {
		line += l
	}

	expressions := parseWithToggle(line)
	fmt.Println(expressions)

	return evaluate(expressions)
}

func parseWithToggle(line string) []string {
	expressions := []string{}
	enabled := true
	temp := ""
	i := 0
	for i < len(line) {
		if i+len(doToken)-1 < len(line) && line[i:i+len(doToken)] == doToken {
			enabled = true
			i += len(doToken)
		} else if i+len(dontToken)-1 < len(line) && line[i:i+len(dontToken)] == dontToken {
			enabled = false
			i += len(dontToken)
		} else if i+3 < len(line) && line[i:i+4] == mulPrefix {
			// Check if the next 4 characters match the prefix we're looking for.
			temp = "mul("
			isValid := true
			// Check if the next 3 characters are digits.
			for n := 0; n < 4; n++ {
				char := line[i+len(temp)]
				if unicode.IsDigit(rune(char)) {
					temp += string(char)
				} else if char == ',' {
					temp += string(char)
					break
				} else {
					isValid = false
					break
				}
			}

			if !isValid {
				i++
				continue
			}

			// Check if the next 3 characters are digits.
			for n := 0; n < 4; n++ {
				char := line[i+len(temp)]
				if unicode.IsDigit(rune(char)) {
					temp += string(char)
				} else if line[i+len(temp)] == ')' {
					temp += string(char)
					break
				} else {
					isValid = false
					break
				}
			}

			if !isValid {
				i++
				continue
			}
			if !enabled {
				i++
				continue
			}
			expressions = append(expressions, temp)
			i += len(temp)
		} else {
			i++
		}
	}

	return expressions
}
