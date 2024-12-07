// Package main of Advent of Code 2024 day2
package main

import (
	"regexp"
	"strconv"
    "log"
    "os"
    "fmt" 
)

func readInput(input string) (*string) {
    file, err := os.ReadFile(input)
    if err != nil {
        log.Fatal(err)
    }
    data := string(file)
    return &data
}

// SumOfMultiplication calculates the sum found by regex
func SumOfMultiplication(input *string) (sum int) {
    re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

    matches := re.FindAllStringSubmatch(*input, -1)

    for _, match := range matches {
        num1, _ := strconv.Atoi(match[1])
        num2, _ := strconv.Atoi(match[2])
        sum += num1 * num2
    }
    return
}

// SumOfMultiplicationWithCondition calculates the sum found by regex with do and don't
func SumOfMultiplicationWithCondition(input *string) (sum int) {
    re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
    matches := re.FindAllStringSubmatch(*input, -1)
	
    mulEnabled := true

    for _, match := range matches {
        switch match[0] {
            case "do()": 
                mulEnabled = true
            case "don't()": 
                mulEnabled = false
            default:
                if mulEnabled {
                    num1, _ := strconv.Atoi(match[1])
                    num2, _ := strconv.Atoi(match[2])
                    sum += num1 * num2
                }
        }
    }
    return
}

func main() {
    data := readInput("input.txt")
    fmt.Println("Sum of multiplication is:", SumOfMultiplication(data))
    fmt.Println("Sum of multiplication with condition is:", SumOfMultiplicationWithCondition(data))
}
