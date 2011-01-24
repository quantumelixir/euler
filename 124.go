package main

import (
    "fmt"
    "sort"
    "utils"
)

func pow(n, k int) int {
    c := 1
    for i := 1; i <= k; i++ { c *= n }
    return c
}

type Tuple [2]int
type TupleArray []Tuple

func (t TupleArray) Len() int {
    return len(t)
}

func (t TupleArray) Less(i, j int) bool {
    if t[i][1] == t[j][1] { return t[i][0] < t[j][0] }
    return t[i][1] < t[j][1]
}

func (t TupleArray) Swap(i, j int) {
    t[i], t[j] = t[j], t[i]
}

func main() {
    const upto = 100000
    primes := utils.Primes(upto)

    rad := make(TupleArray, upto + 1)
    for i := 1; i <= upto; i++ {
        rad[i][0], rad[i][1] = i, i
    }

    for _, i := range primes {
        for k := 2; pow(i, k) <= upto && pow(i, k) > 0; k++ {
            for j := pow(i, k); j <= upto; j += pow(i, k) {
                rad[j][1] /= i
            }
        }
    }

    sort.Sort(rad)
    fmt.Println(rad[10000][0])
}
