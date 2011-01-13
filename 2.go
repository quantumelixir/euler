package main

import "fmt"

func main() {

    c, sum := 0, 0

    for a, b := 1,1; a < 4000000; c = a + b {
        a = b
        b = c
        c = a + b
        if a % 2 == 0 {
            sum += a
        }
    }

    fmt.Printf("%d\n", sum)
}
