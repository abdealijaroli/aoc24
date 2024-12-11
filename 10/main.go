package main

import (
    "bufio"
    "fmt"
    "os"
)

type Point struct {
    x, y int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readLines(path string) []string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func parseGrid(lines []string) [][]int {
    grid := make([][]int, len(lines))
    for i := range grid {
        grid[i] = make([]int, len(lines[0]))
        for j, c := range lines[i] {
            grid[i][j] = int(c - '0')
        }
    }
    return grid
}

func findTrailheads(grid [][]int) []Point {
    var trailheads []Point
    for y := 0; y < len(grid); y++ {
        for x := 0; x < len(grid[0]); x++ {
            if grid[y][x] == 0 {
                trailheads = append(trailheads, Point{x, y})
            }
        }
    }
    return trailheads
}

func isValid(x, y int, grid [][]int) bool {
    return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}

func countPaths(grid [][]int, start Point) int {
    paths := 0
    visited := make(map[Point]bool)
    
    var dfs func(p Point, height int)
    dfs = func(p Point, height int) {
        if !isValid(p.x, p.y, grid) || grid[p.y][p.x] != height {
            return
        }
        
        if visited[p] {
            return
        }
        
        if height == 9 {
            paths++
            return
        }
        
        visited[p] = true
        dirs := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
        for _, d := range dirs {
            next := Point{p.x + d.x, p.y + d.y}
            dfs(next, height+1)
        }
        visited[p] = false
    }
    
    dfs(start, 0)
    return paths
}

func solve2(grid [][]int) int {
    trailheads := findTrailheads(grid)
    sum := 0
    for _, th := range trailheads {
        sum += countPaths(grid, th)
    }
    return sum
}

func main() {
    lines := readLines("10/input.txt")
    grid := parseGrid(lines)
    result := solve2(grid)
    fmt.Println(result)
}