package main

import (
    "testing"
)

var got = []string {
    "MMMSXXMASM",
    "MSAMXMSMSA",
    "AMXSXMAAMM",
    "MSAMASMSMX",
    "XMASAMXAMM",
    "XXAMMXXAMA",
    "SMSMSASXSS",
    "SAXAMASAAA",
    "MAMMMXMMMM",
    "MXMXAXMASX",
}

func Test1(t *testing.T) {
    want := 18
    
    if save:= CalculateXmas(got); save != want {
        t.Errorf("CalculateXmas(%v) = %v, want %v", got, save, want)
    }
}

func Test2(t *testing.T) {
    want := 9
    
    if save:= CalculateXmasShape(got); save != want {
        t.Errorf("CalculateXmasShape(%v) = %v, want %v", got, save, want)
    }
}
