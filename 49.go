package main

import (
    "fmt"
    "utils"
)

func arePermutations(a, b int) (bool) {

    list := make([]int, 10)

    for a != 0 {
        list[a % 10]++
        a /= 10
    }
    for b != 0 {
        list[b % 10]--
        b /= 10
    }

    return utils.AbsSum(list) == 0
}

func main() {

    primes := utils.Primes(10000)
    var a, b, c int

    for i := 0; i < len(primes); i++ {
        for j := i + 1; j < len(primes); j++ {
            if arePermutations(primes[i], primes[j]) {
                k := 2*primes[j] - primes[i]
                if utils.IsPrime(k) && arePermutations(primes[i], k) {
                    a, b, c = primes[i], primes[j], k
                }
            }
        }
    }

    fmt.Printf("%d%d%d\n", a, b, c)
}
