package euler

func E15() uint64 {
    const n = 21
    ways := make([][]uint64, n)
    for i := 0; i < n; i++ {
        ways[i] = make([]uint64, n)
        ways[i][0], ways[0][i] = 1, 1
    }

    for i := 1; i < n; i++ {
        for j := 1; j < n; j++ {
            ways[i][j] = ways[i][j-1] + ways[i-1][j]
        }
    }

    return ways[n-1][n-1]
}
