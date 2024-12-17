package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type robot struct {
	pos [2]int
	vel [2]int
}

func readRobots(lines []string) []robot {
	robots := []robot{}
	for _, line := range lines {
		splitResult1 := strings.Split(line, "=")
		splitResult2 := strings.Split(splitResult1[1], " ")
		positions := strings.Split(splitResult2[0], ",")
		velocities := strings.Split(splitResult1[2], ",")
		posX, _ := strconv.Atoi(positions[0])
		posY, _ := strconv.Atoi(positions[1])
		velX, _ := strconv.Atoi(velocities[0])
		velY, _ := strconv.Atoi(velocities[1])
		robots = append(robots, robot{pos: [2]int{posX, posY}, vel: [2]int{velX, velY}})
	}
	return robots
}

func printRobotMap(robots []robot, gridSize [2]int) {
	fmt.Println()
	for j := 0; j < gridSize[1]; j++ {
		for i := 0; i < gridSize[0]; i++ {
			robotCounter := 0
			for _, robot := range robots {
				if robot.pos == [2]int{i, j} {
					robotCounter++
				}
			}
			if robotCounter == 0 {
				fmt.Print(".")

			} else {
				fmt.Print(robotCounter)
			}
		}
		fmt.Println()
	}
}

func moveRobots(robots []robot, gridSize [2]int, seconds int) []robot {
	newRobots := []robot{}
	for _, currentRobot := range robots {
		newPosX := (currentRobot.pos[0] + currentRobot.vel[0]*seconds) % gridSize[0]
		newPosY := (currentRobot.pos[1] + currentRobot.vel[1]*seconds) % gridSize[1]
		if newPosX < 0 {
			newPosX += gridSize[0]
		}
		if newPosY < 0 {
			newPosY += gridSize[1]
		}
		newRobots = append(newRobots, robot{pos: [2]int{newPosX, newPosY}, vel: currentRobot.vel})
	}
	return newRobots
}

func calcSafetyFactor(robots []robot, gridSize [2]int) int {
	quadUpRight := 0
	quadUpLeft := 0
	quadDownRight := 0
	quadDownLeft := 0
	for _, robot := range robots {
		if robot.pos[0] < gridSize[0]/2 {
			if robot.pos[1] < gridSize[1]/2 {
				quadUpLeft++
			} else if robot.pos[1] > gridSize[1]/2 {
				quadUpRight++
			} else {
				continue
			}
		} else if robot.pos[0] > gridSize[0]/2 {
			if robot.pos[1] < gridSize[1]/2 {
				quadDownLeft++
			} else if robot.pos[1] > gridSize[1]/2 {
				quadDownRight++
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return quadDownLeft * quadDownRight * quadUpLeft * quadUpRight
}

func isThereATree(robots []robot, gridSize [2]int) bool {
	robotMap := map[[2]int]int{}
	for j := 0; j < gridSize[1]; j++ {
		for i := 0; i < gridSize[0]; i++ {
			robotMap[[2]int{i, j}] = 0
			for _, robot := range robots {
				if robot.pos == [2]int{i, j} {
					robotMap[[2]int{i, j}]++
				}
			}
		}
	}
	maxVerticalSequence := 0
	for i := 0; i < gridSize[0]; i++ {
		verticalSequence := 0
		for j := 0; j < gridSize[1]; j++ {
			if robotMap[[2]int{i, j}] >= 1 {
				verticalSequence++
			}
		}
		if verticalSequence > maxVerticalSequence {
			maxVerticalSequence = verticalSequence
		}
	}
	maxHorizontalSequence := 0
	for j := 0; j < gridSize[1]; j++ {
		horizontalSequence := 0
		for i := 0; i < gridSize[0]; i++ {
			if robotMap[[2]int{i, j}] >= 1 {
				horizontalSequence++
			}
		}
		if horizontalSequence > maxHorizontalSequence {
			maxHorizontalSequence = horizontalSequence
		}
	}
	if maxVerticalSequence > 20 && maxHorizontalSequence > 20 {
		return true
	} else {
		return false
	}
}

func main() {
	var gridSize = [2]int{101, 103}

	lines := readLines("14/input.txt")
	initialRobots := readRobots(lines)
	robots := moveRobots(initialRobots, gridSize, 100)
	// safetyFactor := calcSafetyFactor(robots, gridSize)

	for time := 1; time < 10000; time++ {
		robots = moveRobots(robots, gridSize, 1)
		if isThereATree(robots, gridSize) {
			printRobotMap(robots, gridSize)
			fmt.Println("seconds: ", 100+time)
		}
	}
}
