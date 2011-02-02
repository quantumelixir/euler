package main

import (
    "fmt"
)

func pow(a, b, mod uint64) uint64 {
    r := uint64(1)
    for b != 0 {
        if b & 0x1 == 1 {
            r = (r * a) % mod; b--
        }
        a = (a * a) % mod; b >>= 1
    }
    return r
}

func hyperexp(a, b, mod uint64) uint64 {
    if b == 1 { return a % mod }
    return pow(a, hyperexp(a, b - 1, mod), mod)
}

func main() {
    var a, b, mod uint64 = 1777, 1855, 1E8
    fmt.Println(hyperexp(a, b, mod))
}
