package main

import (
    "fmt"
    "utils"
)

func mostConsecutivePrimeBelow(n int) (int) {

    sieve := utils.Sieve(n)
    count := 0

    // get all primes below n
    for _, i := range sieve { if !i { count++ } }
    k, primes := 0, make([]int, count)
    for i, j := range sieve {
        if !j {
            primes[k] = i
            k++
        }
    }

    // find the prime expressible as the sum
    // of most number of consecutive primes
    prime, max := 0, 0
    for i := 0; i < count; i++ {
        for sum, j := primes[i], i; j + 1 < count && sum < n; j++ {
            if !sieve[sum] && j + 1 - i > max {
                max = j + 1 - i
                prime = sum
            }
            sum += primes[j + 1]
        }
    }

    return prime
}

func main() {
    fmt.Println(mostConsecutivePrimeBelow(1000000))
}
