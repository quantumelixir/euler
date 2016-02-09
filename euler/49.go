package euler

import (
    "fmt"
    "euler/utils"
)

func E49() string {

    primes := utils.Primes(10000)
    var a, b, c int

    for i := 0; i < len(primes); i++ {
        for j := i + 1; j < len(primes); j++ {
            if utils.ArePermutations(primes[i], primes[j]) {
                k := 2*primes[j] - primes[i]
                if utils.IsPrime(k) && utils.ArePermutations(primes[i], k) {
                    a, b, c = primes[i], primes[j], k
                }
            }
        }
    }

    return fmt.Sprintf("%d%d%d", a, b, c)
}
