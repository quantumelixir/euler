package euler

import (
    "utils"
)

func E12() (ans int) {

    limit := 500

    primes := utils.Primes(3)
    for i := 1; ; i++ {
        if primes[len(primes)-1] <= i + 1 {
            primes = utils.Primes(2*(i + 1))
        }
        n, d := i*(i+1)/2, 1
        for t, j := 0, 0; n != 1; j++ {
            for t = 0; n % primes[j] == 0; {
                n /= primes[j]
                t++
            }
            d *= t + 1
        }
        if d > limit {
            ans = i*(i+1)/2
            break
        }
    }

    return ans
}
