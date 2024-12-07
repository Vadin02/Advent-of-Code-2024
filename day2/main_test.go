package main

import (
    "testing"
)

var got = [][]int {
    {7, 6, 4, 2, 1},
    {1, 2, 7, 8, 9},
    {9, 7, 6, 2, 1},
    {1, 3, 2, 4, 5},
    {8, 6, 4, 4, 1},
    {1, 3, 6, 7, 9}}

func Test1(t *testing.T) {
    want := 2

    if save := CalculateSave(got); save != want {
        t.Errorf("CalculateSave(%v) = %v, want %v", got, save, want)
    }
}

func Test2(t *testing.T) {
    want := 4
    
    if save:= CalculateSaveDampener(got); save != want {
        t.Errorf("CalculateSaveDampener(%v) = %v, want %v", got, save, want)
    }
}
