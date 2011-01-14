package main

import "fmt"

func main() {

    var prime int
    nth := 10001

    switch nth {
        case 1: prime = 2
        case 2: prime = 3
    }

    sqrt := 1
    I: for p, i := 2, 5; p < nth; i+=2 {
        for sqrt*sqrt < i {
            sqrt++
        }
        if i % 3 == 0 {
            continue I
        }
        for j := 5; j <= sqrt; j+=6 {
            if i % j == 0 || i % (j + 2) == 0 {
                continue I
            }
        }
        prime = i
        p++
    }

    fmt.Printf("%d\n", prime)
}
