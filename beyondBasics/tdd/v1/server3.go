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

// refactoring gives us the opportunity to study our code
func PlayerServer3(w http.ResponseWriter, r *http.Request) {
        player := strings.TrimPrefix(r.URL.Path, "/players/")

        fmt.Fprint(w, GetPlayerScore(player))  // this function will grab the score of the looked up player and return it
}

// new function presented by refactoring
// we don't have tests for anything more than Pepper and Floyd ATM
// so we will resis the temptation from doing anything more than necessary
// TDD implies you will build features as tests demand them, not before
func GetPlayerScore(name string) string {
        if name == "Pepper" {
                return "20"
        }

        if name == "Floyd" {
                return "10"
        }

        return ""
}

