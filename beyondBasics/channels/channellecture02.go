/* Alta3 Research | RZFeeser
   Channels - A simple example.
   Think of a channel like a wormhole. Anything placed into it,
   is instantly moved between goroutines.*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)


// this wants a channel to be passed to it
// think of a channel like a 'wormhole'
func CalculateValue(values chan int, duration time.Duration) {
    value := rand.Intn(10)  // create a random value
    fmt.Println("Calculated Random Value: {}", value)
    time.Sleep( time.Second * duration) // wait 10 seconds!
    values <- value  // send our value back through the wormhole
    // after the data is through the wormhole
    fmt.Println("I run after the data is added to the channel")
}

func main() {
    fmt.Println("Go Channel Tutorial")

    // create an empty channel called "values" (UNBUFFERED)
    values := make(chan int)
    defer close(values)

    go CalculateValue(values, 3) // pass our channel to our function
    go CalculateValue(values, 6)
    go CalculateValue(values, 9)
    go CalculateValue(values, 12)

    // wait until a value is returned
    value := <-values // after 10 seconds a value will return

    time.Sleep( time.Second ) // wait a single second (after data is read from the channel)
    fmt.Println(value) // I run after the data is read from the channel
}
/*
student@bchd:~/static/alta3golang/beyondBasics/channels$ go run ./channellecture02.go
Go Channel Tutorial
Calculated Random Value: {} 1
Calculated Random Value: {} 7
Calculated Random Value: {} 7
Calculated Random Value: {} 9
I run after the data is added to the channel
1
student@bchd:~/static/alta3golang/beyondBasics/channels$ go run ./channellecture02.go
Go Channel Tutorial
Calculated Random Value: {} 1
Calculated Random Value: {} 7
Calculated Random Value: {} 7
Calculated Random Value: {} 9
I run after the data is added to the channel
7
*/
