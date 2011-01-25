package main

import (
    "fmt"
)

func main() {
    const upto = 10000

    sumdivs := make([]int, upto + 1)
    for i := 1; i <= upto; i++ {
        for j := i*2; j <= upto; j+=i {
            sumdivs[j] += i
        }
    }

    sum := 0
    for i := 1; i <= upto; i++ {
        a := sumdivs[i]
        if a <= i || a > upto { continue }
        b := sumdivs[a]
        if b == i { sum += a + b }
    }

    fmt.Println(sum)
}
