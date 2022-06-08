/* Alta3 Research | RZFeeser
   Refer - The panic(), defer, and recover() */

package main

import (  
    "fmt"
)

// similar to "finally" in Python/Java
// Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect.
// If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
func recoverFullName() {  
    if r := recover(); r!= nil {
        fmt.Println("recovered from ", r)
    }
}

func fullName(firstName *string, lastName *string) {  
    defer recoverFullName()
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("will never print this line: `:returned normally from fullName")
}

func main() {  
    defer fmt.Println("deferred call in main")
    firstName := "Elon"
    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}
/*
recovered from  runtime error: last name cannot be nil
returned normally from main
deferred call in main
*/
