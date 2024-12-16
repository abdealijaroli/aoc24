package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct {
	x, y int
}

func readInput(filename string) map[P]uint8 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	m := make(map[P]uint8)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			m[P{x, y}] = uint8(char)
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return m
}

// BFS exploration function
func explore(m map[P]uint8, p P, visited map[P]struct{}) (int, int) {
	visited[p] = struct{}{}

	id := m[p]
	area, corners := 0, 0

	var c P
	toDo := []P{p}
	for len(toDo) > 0 {
		c, toDo = toDo[0], toDo[1:]

		area++

		for _, d := range []P{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := P{c.x + d.x, c.y + d.y}
			val, ok := m[next]
			if ok && val == id {
				if _, ok := visited[next]; !ok {
					toDo = append(toDo, next)
					visited[next] = struct{}{}
				}
			}
		}

		for _, d := range []P{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}} {
			// convex corner
			if m[P{c.x + d.x, c.y}] != id &&
				m[P{c.x, c.y + d.y}] != id {
				corners++
			}
			// concave corner
			if m[P{c.x + d.x, c.y}] == id &&
				m[P{c.x, c.y + d.y}] == id &&
				m[P{c.x + d.x, c.y + d.y}] != id {
				corners++
			}
		}
	}

	return area, corners
}

func solve(m map[P]uint8) (int, int) {
	visited := make(map[P]struct{})
	var cost1, cost2 int

	for p := range m {
		if _, ok := visited[p]; !ok {
			area, corners := explore(m, p, visited)
			cost1 += area * 4
			cost2 += area * corners
		}
	}

	return cost1, cost2
}

func main() {
	m := readInput("12/input.txt")

	cost1, cost2 := solve(m)

	fmt.Println("Answer for part 1:", cost1)
	fmt.Println("Answer for part 2:", cost2)
}
