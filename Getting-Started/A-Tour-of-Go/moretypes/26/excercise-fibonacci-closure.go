package main

import "fmt"

func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        actual := a
        a, b = b, a+b
        return actual
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

