/* Alta3 Research | RZFeeser
   TDD - Feature Design example (refactored)
   
   PlayerServer - Triggered by a lookup to our server

   GetPlayerScore - Performing the work of returning a player score  */


package main

import (
        "fmt"
        "net/http"
        "strings"
)

// move our function we refactored into an interface
type PlayerStore interface {
        GetPlayerScore(name string) int
}

// For our PlayerServer to be ale to use a PlayerStore, it needs to be referenced to one
// PlayerServer is now a struct (previously a function)
type PlayerServer struct {
        store PlayerStore
}


// create the handler interface 
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        player := strings.TrimPrefix(r.URL.Path, "/players/")
        fmt.Fprint(w, p.store.GetPlayerScore(player))   // this gets the score rather than the local function we had previously
}

