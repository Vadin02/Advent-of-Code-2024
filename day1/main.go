// Package main of Advent of Code 2024 day1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func readInput(input string) (left []int, right []int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l, r int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, l)
		right = append(right, r)
	}
	return
}

// CalculateDifferences calculates the total distance between two slices of integers as sum from lowest to highest
func CalculateDifferences(left []int, right []int) (totalDistance int) {
	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		if left[i] > right[i] {
			totalDistance += left[i] - right[i]
		} else {
			totalDistance += right[i] - left[i]
		}
	}
    return
}
// CalculateSimilarity calculates the similarity between two slices
func CalculateSimilarity(left []int, right []int) (totalSimilarity int) {
    frequencyMap := make(map[int]int)
    for _, num := range right {
        frequencyMap[num]++
    }

    for _, num := range left {
        totalSimilarity += num * frequencyMap[num]
    }

    return 
}

func main() {
    left, right := readInput("input.txt")
    fmt.Println("Total distance between the two slices is:", CalculateDifferences(left, right))
    fmt.Println("Total similarity between the two slices is:", CalculateSimilarity(left, right))
}
