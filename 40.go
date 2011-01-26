package main

import (
    "fmt"
    "utils"
)

func d(n int) int {
    tot, i, ti, digit := 0, 1, 1, 0
    for {
        if n <= tot + 9*i*ti {
            rem := n - tot - 1
            d, p := rem/i + ti, i - rem%i
            for p != 0 {
                digit = d%10
                d/=10
                p--
            }
            break
        }
        tot += 9*i*ti
        i++
        ti *= 10
    }
    return digit
}

func main() {
    prod := 1
    for i := 0; i <= 6; i++ {
        prod *= d(utils.Pow(10, i))
    }
    fmt.Println(prod)
}
