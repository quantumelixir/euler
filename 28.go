package main

import "fmt"

func main() {

    size, sum := 1001, 1
    for i := 3; i <= size; i+=2 {
        sum += i*i + i*i - (i - 1) + i*i - 2*(i - 1) + i*i - 3*(i - 1)
    }

    fmt.Printf("%d\n", sum)
}
