package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidOrder(nums []int, rules map[int][]int) bool {
	pos := make(map[int]int)
	for i, num := range nums {
		pos[num] = i
	}

	for i := range nums {
		num := nums[i]
		if deps, exists := rules[num]; exists {
			for _, dep := range deps {
				if depPos, found := pos[dep]; found && depPos <= i {
					return false
				}
			}
		}
	}
	return true
}

func correctOrder(nums []int, rules map[int][]int) []int {
	corrected := make([]int, len(nums))
	copy(corrected, nums)

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(corrected)-1; i++ {
			num1, num2 := corrected[i], corrected[i+1]
			if deps, exists := rules[num2]; exists {
				for _, dep := range deps {
					if dep == num1 {
						corrected[i], corrected[i+1] = corrected[i+1], corrected[i]
						changed = true
						break
					}
				}
			}
		}
	}
	return corrected
}

func parseInput(input string) (map[int][]int, [][]int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rules := make(map[int][]int)
	var updates [][]int

	inRules := true
	for _, line := range lines {
		if line == "" {
			inRules = false
			continue
		}

		if inRules {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
				y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err1 == nil && err2 == nil {
					rules[x] = append(rules[x], y)
				}
			}
		} else {
			var nums []int
			for _, numStr := range strings.Split(line, ",") {
				n, err := strconv.Atoi(strings.TrimSpace(numStr))
				if err == nil {
					nums = append(nums, n)
				}
			}
			if len(nums) > 0 {
				updates = append(updates, nums)
			}
		}
	}

	return rules, updates, nil
}

func main() {
	data, err := os.ReadFile("5/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	content := string(data)
	if !strings.Contains(content, "\n\n") {
		lines := strings.Split(content, "\n")
		for i := len(lines) - 1; i >= 0; i-- {
			if strings.Contains(lines[i], "|") {
				content = strings.Join(lines[:i+1], "\n") + "\n\n" + strings.Join(lines[i+1:], "\n")
				break
			}
		}
	}

	rules, updates, err := parseInput(content)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	sum := 0
	for _, update := range updates {
		if !isValidOrder(update, rules) {
			corrected := correctOrder(update, rules)
			midIndex := len(corrected) / 2
			sum += corrected[midIndex]
		}
	}

	fmt.Println("Sum:", sum)
}
