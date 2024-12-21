package day2

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input []string) (int, error) {
	reports := [][]int{}
	for _, line := range input {
		nums := strings.Fields(line)
		report := []int{}
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			report = append(report, i)
		}
		reports = append(reports, report)
	}

	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}
	return safeReports, nil
}

func isSafe(report []int) bool {
	previousLevel := report[0]
	for _, level := range report[1:] {
		diff := abs(previousLevel - level)
		if diff < 1 || diff > 3 {
			return false
		}
		previousLevel = level
	}

	if isIncreasing(report) || isDecreasing(report) {
		return true
	}
	return false
}

func isIncreasing(report []int) bool {
	previousLevel := report[0]
	for _, level := range report[1:] {
		if level <= previousLevel {
			return false
		}
		previousLevel = level
	}
	return true
}

func isDecreasing(report []int) bool {
	previousLevel := report[0]
	for _, level := range report[1:] {
		if level >= previousLevel {
			return false
		}
		previousLevel = level
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// TODO: Figure out if things are mostly safe. What if we took out 1 bad level?
func Part2(input []string) (int, error) {
	reports := [][]int{}
	for _, line := range input {
		nums := strings.Fields(line)
		report := []int{}
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			report = append(report, i)
		}
		reports = append(reports, report)
	}

	safeReports := 0
	for _, report := range reports {
		if canBeSafe(report) {
			safeReports++
			continue
		}
	}

	return safeReports, nil
}

func canBeSafe(report []int) bool {
	if isMostlySafe(report) {
		return true
	}

	// Problem dampener.
	for i := range report {
		// Try removing the item at this index and see if that fixes the problem.
		sliced := slices.Concat(report[:i], report[i+1:])
		if isMostlySafe(sliced) {
			return true
		}
	}
	return false
}

func isMostlySafe(report []int) bool {
	previousLevel := report[0]
	for _, level := range report[1:] {
		diff := abs(previousLevel - level)
		if diff < 1 || diff > 3 {
			return false
		}
		previousLevel = level
	}

	if isIncreasing(report) || isDecreasing(report) {
		return true
	}

	return false
}
