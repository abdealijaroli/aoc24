package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TransformResult struct {
	count   int
	pattern []int
}

var memo = make(map[string]TransformResult)

func predictGrowth(n int, depth int) int {
	key := fmt.Sprintf("%d_%d", n, depth)
	if result, exists := memo[key]; exists {
		return result.count
	}

	if depth == 0 {
		return 1
	}

	var nextCount int
	if n == 0 {
		nextCount = predictGrowth(1, depth-1)
	} else if len(strconv.Itoa(n))%2 == 0 {
		s := strconv.Itoa(n)
		mid := len(s) / 2
		left, _ := strconv.Atoi(s[:mid])
		right, _ := strconv.Atoi(s[mid:])
		nextCount = predictGrowth(left, depth-1) + predictGrowth(right, depth-1)
	} else {
		nextCount = predictGrowth(n*2024, depth-1)
	}

	memo[key] = TransformResult{nextCount, nil}
	return nextCount
}

func solve(nums []int, depth int) int {
	total := 0
	for _, n := range nums {
		total += predictGrowth(n, depth)
	}
	return total
}

func main() {
	file, err := os.Open("11/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int

	if scanner.Scan() {
		for _, numStr := range strings.Fields(scanner.Text()) {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
	}

	result := solve(nums, 75)
	fmt.Println(result)
}
