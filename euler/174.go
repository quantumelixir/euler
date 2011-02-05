package euler

func countLaminae(t int) int {
    count := 0
    for n := 1; ; n++ {
        m := t/4/n - n
        if m <= 0 { break }
        if 4*n*(m+n) == t { count++ }
    }
    return count
}

func E174() int {
    const N, upto = 15, 1E6
    count := 0
    for i := 1; i <= upto; i++ {
        c := countLaminae(i)
        if 1 <= c && c <= 10 { count++ }
    }
    return count
}
