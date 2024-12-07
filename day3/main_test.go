package main

import (
    "testing"
)


func Test1(t *testing.T) {
    got := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    want := 161

    if save:= SumOfMultiplication(&got); save != want {
        t.Errorf("SumOfMultiplication(%v) = %v, want %v", got, save, want)
    }
}

func Test2(t *testing.T) {
    got := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
    want := 48

    if save:= SumOfMultiplicationWithCondition(&got); save != want {
        t.Errorf("SumOfMultiplicationWithCOndition(%v) = %v, want %v", got, save, want)
    }
}
