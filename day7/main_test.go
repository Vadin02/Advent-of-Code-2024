package main

import (
    "testing"
)


var got = []Data {
    {190, []int{10, 19}},
    {3267, []int{81, 40, 27}},
    {83, []int{17, 5}},
    {156, []int{15, 6}},
    {7290, []int{6, 8, 6, 15}},
    {161011, []int{16, 10, 13}},
    {192, []int{17, 8, 14}},
    {21037, []int{9, 7, 18, 13}},
    {292, []int{11, 6, 16, 20}},
}

func Test1(t *testing.T) {
    want := 3749

    if result := TotalCalibrationResult(got, false); result != want {
        t.Errorf("TotalCalibrationResult(%v, false) = %v, want %v", got, result, want)
    }
}

func Test2(t *testing.T) {
    want := 11387

    if result := TotalCalibrationResult(got, true); result != want {
        t.Errorf("TotalCalibrationResult(%v, true) = %v, want %v", got, result, want)
    }
}
