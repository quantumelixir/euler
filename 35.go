package main

import "fmt"

func rotate(n int) (int){

    sign := 1
    if n < 0 {
        sign = -1
        n = -n
    }

    r := n % 10
    p := 1
    for p < n {
        p *= 10
    }

    return sign*(r*p + n)/10
}

func isCircularPrime(n int, list []bool) (bool) {

    if list[n] {
        return false
    }

    orig := n
    for n = rotate(n); n != orig; n = rotate(n) {
        if list[n] {
            return false
        }
    }

    return true
}

func main() {

    limit := 1000000
    list := make([]bool, limit)
    list[0], list[1] = true, true

    for i := 2; i < len(list); i++ {
        if !list[i] {
            for j := i*2; j < len(list); j+=i {
                list[j] = true
            }
        }
    }

    count := 0
    for i, _ := range list {
        if isCircularPrime(i, list) {
            count++
        }
    }

    fmt.Printf("%d\n", count)
}
