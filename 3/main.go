package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func sanitizeInput(inputStrings []string) int {
	var validStrArr []string
	mulRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	for _, str := range inputStrings {
		validStr := mulRegex.FindAllString(str, 1e10)
		validStrArr = append(validStrArr, validStr...)
	}
	// fmt.Println(validStrArr)
	return extractNumMatrix(validStrArr)
}

func extractNumMatrix(validStrings []string) int {
	result := make([][]int, len(validStrings))
	numRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for i, str := range validStrings {
		matches := numRegex.FindStringSubmatch(str)
		if len(matches) == 3 {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			result[i] = []int{x, y}
		}
	}
	// fmt.Println(result)
	return calcProduct(result)
}

func calcProduct(numMatrix [][]int) int {
	product := 0
	for _, pair := range numMatrix {
		product += pair[0] * pair[1]
	}
	// fmt.Println(product)
	return product
}

func main() {
	var sb strings.Builder

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sb.WriteString(scanner.Text())
		sb.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	input := sb.String()

	inputStrArr := strings.Split(strings.TrimSpace(input), "\n")

	if len(inputStrArr) == 0 {
		fmt.Fprintln(os.Stderr, "no input provided")
		os.Exit(1)
	}

	fmt.Println(sanitizeInput(inputStrArr))
}
