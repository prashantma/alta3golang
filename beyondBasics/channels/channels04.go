/* Alta3 Research | RZFeeser
   Channels and goroutines - skeleton for HTTP lookups with 3 workers */

package main

import (
        "fmt"
        "sync"
        "time"
)

// this is our "worker" function
// the worker wants to harvest data from the channel (linkChan)
func worker(linkChan chan string, wg *sync.WaitGroup, workerId int) {
        // Decreasing internal counter for wait-group as soon as goroutine finishes
        defer wg.Done()      

        // loop across the incoming channel (it is treated as a slice)
        for url := range linkChan {
                time.Sleep(1 * time.Second)                     // this would be work we want to perform
                fmt.Printf("Done processing link #%s on worker #%d\n", url, workerId)   // this would be work we want to perform
        }

}


func main() {
        // imagine we have a slice of links we want to perform HTTP GET requests on
        yourLinksSlice := make([]string, 50)
        
        // write 1 to 50 in the slice we just created
        for i := 0; i < 50; i++ {
                yourLinksSlice[i] = fmt.Sprintf("%d", i+1)
        }

        // use make or the channel won't be useful
        lCh := make(chan string)
        
        
        wg := new(sync.WaitGroup) // returns a pointer to initialized memory

        // Adding routines to workgroup and running then
        for i := 0; i < 3; i++ {
                wg.Add(1)
                go worker(lCh, wg, i)
        }

        // Processing all links by spreading them to `free` goroutines
        for _, link := range yourLinksSlice {
                lCh <- link
        }

        // Closing channel (waiting in goroutines won't continue any more)
        close(lCh) // no more jobs will be sent here

        // Waiting for all goroutines to finish (otherwise they die as main routine dies)
        wg.Wait()
}
/*
student@bchd:~/static/alta3golang/beyondBasics/channels$ go run ./channels04.go
Done processing link #3
Done processing link #1
Done processing link #2
Done processing link #4
Done processing link #6
Done processing link #5
Done processing link #7
Done processing link #9
Done processing link #8
Done processing link #10
Done processing link #12
Done processing link #11
Done processing link #13
Done processing link #15
Done processing link #14
Done processing link #16
Done processing link #18
Done processing link #17
Done processing link #19
Done processing link #21
Done processing link #20
Done processing link #22
Done processing link #24
Done processing link #23
Done processing link #25
Done processing link #27
Done processing link #26
Done processing link #28
Done processing link #30
Done processing link #29
Done processing link #31
Done processing link #33
Done processing link #32
Done processing link #34
Done processing link #36
Done processing link #35
Done processing link #37
Done processing link #39
Done processing link #38
Done processing link #40
Done processing link #42
Done processing link #41
Done processing link #43
Done processing link #45
Done processing link #44
Done processing link #46
Done processing link #48
Done processing link #47
Done processing link #49
Done processing link #50
*/
/*student@bchd:~/static/alta3golang/beyondBasics/channels$ go run ./channels04.go
Done processing link #3 on worker #1
Done processing link #2 on worker #0
Done processing link #1 on worker #2
Done processing link #4 on worker #1
Done processing link #5 on worker #0
Done processing link #6 on worker #2
Done processing link #7 on worker #1
Done processing link #8 on worker #0
Done processing link #9 on worker #2
Done processing link #10 on worker #1
Done processing link #11 on worker #0
Done processing link #12 on worker #2
Done processing link #13 on worker #1
Done processing link #14 on worker #0
Done processing link #15 on worker #2
Done processing link #16 on worker #1
Done processing link #17 on worker #0
Done processing link #18 on worker #2
Done processing link #19 on worker #1
Done processing link #20 on worker #0
Done processing link #21 on worker #2
Done processing link #22 on worker #1
Done processing link #23 on worker #0
Done processing link #24 on worker #2
Done processing link #25 on worker #1
Done processing link #26 on worker #0
Done processing link #27 on worker #2
Done processing link #28 on worker #1
Done processing link #29 on worker #0
Done processing link #30 on worker #2
Done processing link #31 on worker #1
Done processing link #32 on worker #0
Done processing link #33 on worker #2
Done processing link #34 on worker #1
Done processing link #35 on worker #0
Done processing link #36 on worker #2
Done processing link #37 on worker #1
Done processing link #38 on worker #0
Done processing link #39 on worker #2
Done processing link #40 on worker #1
Done processing link #41 on worker #0
Done processing link #42 on worker #2
Done processing link #43 on worker #1
Done processing link #44 on worker #0
Done processing link #45 on worker #2
Done processing link #46 on worker #1
Done processing link #47 on worker #0
Done processing link #48 on worker #2
Done processing link #49 on worker #1
Done processing link #50 on worker #0
*/
