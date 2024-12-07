// Pakcage main of Advent of Code 2024 day7
package main

import (
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "fmt"
)

type Data struct {
    target int
    numbers []int
}

func readInput(input string) (data []Data) {
    file, err := os.Open(input)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ":")
        target, _ := strconv.Atoi(parts[0])
        numbersStr := strings.Split(parts[1], " ")
        
        var numbers []int
        for _, number := range numbersStr {
            number, _ := strconv.Atoi(number)
            numbers = append(numbers, number)
        }
        

        data = append(data, Data{target, numbers})
    }
    return
}

func canSolve(target int, numbers[]int, current int, index int, concat bool) bool {
    if index == len(numbers) {
        return current == target
    }

    if canSolve(target, numbers, current + numbers[index], index + 1, concat) {
        return true
    }

    if canSolve(target, numbers, current * numbers[index], index + 1,  concat) {
        return true
    }
    
    if concat {
        concatinated := strconv.Itoa(current) + strconv.Itoa(numbers[index])
        concatinatedInt, _ := strconv.Atoi(concatinated)
        if canSolve(target, numbers, concatinatedInt, index + 1,  concat) {
            return true
        }
    }
    return false
}

// TotalCalibrationResult calculates the sum of the target values that can be solved
func TotalCalibrationResult (data []Data, concat bool) (result int) {
    for _, check := range data {
        if canSolve(check.target, check.numbers, check.numbers[0], 1, concat) {
            result += check.target
        }

    }
    return
}

func main() {
    data := readInput("input.txt")
    fmt.Println("The sum of the target values that can be solved is:", TotalCalibrationResult(data, false))
    fmt.Println("The sum of the target values that can be solved including concatination is:", TotalCalibrationResult(data, true))
}


