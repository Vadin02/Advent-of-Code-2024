package main

import (
    "testing"
)


var got = [][]rune{
{'.','.','.','.','#','.','.','.','.','.'},
{'.','.','.','.','.','.','.','.','.','#'},
{'.','.','.','.','.','.','.','.','.','.'},
{'.','.','#','.','.','.','.','.','.','.'},
{'.','.','.','.','.','.','.','#','.','.'},
{'.','.','.','.','.','.','.','.','.','.'},
{'.','#','.','.','^','.','.','.','.','.'},
{'.','.','.','.','.','.','.','.','#','.'},
{'#','.','.','.','.','.','.','.','.','.'},
{'.','.','.','.','.','.','#','.','.','.'},
}

func Test1(t *testing.T) {
    want := 41
    if _, result := CalculateFields(got, true); result != want {
        t.Errorf("CalculateFields() = %v, want %v", result, want)
    }
}

func Test2(t *testing.T) {
    want := 6
    if result := FindLoops(got); result != want {
        t.Errorf("FindLoops() = %v, want %v", result, want)
}
}
