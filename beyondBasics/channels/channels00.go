/* Alta3 Research | RZFeeser
   Channels - Create a basic channel     */

package main
  
import "fmt"
  
func main() {
  
    // Creating a channel
    // Using var keyword - initializes with 'nil' - cannot transport data with us
    var mychannel chan int
    fmt.Println("Value of the channel: ", mychannel)
    fmt.Printf("Type of the channel: %T ", mychannel)
  
    // Creating a channel using make() function
    mychannel1 := make(chan int)
    fmt.Println("\nValue of the channel1: ", mychannel1)
    fmt.Printf("Type of the channel1: %T ", mychannel1)
}
/*
Value of the channel:  <nil>
Type of the channel: chan int
Value of the channel1:  0xc000062060
Type of the channel1: chan int
*/
