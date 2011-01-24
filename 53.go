package main

import (
    "fmt"
)

func exceeds(n, r int, l float64) bool {
    if n < r { return false }
    c := 1.0
    for i := 1; i <= r && c <= l; i++ { c = c*float64(n-i+1)/float64(r-i+1) }
    if c > l { return true }
    return false
}

func main() {
    const upto = 100

    var sum int
    curr := make([]int, upto + 1)
    prev := make([]int, upto + 1)
    prev[0], prev[1] = 1, 1
    curr[0] = 1

    for n := 2; n <= upto; n++ {
        for r := 1; r <= n; r++ {
            curr[r] = prev[r-1] + prev[r]
            if curr[r] > 1E6 {
                curr[r] = 1E6
                sum += n - 2*r + 1
                break
            }
        }
        prev, curr = curr, prev
    }

    fmt.Println(sum)
}
