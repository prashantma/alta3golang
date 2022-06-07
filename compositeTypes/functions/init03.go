/* Alta3 Research | RZFeeser
   init - order of initialization  */

package main 

import "fmt"

// execution of "init" is done after declaration
func init() {
    WhatIsThe = 0
}
// declaration
var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
    fmt.Println("Called before init!!")
    return 42
}

func main() {
    if WhatIsThe == 0 {
        fmt.Println("The cake is a lie.")
    }
}

/*
Called before init!!
The cake is a lie.
*/
