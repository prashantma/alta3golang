/* Alta3 Research | RZFeeser
   Refer - Basic patterns - panic(), defer, and recover()  */


package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    // encuntered only once in execution path
    defer func() {
        fmt.Println("Defer in f is called only once")
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }() // expression in defer must be function call
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
/*
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Defer in f is called only once
Recovered in f 4
Returned normally from f.
*/
