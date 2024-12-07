// Package main of Advent of Code 2024 day4
package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "slices"
    "strconv"
)

func readInput(input string) (*[]string , *[]string) {
    file, err := os.ReadFile(input)
    if err != nil {
        log.Fatal(err)
    }
    split := strings.Split(strings.TrimSpace(string(file)), "\n\n")
    rules := strings.Split(split[0], "\n")
    data := strings.Split(split[1], "\n")
    return &rules, &data
}


func compair(a, b string, rules *[]string) int {
    for _, rule := range *rules {
        if s := strings.Split(rule, "|"); s[0] == a && s[1] == b {
            return -1
        }
    }
    return 0
}

// Calculate calculates the sum of the middle element of each line sorted
func Calculate (correctlySorted bool, data *[]string, rules *[]string) (sum int) {
    for _, line := range *data {

        sortFunc := func(a, b string) int {
            return compair(a, b, rules)
        }

        if s := strings.Split(line, ","); slices.IsSortedFunc(s, sortFunc) == correctlySorted {
            if !correctlySorted {
                slices.SortFunc(s, sortFunc)
            }
            middle, _ := strconv.Atoi(s[len(s)/2])
            sum += middle
        }
    }
    return
}

func main() {
    rules, data := readInput("input.txt")
    fmt.Println("The middle value of all correcty sorted sumed up is:", Calculate(true, data, rules))
    fmt.Println("The middle value of all incorrecty sorted sumed up is:", Calculate(false, data, rules))
}
