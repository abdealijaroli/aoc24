package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type Block struct {
    fileID int
    size   int
    isFile bool
}

type FileBlock struct {
    fileID int
    start  int
    size   int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readLines(path string) string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    return scanner.Text()
}

func parseInput(input string) []Block {
    blocks := []Block{}
    fileID := 0
    isFile := true
    
    for i := 0; i < len(input); i++ {
        size, _ := strconv.Atoi(string(input[i]))
        blocks = append(blocks, Block{fileID, size, isFile})
        if isFile {
            fileID++
        }
        isFile = !isFile
    }
    return blocks
}

func expandBlocks(blocks []Block) []int {
    var disk []int
    
    for _, block := range blocks {
        for j := 0; j < block.size; j++ {
            if block.isFile {
                disk = append(disk, block.fileID)
            } else {
                disk = append(disk, -1)
            }
        }
    }
    return disk
}

func getFileBlocks(disk []int) []FileBlock {
    var files []FileBlock
    i := 0
    for i < len(disk) {
        if disk[i] != -1 {
            fileID := disk[i]
            start := i
            size := 0
            for i < len(disk) && disk[i] == fileID {
                size++
                i++
            }
            files = append(files, FileBlock{fileID, start, size})
        } else {
            i++
        }
    }
    return files
}

func compactDiskPart2(disk []int) []int {
    result := make([]int, len(disk))
    copy(result, disk)
    
    files := getFileBlocks(result)
    
    for i := len(files) - 1; i >= 0; i-- {
        file := files[i]
        bestPos := -1
        
        for pos := 0; pos < file.start; pos++ {
            if result[pos] == -1 {
                fits := true
                for j := 0; j < file.size; j++ {
                    if pos+j >= len(result) || result[pos+j] != -1 {
                        fits = false
                        break
                    }
                }
                if fits {
                    bestPos = pos
                    break
                }
            }
        }
        
        if bestPos != -1 {
            for j := 0; j < file.size; j++ {
                result[bestPos+j] = file.fileID
                result[file.start+j] = -1
            }
        }
    }
    
    return result
}

func calculateChecksum(disk []int) int {
    sum := 0
    for pos, fileID := range disk {
        if fileID != -1 {
            sum += pos * fileID
        }
    }
    return sum
}

func solve2(input string) int {
    blocks := parseInput(input)
    disk := expandBlocks(blocks)
    compactedDisk := compactDiskPart2(disk)
    return calculateChecksum(compactedDisk)
}

func main() {
    input := readLines("9/input.txt")
    result := solve2(input)
    fmt.Println(result)
}