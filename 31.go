package main

import (
    "fmt"
)

var memo [][]int

func ways(n int, denom []int, ind int) int {
    if memo[n][ind] != 0 { return memo[n][ind] }

    if ind == len(denom) { return 0 }
    if denom[ind] > n    { return 0 }
    if denom[ind] == n   { return 1 }

    memo[n][ind] = ways(n, denom, ind + 1) + ways(n - denom[ind], denom, ind)
    return memo[n][ind]
}

func main() {
    upto := 200
    denom := []int{1, 2, 5, 10, 20, 50, 100, 200}
    memo = make([][]int, upto + 1)
    for i := 0; i <= upto; i++ {
        memo[i] = make([]int, len(denom))
    }

    fmt.Println(ways(upto, denom, 0))
}
