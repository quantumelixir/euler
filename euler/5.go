package euler

import (
    "utils"
)

func E5() int {
    res := 1
    for i := 1; i <= 20; i++ {
        res *= i / utils.GCD(res, i)
    }
    return res
}
