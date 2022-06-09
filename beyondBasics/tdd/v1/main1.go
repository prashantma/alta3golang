/* Alta3 Research | RZFeeser
   Intro to TDD operations

   You have been asking to create a web server where users can track how many games
   players have won.
   
   GET /players/{name} should return a number indicating the total number of wins
   POST /players/{name} should record a win for that name, incrementing for every subsequent POST
   
   Following a TDD approach means we need to make small iterative improvements until we have a solution.
   
   Remember...
   - Keep the problem narrow
   - Don't fall down rabbit holes
   - If you get stuck / lost, reverting to a previous state will not lose much work */

package main

import (
        "log"
        "net/http"
)

func main1() {

        // this is a handler, we want to run against this via a goroutine
        handler := http.HandlerFunc(PlayerServer3)

        // starts a webserver on a port, creating a goroutine for every request and running it against a handler
        log.Fatal(http.ListenAndServe(":5000", handler))   // run on port 5000, against the handler (this returns an error on failing)
}

