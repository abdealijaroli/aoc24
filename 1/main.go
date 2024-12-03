package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calcTotalDistance(list1, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		distance := int(math.Abs(float64(list1[i] - list2[i])))
		totalDistance += distance
	}

	return totalDistance
}

func main() {
	var list1, list2 []int

	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			continue
		}

		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	fmt.Println(calcTotalDistance(list1, list2))
}
