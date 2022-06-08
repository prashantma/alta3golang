/* Alta3 Research | RZFeeser
   Calling the panic() function directly */
   
package main

import (
    "fmt"
)

func main() {

    // panic produces a quick exit
    panic("Jim, we have a problem.") // this line will be displayed

    fmt.Println("You will not even see this line. The panic creates a fast fail.")
}
/*
panic: Jim, we have a problem.

goroutine 1 [running]:
main.main()
        /home/student/static/alta3golang/idiomatic/panic/panic01.go:13 +0x27
exit status 2
*/
