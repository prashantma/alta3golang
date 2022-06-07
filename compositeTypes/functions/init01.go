/* Alta3 Research | RZFeeser
   init function              */

package main

import "fmt"

/*
Useful in cases:
grab passwords, ssh key
connecting to database
critical services like heartbeats
*/
func init() {
  fmt.Println("Think of the init function like a prelude to main.")
}

func main() {
  fmt.Println("This would run 'after' the init.")
}
/*
Think of the init function like a prelude to main.
This would run 'after' the init.
*/
