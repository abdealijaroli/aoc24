package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func evaluate(nums []int, ops []string) int {
    result := nums[0]
    for i := 0; i < len(ops); i++ {
        switch ops[i] {
        case "+":
            result += nums[i+1]
        case "*":
            result *= nums[i+1]
        case "||":
            resultStr := fmt.Sprintf("%d%d", result, nums[i+1])
            result, _ = strconv.Atoi(resultStr)
        }
    }
    return result
}

func canMakeTarget(target int, nums []int) bool {
    if len(nums) == 1 {
        return nums[0] == target
    }
    
    numOps := len(nums) - 1
    maxCombinations := 1
    for i := 0; i < numOps; i++ {
        maxCombinations *= 3
    }
    
    for i := 0; i < maxCombinations; i++ {
        ops := make([]string, numOps)
        temp := i
        for j := 0; j < numOps; j++ {
            switch temp % 3 {
            case 0:
                ops[j] = "+"
            case 1:
                ops[j] = "*"
            case 2:
                ops[j] = "||"
            }
            temp /= 3
        }
        
        if evaluate(nums, ops) == target {
            return true
        }
    }
    
    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    sum := 0
    
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ":")
        target, _ := strconv.Atoi(parts[0])
        
        numStr := strings.Fields(parts[1])
        nums := make([]int, len(numStr))
        for i, str := range numStr {
            nums[i], _ = strconv.Atoi(str)
        }
        
        if canMakeTarget(target, nums) {
            sum += target
        }
    }
    
    fmt.Println(sum)
}