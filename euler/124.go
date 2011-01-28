package euler

import (
    "sort"
    "utils"
)

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

func E124() int {
    const upto = 100000
    primes := utils.Primes(upto)

    rad := make(TupleArray, upto + 1)
    for i := 1; i <= upto; i++ {
        rad[i][0], rad[i][1] = i, i
    }

    for _, i := range primes {
        for k := 2; utils.Pow(i, k) <= upto && utils.Pow(i, k) > 0; k++ {
            for j := utils.Pow(i, k); j <= upto; j += utils.Pow(i, k) {
                rad[j][1] /= i
            }
        }
    }

    sort.Sort(rad)
    return rad[10000][0]
}
