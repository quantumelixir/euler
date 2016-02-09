package euler

import (
    "euler/utils"
)

func isPandigital(l, m, n int) bool {
    var used [10]int

    for l != 0 { used[l%10]++; l /= 10 }
    for m != 0 { used[m%10]++; m /= 10 }
    for n != 0 { used[n%10]++; n /= 10 }

    if used[0] != 0 { return false }

    for i := 1; i < 10; i++ {
        if used[i] != 1 { return false }
    }

    return true
}

func E32() (sum int) {

    products := make(map[int]bool)

    // i, j = number of digits in the multiplicand, multiplier
    // i + j + (i + j - 1) = 2*i + 2*j - 1 = 9 => i + j = 5
    for d := 1; d < 5; d++ {
        for i := utils.Pow(10, d - 1); i < utils.Pow(10, d); i++ {
            for j := utils.Pow(10, 5 - d - 1); i <= j && j < utils.Pow(10, 5 - d); j++ {
                if isPandigital(i, j, i*j) {
                    products[i*j] = true
                }
            }
        }
    }

    for k, _ := range products {
        sum += k
    }

    return sum
}
