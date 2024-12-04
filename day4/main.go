// Package main of Advent of Code 2024 day4
package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "slices"
)

func readInput(input string) (grid []string) {
    file, err := os.Open(input)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        grid = append(grid, scanner.Text())
    }
    return
}

func checkDirection(grid []string, r, c, dr, dc int) bool {
    for _, char := range "XMAS" {
        if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
            return false
        }

        if rune(grid[r][c]) != char {
            return false
        }

        r += dr
        c += dc
    }
    return true
}

// CalculateXmas calculates the number of times the word XMAS appears in the grid
func CalculateXmas(grid []string) (sum int) {
    directions := [][2]int{
        {-1, 0}, {1, 0}, {0, -1}, {0, 1},
        {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
    }

    for r, row := range grid {
        for c := range row {
            for _, dir := range directions {
                dr, dc := dir[0], dir[1]
                if checkDirection(grid, r, c, dr, dc) {
                    sum++
                }
            }
        }
    }
    return
}

func checkXmasShape(grid []string, r, c int) bool {
    if grid[r][c] != 'A' {
        return false
    }

    diagonals := [][2]int{
        {r - 1, c - 1}, {r - 1, c + 1},
        {r + 1, c - 1}, {r + 1, c + 1},
    }

    possibleConfig := []string{
        "MMSS", "MSMS", "SSMM", "SMSM",
    }

    config:= ""
    for _, d := range diagonals {
        dr, dc := d[0], d[1]
        config += string(grid[dr][dc])
    }
    return slices.Contains(possibleConfig, config)
}

// CalculateXmasShape counts how many times the X-MAS shape appears in the grid
func CalculateXmasShape(grid []string) (sum int) {
    for r := 1; r < len(grid)-1; r++ {
        for c := 1; c < len(grid[r])-1; c++ {
            if checkXmasShape(grid, r, c) {
                sum++
            }
        }
    }
    return
}


func main() {
    grid := readInput("input.txt")
    fmt.Println("CalculateXmas returns", CalculateXmas(grid))
    fmt.Println("calculateXmasShape returns", CalculateXmasShape(grid))
}
