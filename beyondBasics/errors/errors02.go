/* Alta3 Research | RZFeeser
   Errors - Creating basic errors with fmt.Errorf */

package main

import (
    "fmt"
    "time"
)

func main() {

    // build a error message with the current time
    err := fmt.Errorf("error occurred at: %v", time.Now()) // use a formatting verb to add info to the error
    fmt.Println("An error happened:", err)
}
/*
An error happened: error occurred at: 2022-06-09 13:35:02.685663177 +0000 UTC m=+0.000037340
*/
