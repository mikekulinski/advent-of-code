package day1

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input []string) (int, error) {
	list1 := []int{}
	list2 := []int{}
	for _, line := range input {
		nums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		list1 = append(list1, num1)
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		list2 = append(list2, num2)
	}

	slices.Sort(list1)
	slices.Sort(list2)
	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		totalDistance += abs(list1[i] - list2[i])
	}
	return totalDistance, nil
}

func Part2(input []string) (int, error) {
	list1 := []int{}
	list2 := []int{}
	for _, line := range input {
		nums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		list1 = append(list1, num1)
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		list2 = append(list2, num2)
	}

	// Put the second list into a map defined as (num, # times it appears)
	appearances := map[int]int{}
	for _, num := range list2 {
		appearances[num]++
	}

	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		num := list1[i]
		totalDistance += num * appearances[num]
	}
	return totalDistance, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
