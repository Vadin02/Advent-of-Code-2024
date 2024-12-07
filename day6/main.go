// Package main of Advent of Code 2024 day6
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	Up    = [2]int{-1, 0} // Up facing
	Right = [2]int{0, 1} // Right facing
	Down  = [2]int{1, 0} // Down facing
	Left  = [2]int{0, -1} // Left facing
)

func readInput(input string) (grid [][]rune) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return
}

// CalculateFields changes where solder goes and now many field are visited
func CalculateFields(grid [][]rune, shouldPrint bool) (loop bool, result int) {
	grid = copyGrid(grid)
    var x, y int
	var dir [2]int
    visited := make(map[[2]int][2]int)



	rows, cols := len(grid), len(grid[0])
	for r, row := range grid {
		for c, col := range row {
			switch col {
			case '^':
				x, y, dir = r, c, Up
			case '>':
				x, y, dir = r, c, Right
			case 'v':
				x, y, dir = r, c, Down
			case '<':
				x, y, dir = r, c, Left
			}
		}
	}

    grid[x][y] = 'X'
    result++

	for {
		nx, ny := x+dir[0], y+dir[1]
		if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
            break
        }

        if grid[nx][ny] == '#' {
			switch dir {
			case Up:
				dir = Right
			case Right:
				dir = Down
			case Down:
				dir = Left
			case Left:
				dir = Up
            }
            continue
        }

        if visited[[2]int{nx, ny}] == dir {
            return true, result
        }

        visited[[2]int{nx, ny}] = dir

        x, y = nx, ny
        if grid[x][y] != 'X' {
            result++
            grid[x][y] = 'X'
        }
	}
    if shouldPrint {
        for _, row := range grid {
            fmt.Println(string(row))
        }
    }
	return false, result
}

// FindLoops finds loops in the grid where the guard gets stuck
func FindLoops(grid [][]rune) (result int) {
    grid = copyGrid(grid)
    for r, row := range grid {
        for c := range row {
            if grid[r][c] == '.' {
                grid[r][c] = '#'
                if loop, _ := CalculateFields(copyGrid(grid), false); loop {
                    result++
                    grid[r][c] = '0'
                } else {
                    grid[r][c] = '.'
                }
            }
        }
    }
    for _, row := range grid {
        fmt.Println(string(row))
    }
    return
}
func copyGrid(grid [][]rune) [][]rune {
    newGrid := make([][]rune, len(grid))
    for i := range grid {
        newGrid[i] = append([]rune(nil), grid[i]...)
    }
    return newGrid
}

func main() {
    grid := readInput("input.txt")
    _, result := CalculateFields(grid, true)
    fmt.Println("The solder visited", result, "fields before leaving the grid")
    loop := FindLoops(grid)
    fmt.Println("There were", loop, "ways found to create a loop")
}
