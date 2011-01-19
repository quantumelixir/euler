package main

import (
    "fmt"
    "utils"
)

func main() {
    var sum uint64
    sieve := utils.Sieve(2000000)
    for i, k := range sieve {
        if !k {
            sum += uint64(i)
        }
    }
    fmt.Println(sum)
}
