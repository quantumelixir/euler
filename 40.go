package main

import (
    "fmt"
    "utils"
)

func getDigit(n, p int) (digit int) {
    if n == 0 { return 0 }

    t, tot := n, 0
    for t != 0 {
        t/=10
        tot++
    }

    p = tot - p + 1
    for p != 0 {
        digit = n%10
        n/=10
        p--
    }

    return digit
}

func d(n int) int {
    tot, i, ti := 0, 1, 1
    for {
        if n <= tot + 9*i*ti {
            rem := n - tot - 1
            return getDigit(rem/i + ti, rem%i + 1)
        }
        tot += 9*i*ti
        i++
        ti *= 10
    }
    return i
}

func main() {
    prod := 1
    for i := 0; i <= 6; i++ {
        prod *= d(utils.Pow(10, i))
    }
    fmt.Println(prod)
}
