package main

func IsPrime(n int) bool {
    if n < 2 {
        return false
    }
    for k := 2; k < n; k++ {
        if n%k == 0 {
            return false
        }
    }
    return true
}

func main() {
    // Example usage
    println(IsPrime(6))
    println(IsPrime(101))
    println(IsPrime(11))
    println(IsPrime(13441))
    println(IsPrime(61))
    println(IsPrime(4))
    println(IsPrime(1))
}
