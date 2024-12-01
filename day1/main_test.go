package main

import (
    "testing"
)

var gotLeft = []int{3, 4, 2, 1, 3, 3}
var gotRight = []int{4, 3, 5, 3, 9, 3}

func Test1(t *testing.T) {
    want := 11

    if difference := CalculateDifferences(gotLeft, gotRight); difference != want { 
        t.Errorf("CalculateDifferences(%v, %v) = %v, want %v", gotLeft, gotRight, difference, want)
    }
}


func Test2(t *testing.T) {
    want := 31

    if similarity := CalculateSimilarity(gotLeft, gotRight); similarity != want {
        t.Errorf("CalculateSimilarity(%v, %v) = %v, want %v", gotLeft, gotRight, similarity, want)
    }
}
