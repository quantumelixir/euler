package euler

import (
    "euler/utils"
)

func E10() uint64 {
    var sum uint64
    sieve := utils.Sieve(2000000)
    for i, k := range sieve {
        if !k {
            sum += uint64(i)
        }
    }
    return sum
}
