package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sync"
)

type cacheKey struct {
    num   int
    depth int
}

var cache = make(map[cacheKey][]int)
var cacheMutex sync.RWMutex

func transformZero(n int) []int {
    if n == 0 {
        return []int{1}
    }
    return []int{n}
}

func splitEvenDigits(n int) []int {
    s := strconv.Itoa(n)
    if len(s)%2 == 0 {
        mid := len(s)/2
        left, _ := strconv.Atoi(s[:mid])
        right, _ := strconv.Atoi(s[mid:])
        return []int{left, right}
    }
    return []int{n}
}

func multiplyBy2024(n int) []int {
    return []int{n * 2024}
}

func transformWithCache(stone int, depth int) []int {
    key := cacheKey{stone, depth}
    
    cacheMutex.RLock()
    if result, ok := cache[key]; ok {
        cacheMutex.RUnlock()
        return result
    }
    cacheMutex.RUnlock()
    
    var result []int
    if depth == 0 {
        return []int{stone}
    }
    
    if stone == 0 {
        result = transformZero(stone)
    } else if len(strconv.Itoa(stone))%2 == 0 {
        result = splitEvenDigits(stone)
    } else {
        result = multiplyBy2024(stone)
    }
    
    var nextResult []int
    for _, n := range result {
        nextResult = append(nextResult, transformWithCache(n, depth-1)...)
    }
    
    cacheMutex.Lock()
    cache[key] = nextResult
    cacheMutex.Unlock()
    
    return nextResult
}

func transformParallel(stones []int, depth int) []int {
    var wg sync.WaitGroup
    resultChan := make(chan []int, len(stones))
    
    for _, stone := range stones {
        wg.Add(1)
        go func(s int) {
            defer wg.Done()
            resultChan <- transformWithCache(s, depth)
        }(stone)
    }
    
    go func() {
        wg.Wait()
        close(resultChan)
    }()
    
    var result []int
    for r := range resultChan {
        result = append(result, r...)
    }
    return result
}

func main() {
    file, err := os.Open("11/input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var nums []int

    if scanner.Scan() {
        line := scanner.Text()
        for _, numStr := range strings.Fields(line) {
            num, _ := strconv.Atoi(numStr)
            nums = append(nums, num)
        }
    }

    stones := transformParallel(nums, 25)
    fmt.Println(len(stones))
}