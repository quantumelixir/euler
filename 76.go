package main

import (
    "fmt"
)

var memo [][]int

func ways(n int, from int) int {
    if memo[n][from] != 0 { return memo[n][from] }

    if from > n    { return 0 }
    if from == n   { return 1 }

    memo[n][from] = ways(n, from + 1) + ways(n - from, from)
    return memo[n][from]
}

func main() {
    upto := 100

    memo = make([][]int, upto + 1)
    for i := 0; i <= upto; i++ {
        memo[i] = make([]int, upto + 1)
    }

    fmt.Println(ways(upto, 1) - 1) // exclude the single element "sum"
}
