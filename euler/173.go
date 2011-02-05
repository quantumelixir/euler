package euler

func E173() int {
    // (m + 2*n)^2 - m^2 <= tiles

    const tiles = 1E6
    count := 0
    for n := 1; ; n++ {
        m := tiles/4/n - n
        if m <= 0 { break }
        count += m
    }

    return count
}
