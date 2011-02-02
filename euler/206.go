package main

import (
    "fmt"
)

func main() {
    var i, res uint64
    I: for i = 1E7; ; i+=10 {
        for _, j := range [2]uint64{3, 7} {
            s := (i + j) * (i + j) / 100
            if s % 10 != 8 { continue }
            s /= 100
            if s % 10 != 7 { continue }
            s /= 100
            if s % 10 != 6 { continue }
            s /= 100
            if s % 10 != 5 { continue }
            s /= 100
            if s % 10 != 4 { continue }
            s /= 100
            if s % 10 != 3 { continue }
            s /= 100
            if s % 10 != 2 { continue }
            s /= 100
            if s % 10 != 1 { continue }
            res = (i + j) * 10
            break I
        }
    }
    fmt.Println(res)
}
