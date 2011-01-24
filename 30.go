package main

import (
    "fmt"
)

func sumOfPowers(n int) int {
    r := 0
    for n != 0 {
        t := n%10
        r += t*t*t*t*t
        n /= 10
    }
    return r
}

func main() {
    sum := 0
    for i := 2; i < 1E6; i++ {
        if sumOfPowers(i) == i { sum += i }
    }
    fmt.Println(sum)
}
