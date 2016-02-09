package euler

import (
    "euler/utils"
)

func isCircularPrime(n int, list []bool) (bool) {

    if list[n] {
        return false
    }

    orig := n
    for n = utils.RotateInt(n); n != orig; n = utils.RotateInt(n) {
        if list[n] {
            return false
        }
    }

    return true
}

func E35() int {

    limit := 1000000
    list := make([]bool, limit)
    list[0], list[1] = true, true

    for i := 2; i < len(list); i++ {
        if !list[i] {
            for j := i*2; j < len(list); j+=i {
                list[j] = true
            }
        }
    }

    count := 0
    for i, _ := range list {
        if isCircularPrime(i, list) {
            count++
        }
    }

    return count
}
