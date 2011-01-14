package main

import "fmt"

func main() {

    var prime int
    nth := 10001

    if nth == 1 {
        prime = 2
    }

    sqrt := 1
    I: for p, i := 1, 3; p < nth; i+=2 {
        for sqrt*sqrt < i {
            sqrt++
        }
        for j := 2; j < sqrt + 1; j++ {
            if i % j == 0 {
                continue I
            }
        }
        prime = i
        p++
    }

    fmt.Printf("%d\n", prime)
}
