package euler

import (
    "utils"
)

func isRightTruncatablePrime(n int) (bool) {
    for n != 0 {
        if !utils.IsPrime(n) {
            return false
        }
        n /= 10
    }
    return true
}

func isLeftTruncatablePrime(n int) (bool) {
    n = utils.ReverseInt(n)
    for n != 0 {
        if !utils.IsPrime(utils.ReverseInt(n)) {
            return false
        }
        n /= 10
    }
    return true
}

func isTruncatablePrime(n int) (bool) {
    return isRightTruncatablePrime(n) && isLeftTruncatablePrime(n)
}

func E37() int {

    sum := 0
    for count, i := 0, 10; count < 11; i++ {
        if isTruncatablePrime(i) {
            sum += i
            count++
        }
    }

    return sum
}
