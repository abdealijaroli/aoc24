package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type C struct{ x, y int }

type Config struct {
	A, B, P C
}

func parseConfigs() []Config {
	scanner := bufio.NewScanner(os.Stdin)

	configs := []Config{}
	config := Config{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			config = Config{}
		}

		var x, y int
		switch {
		case strings.HasPrefix(line, "Button A"):
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
			config.A = C{x, y}

		case strings.HasPrefix(line, "Button B"):
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
			config.B = C{x, y}

		case strings.HasPrefix(line, "Prize"):
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
			config.P = C{x, y}
			configs = append(configs, config)
		}
	}
	return configs
}

func solve(c Config) int {
	b := (c.A.x*c.P.y - c.A.y*c.P.x) / (c.A.x*c.B.y - c.A.y*c.B.x)
	a := (c.P.x - b*c.B.x) / c.A.x

	if a*c.A.x+b*c.B.x == c.P.x && a*c.A.y+b*c.B.y == c.P.y {
		return 3*a + b
	}

	return 0
}

func main() {
	configs := parseConfigs()

	for i := range configs {
		configs[i].P.x += 10000000000000
		configs[i].P.y += 10000000000000
	}

	tokens := 0
	for _, config := range configs {
		tokens += solve(config)
	}
	fmt.Println(tokens)
}