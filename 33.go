package main

import (
    "fmt"
)

func gcd(a, b int) int {
    if a < 0  { return gcd(-a, b) }
    if b < 0  { return gcd(a, -b) }
    if a < b  { return gcd(b, a) }
    if b == 0 { return a }
    return gcd(b, a%b)
}

func areEqual(a, b int, x, y int) bool {
    c, z := gcd(a, b), gcd(x, y)
    if c == 0 || z == 0 { return c == z }
    return a/c == x/z && b/c == y/z
}

func main() {
    n, d := 1, 1
    for i := 10; i < 100; i++ {
        for j := i + 1; j < 100; j++ {
            if i%10 == 0 || j%10 == 0 { continue }
            k, l := 1, 1
            if i/10 == j/10 && areEqual(i, j, i%10, j%10) { k, l = i%10, j%10 }
            if i/10 == j%10 && areEqual(i, j, i%10, j/10) { k, l = i%10, j/10 }
            if i%10 == j/10 && areEqual(i, j, i/10, j%10) { k, l = i/10, j%10 }
            if i%10 == j%10 && areEqual(i, j, i/10, j/10) { k, l = i/10, j/10 }
            if k != 1 || l != 1 { fmt.Println(i, j, k, l) }
            n *= k
            d *= l
        }
    }
    fmt.Println(d/gcd(n, d))
}
