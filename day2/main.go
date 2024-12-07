// Package main of Advent of Code 2024 day2
package main

import (
    "fmt"
    "bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(input string) (reports [][]int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)
		report := make([]int, len(levels))

		for i := range levels {
			num, err := strconv.Atoi(levels[i])
			if err != nil {
				log.Fatal(err)
			}
			report[i] = num
		}
		reports = append(reports, report)
	}
	return
}

// CalculateSave calculates the number of reports that are save
func CalculateSave(reports [][]int) (totalSave int) {
	for _, report := range reports {
	    if isSave(report) {
            totalSave++
        }
    }
    return
}

// CalculateSaveDampener calculates the number of reports that are save with one number removed
func CalculateSaveDampener(reports [][]int) (totalSave int) {
    for _, report := range reports {
        if isSave(report) || isSaveDampener(report) {
            totalSave++
        }
    }
    return
}

func isSave(report []int) bool {
    increasing := false
    decreasing := false

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

        switch {
        case diff == 0: return false
        case diff < 0: decreasing = true
        case diff > 0: increasing = true
        }

		switch {
        case increasing:
			if diff < 1 || diff > 3 {
				return false
			}
        case decreasing:
			if diff > -1 || diff < -3 {
				return false
			}
		}
        if increasing && decreasing {
            return false
        }
	}
    return true
}

func isSaveDampener(report []int) bool {
    for i := range report {
        newReport := append([]int(nil), report[:i]...)
        newReport = append(newReport, report[i+1:]...)
        if isSave(newReport) {
            return true
        }
    }
    return false
}

func main() {
    data := readInput("input.txt")
    fmt.Println("Total number of reports that are save:", CalculateSave(data))
    fmt.Println("Total number of reports that are save with one number removed:", CalculateSaveDampener(data))
}
