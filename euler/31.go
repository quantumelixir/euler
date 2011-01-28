package euler

func ways31(n int, denom []int, ind int, memo [][]int) int {
    if memo[n][ind] != 0 { return memo[n][ind] }

    if ind == len(denom) { return 0 }
    if denom[ind] > n    { return 0 }
    if denom[ind] == n   { return 1 }

    memo[n][ind] = ways31(n, denom, ind + 1, memo) + ways31(n - denom[ind], denom, ind, memo)
    return memo[n][ind]
}

func E31() int {
    upto := 200
    denom := []int{1, 2, 5, 10, 20, 50, 100, 200}
    memo := make([][]int, upto + 1)
    for i := 0; i <= upto; i++ {
        memo[i] = make([]int, len(denom))
    }

    return ways31(upto, denom, 0, memo)
}
