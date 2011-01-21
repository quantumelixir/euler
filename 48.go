package main

import (
    "fmt"
)

func raiseMod(a, b, m uint64) (r uint64) {
    var i uint64
    for i, r = 1, 1; i <= b; i++ {
        r = (r * a) % m
    }
    return r
}

func main() {
    var mod, res, i uint64

    mod = 10000000000
    for i = 1; i <= 1000; i++ {
        res = (res + raiseMod(i, i, mod)) % mod
    }

    fmt.Println(res)
}
