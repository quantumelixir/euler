package euler

import (
    "euler/utils"
)

func E52() (i int) {
    for i = 1; ; i++ {
        if  utils.ArePermutations(i, 2*i) &&
            utils.ArePermutations(i, 3*i) &&
            utils.ArePermutations(i, 4*i) &&
            utils.ArePermutations(i, 5*i) &&
            utils.ArePermutations(i, 6*i) {
            break;
        }
    }
    return i
}
