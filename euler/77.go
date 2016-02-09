package euler

import (
    "euler/utils"
)

func ways77(n int, denom []int, ind int, memo [][]int) int {
    if memo[n][ind] != -1 { return memo[n][ind] }

    switch {
    case ind == len(denom) : memo[n][ind] = 0
    case denom[ind] > n    : memo[n][ind] = 0
    case denom[ind] == n   : memo[n][ind] = 1
    default:
        memo[n][ind] = ways77(n, denom, ind + 1, memo) +
                       ways77(n - denom[ind], denom, ind, memo)
    }

    return memo[n][ind]
}

func E77() (i int) {
    upto := 100
    denom := utils.Primes(upto + 1)
    memo := make([][]int, upto + 1)

    for i = 0; i <= upto; i++ {
        memo[i] = make([]int, len(denom) + 1)
        for j := 0; j <= len(denom); j++ { memo[i][j] = -1 }
    }

    for i = 0; ; i++ {
        if ways77(i, denom, 0, memo) > 5000 {
            break
        }
    }

    return i
}
