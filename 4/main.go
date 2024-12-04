package main

import (
	"fmt"
	"os"
	"strings"
)

func isXMAS(r, c int, arr []string) bool {
	val11 := arr[r-1][c-1]
	val12 := arr[r-1][c+1]
	val21 := arr[r+1][c-1]
	val22 := arr[r+1][c+1]

	str1 := string(val11) + string(val12)
	str2 := string(val21) + string(val22)

	if str1 == str2 && (str1 == "MS" || str1 == "SM") {
		return true
	}
	if str1 != str2 && ((str1 == "MM" && str2 == "SS") || (str1 == "SS" && str2 == "MM")) {
		return true
	}
	return false
}

func main() {
	data, err := os.ReadFile("4/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	arr := strings.Split(string(data), "\n")
	res := 0

	for r := 1; r < len(arr)-1; r++ {
		for c := 1; c < len(arr[0])-1; c++ {
			if arr[r][c] == 'A' && isXMAS(r, c, arr) {
				res++
			}
		}
	}

	fmt.Println(res)
}
