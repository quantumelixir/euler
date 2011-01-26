package main

import (
    "fmt"
    "utils"
)

var memo [][]int

func ways(n int, denom []int, ind int) int {
    if memo[n][ind] != -1 { return memo[n][ind] }

    switch {
    case ind == len(denom) : memo[n][ind] = 0
    case denom[ind] > n    : memo[n][ind] = 0
    case denom[ind] == n   : memo[n][ind] = 1
    default:
        memo[n][ind] = ways(n, denom, ind + 1) +
                       ways(n - denom[ind], denom, ind)
    }

    return memo[n][ind]
}

func main() {
    upto := 100
    denom := utils.Primes(upto + 1)
    memo = make([][]int, upto + 1)
    for i := 0; i <= upto; i++ {
        memo[i] = make([]int, len(denom) + 1)
        for j := 0; j <= len(denom); j++ { memo[i][j] = -1 }
    }

    for i := 0; ; i++ {
        if ways(i, denom, 0) > 5000 {
            fmt.Println(i)
            break
        }
    }
}
