package main

import (
    "bufio"
    "fmt"
    "os"
)

type Point struct {
    x, y int
}

type Antenna struct {
    pos       Point
    frequency rune
}

func isCollinear(p1, p2, p3 Point) bool {
    dx1 := p2.x - p1.x
    dy1 := p2.y - p1.y
    dx2 := p3.x - p1.x
    dy2 := p3.y - p1.y
    return dx1*dy2 == dx2*dy1
}

func findAntinodes(grid []string) int {
    var antennas []Antenna
    rows, cols := len(grid), len(grid[0])
    
    for y := range grid {
        for x, c := range grid[y] {
            if c != '.' {
                antennas = append(antennas, Antenna{Point{x, y}, c})
            }
        }
    }
    
    antinodes := make(map[Point]bool)
    
    for i := range antennas {
        for j := i + 1; j < len(antennas); j++ {
            if antennas[i].frequency != antennas[j].frequency {
                continue
            }
            
            antinodes[antennas[i].pos] = true
            antinodes[antennas[j].pos] = true
            
            for y := 0; y < rows; y++ {
                for x := 0; x < cols; x++ {
                    p := Point{x, y}
                    if isCollinear(p, antennas[i].pos, antennas[j].pos) {
                        antinodes[p] = true
                        // Check all other antennas of same frequency
                        for k := j + 1; k < len(antennas); k++ {
                            if antennas[k].frequency == antennas[i].frequency && 
                               isCollinear(p, antennas[i].pos, antennas[k].pos) {
                                antinodes[antennas[k].pos] = true
                            }
                        }
                    }
                }
            }
        }
    }
    
    return len(antinodes)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var grid []string
    
    for scanner.Scan() {
        grid = append(grid, scanner.Text())
    }
    
    fmt.Println(findAntinodes(grid))
}


// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
// )

// type Point struct {
//     x, y int
// }

// type Antenna struct {
//     pos       Point
//     frequency rune
// }

// func findAntinodes(grid []string) int {
//     var antennas []Antenna
    
//     for y := range grid {
//         for x, c := range grid[y] {
//             if c != '.' {
//                 antennas = append(antennas, Antenna{Point{x, y}, c})
//             }
//         }
//     }
    
//     antinodes := make(map[Point]bool)
    
//     for i := range antennas {
//         for j := i + 1; j < len(antennas); j++ {
//             if antennas[i].frequency == antennas[j].frequency {
//                 dx := antennas[j].pos.x - antennas[i].pos.x
//                 dy := antennas[j].pos.y - antennas[i].pos.y
                
//                 checkAntinode := func(mult int) {
//                     x1 := antennas[i].pos.x - dx*mult
//                     y1 := antennas[i].pos.y - dy*mult
//                     x2 := antennas[j].pos.x + dx*mult
//                     y2 := antennas[j].pos.y + dy*mult
                    
//                     if x1 >= 0 && x1 < len(grid[0]) && y1 >= 0 && y1 < len(grid) {
//                         antinodes[Point{x1, y1}] = true
//                     }
//                     if x2 >= 0 && x2 < len(grid[0]) && y2 >= 0 && y2 < len(grid) {
//                         antinodes[Point{x2, y2}] = true
//                     }
//                 }
                
//                 checkAntinode(1)
//             }
//         }
//     }
    
//     return len(antinodes)
// }

// func main() {
//     scanner := bufio.NewScanner(os.Stdin)
//     var grid []string
    
//     for scanner.Scan() {
//         grid = append(grid, scanner.Text())
//     }
    
//     fmt.Println(findAntinodes(grid))
// }