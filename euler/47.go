package euler

func countDistinctFactors(limit int) []int {
    count := make([]int, limit)
    count[0], count[1] = -1, 1
    for i := 2; i < limit; i++ {
        if count[i] == 0 {
            for j := 2*i; j < limit; j += i {
                count[j]++
            }
        }
    }
    return count
}

func E47() int {
    const limit, k = 1E6, 4
    count := countDistinctFactors(limit)

    var i int
    for i = 0; i + k - 1 < limit ; i++ {
        flag := true
        for j := 0; flag && j < k; j++ {
            flag = flag && count[i + j] == k
        }
        if flag {
            break
        }
    }
    return i
}
