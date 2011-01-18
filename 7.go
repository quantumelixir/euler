package main

import (
    "fmt"
    "utils"
)

func main() {

    var i int
    nth := 10001

    for count := 0; count != nth; i++ {
        if utils.IsPrime(i) {
            count++
        }
    }

    fmt.Println(i - 1)
}
