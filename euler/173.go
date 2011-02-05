package main

import (
    "fmt"
)

func main() {
    // m^2 - (m-2)^2 = 4m - 4 = tiles
    const tiles = 1E6
    const max = (tiles + 4)/4

    count, start := 0, 0
    for m := 2; m <= max; m+=2 {
        if m % 2 == 0 { start = m - 2
        } else        { start = m - 1 }
        for n := start; m*m - n*n <= tiles && n > 0; n-=2 {
            count++
        }
    }

    for m := 1; m <= max; m+=2 {
        if m % 2 == 0 { start = m - 1
        } else        { start = m - 2 }
        for n := start; m*m - n*n <= tiles && n > 0; n-=2 {
            count++
        }
    }

    fmt.Println(count)
}
