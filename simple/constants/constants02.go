/* Alta3 Research | Author: RZFeeser
   Constants - Introduction to numerical constants   */

package main

import "fmt"

// this is a group of constants being declared at once
const (
    Code    = 585        // go will determine this is a "untyped" integer
    Big     = "New York"
    Country = "USA"
)

func oven(i string) string {
    return i
}

func main() {

    // display our global constants
    fmt.Println(Big)      // displays New York
    fmt.Println(Country)  // displays USA

    // Note: constants may also be lower case
    // they do not have to be uppercase
    const Small = 3    // this is an "untyped" integer
    fmt.Println(Small)

    // if uncommented, this will fail
    // Small = 10  // reassignment is NOT allowed

    // if uncommented, this will fail
    // the value of constants needs to be known at compile time
    // therefore values returned by function calls are illegal
    // ./constants02.go:36:21: oven("vanilla") (value of type string) is not constant
    const cupcake = oven("vanilla")
}

