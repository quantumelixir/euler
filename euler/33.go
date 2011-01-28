package euler

import (
    "utils"
)

func areEqual(a, b int, x, y int) bool {
    c, z := utils.GCD(a, b), utils.GCD(x, y)
    if c == 0 || z == 0 { return c == z }
    return a/c == x/z && b/c == y/z
}

func E33() int {
    n, d := 1, 1
    for i := 10; i < 100; i++ {
        for j := i + 1; j < 100; j++ {
            if i%10 == 0 || j%10 == 0 { continue }
            k, l := 1, 1
            if i/10 == j/10 && areEqual(i, j, i%10, j%10) { k, l = i%10, j%10 }
            if i/10 == j%10 && areEqual(i, j, i%10, j/10) { k, l = i%10, j/10 }
            if i%10 == j/10 && areEqual(i, j, i/10, j%10) { k, l = i/10, j%10 }
            if i%10 == j%10 && areEqual(i, j, i/10, j/10) { k, l = i/10, j/10 }
            n *= k
            d *= l
        }
    }
    return d/utils.GCD(n, d)
}
