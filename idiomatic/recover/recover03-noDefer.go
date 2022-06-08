/* Alta3 Research | RZFeeser
   Refer - panic() and Defer (only)  */


package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")      // we will never see this message
}

func f() {
    // removed the defer function

    fmt.Println("Calling g.")
    g(0)
    
    // we will never reach this point
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))       // this time, will be displayed
    }
    defer fmt.Println("Defer in g", i)    // defer statements will still be processed
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
panic: 4

goroutine 1 [running]:
main.g(0x4)
        /home/student/static/alta3golang/idiomatic/recover/recover03-noDefer.go:27 +0x1dd
main.g(0x3)
        /home/student/static/alta3golang/idiomatic/recover/recover03-noDefer.go:31 +0x127
main.g(0x2)
        /home/student/static/alta3golang/idiomatic/recover/recover03-noDefer.go:31 +0x127
main.g(0x1)
        /home/student/static/alta3golang/idiomatic/recover/recover03-noDefer.go:31 +0x127
main.g(0x0)
        /home/student/static/alta3golang/idiomatic/recover/recover03-noDefer.go:31 +0x127
main.f()
        /home/student/static/alta3golang/idiomatic/recover/recover03-noDefer.go:18 +0x5d
main.main()
        /home/student/static/alta3golang/idiomatic/recover/recover03-noDefer.go:10 +0x19
exit status 2
*/
