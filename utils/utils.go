package utils

func Min(a, b int) (int) {
    if a < b {
        return a
    }

    return b
}

func Max(a, b int) (int) {
    if a > b {
        return a
    }

    return b
}

func IsPrime(n int) (bool) {

    if n == 2 || n == 3 {
        return true
    }

    if n == 1 || n % 2 == 0 || n % 3 == 0 {
        return false
    }

    sqrt := 1
    for sqrt*sqrt < n {
        sqrt++
    }

    for j := 5; j <= sqrt; j += 6 {
        if n%j == 0 || n%(j+2) == 0 {
            return false
        }
    }

    return true
}

func RotateInt(n int) (int) {

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

func ReverseInt(n int) (int) {
    var res int
    for n != 0 {
        res = res*10 + n%10
        n /= 10
    }
    return res
}
