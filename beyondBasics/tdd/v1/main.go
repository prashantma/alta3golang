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

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
        return 42  // hard code the return of the value 42
}

func main() {

        // create a new instance of PlayerServer
        server := &PlayerServer{}

        // use InMemoryPlayerStore to change empty reply in response
        // server := &PlayerServer{&InMemoryPlayerStore{}}

        log.Fatal(http.ListenAndServe(":5000", server))   // run on port 5000, invoke the server instance
}

