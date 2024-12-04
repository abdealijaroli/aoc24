package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(report []int) bool {
    if len(report) < 2 {
        return true
    }
    
    isIncreasing := report[1] > report[0]
    
    for i := 1; i < len(report); i++ {
        diff := report[i] - report[i-1]
        if diff < -3 || diff > 3 || diff == 0 {
            return false
        }
        if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
            return false
        }
    }
    return true
}

func isSafeWithDampener(report []int) bool {
    if isReportSafe(report) {
        return true
    }
    
    for i := range report {
        dampened := make([]int, 0, len(report)-1)
        dampened = append(dampened, report[:i]...)
        dampened = append(dampened, report[i+1:]...)
        
        if isReportSafe(dampened) {
            return true
        }
    }
    return false
}

func calcSafeReports(reports [][]int) int {
    safeReports := 0
    for _, report := range reports {
        if isSafeWithDampener(report) {
            safeReports++
        }
    }
    return safeReports
}

// Part 1 ====================================================
// func calcSafeReports(reports [][]int) int {
// 	safeReports := 0

// 	for _, report := range reports {
// 		isSafe := true
// 		isIncreasing := report[1] > report[0]

// 		for i := 1; i < len(report); i++ {
// 			diff := report[i] - report[i-1]

// 			if diff < -3 || diff > 3 || diff == 0 {
// 				isSafe = false
// 				break
// 			}

// 			if isIncreasing && diff < 0 {
// 				isSafe = false
// 				break
// 			}
// 			if !isIncreasing && diff > 0 {
// 				isSafe = false
// 				break
// 			}
// 		}
// 		if isSafe {
// 			safeReports++
// 		}
// 	}

// 	return safeReports
// }

func main() {
	var reports [][]int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		var levels []int

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				continue
			}
			levels = append(levels, num)
		}

		if len(levels) > 0 {
			reports = append(reports, levels)
		}
	}

	fmt.Println(calcSafeReports(reports))
}
