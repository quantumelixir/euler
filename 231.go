package main

import (
    "fmt"
)

var primes []uint64

func sieve(n uint64) (sieve []bool) {
    sieve = make([]bool, n)
    sieve[0], sieve[1] = true, true
    for i := uint64(2); i < n; i++ {
        for j := i*2; !sieve[i] && j < n; j += i {
            sieve[j] = true
        }
    }
    return sieve
}

func findprimes(n uint64) (primes []uint64) {

    count, sieve := 0, sieve(n)
    for _, i := range sieve { if !i { count++ } }

    k, primes := uint64(0), make([]uint64, count)
    for i := uint64(0); i < uint64(len(sieve)); i++ {
        if !sieve[i] {
            primes[k] = i
            k++
        }
    }
    return primes
}

func clever(n uint64) (sum uint64) {
    for _, v := range primes {
        for t := v; n/t != 0; t *= v {
            sum += (n/t)*v
        }
    }
    return sum
}

func main() {
    const n, r = 2E7, 15E6
    primes = findprimes(n)
    sum := clever(n) - clever(n-r) - clever(r)
    fmt.Println(sum)
}
