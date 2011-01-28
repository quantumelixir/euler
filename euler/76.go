package euler

func ways76(n int, from int, memo [][]int) int {
    if memo[n][from] != 0 { return memo[n][from] }

    if from > n    { return 0 }
    if from == n   { return 1 }

    memo[n][from] = ways76(n, from + 1, memo) + ways76(n - from, from, memo)
    return memo[n][from]
}

func E76() int {
    upto := 100

    memo := make([][]int, upto + 1)
    for i := 0; i <= upto; i++ {
        memo[i] = make([]int, upto + 1)
    }

    return ways76(upto, 1, memo) - 1 // exclude the single element "sum"
}
