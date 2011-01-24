package main

import (
    "fmt"
    "utils"
)

func main() {
    for i := 1; ; i++ {
        if  utils.ArePermutations(i, 2*i) &&
            utils.ArePermutations(i, 3*i) &&
            utils.ArePermutations(i, 4*i) &&
            utils.ArePermutations(i, 5*i) &&
            utils.ArePermutations(i, 6*i) {
            fmt.Println(i)
            break;
        }
    }
}
