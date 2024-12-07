package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

type State struct {
	pos Position
	dir int
}

func findLoopPositions(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	var startPos Position

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '^' {
				startPos = Position{x, y}
				grid[y][x] = '.' 
			}
		}
	}

	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}

	validPositions := 0

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == '.' && (x != startPos.x || y != startPos.y) {
				grid[y][x] = '#'

				if causesLoop(grid, startPos, dx, dy) {
					validPositions++
				}

				grid[y][x] = '.'
			}
		}
	}

	return validPositions
}

func causesLoop(grid [][]byte, startPos Position, dx, dy []int) bool {
	rows, cols := len(grid), len(grid[0])
	visited := make(map[State]int)

	pos := startPos
	dir := 0 
	steps := 0

	for steps < rows*cols*4 { 
		state := State{pos, dir}
		if count, seen := visited[state]; seen {
			return steps-count > 3 
		}
		visited[state] = steps

		nx := pos.x + dx[dir]
		ny := pos.y + dy[dir]

		if nx < 0 || nx >= cols || ny < 0 || ny >= rows {
			return false
		}

		if grid[ny][nx] == '#' {
			dir = (dir + 1) % 4
		} else {
			pos.x = nx
			pos.y = ny
		}

		steps++
	}

	return true 
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte

	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if len(grid) > 0 {
		fmt.Println(findLoopPositions(grid))
	}
}
