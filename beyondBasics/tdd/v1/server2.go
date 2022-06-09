package  main

import (
        "fmt"
        "net/http"
        "strings"    // add the strings package
)

func PlayerServer2(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")     //  look at the incoming request and remove the /players/ prefix

    // test for looked up player is "Pepper"
    if player == "Pepper" {
        fmt.Fprint(w, "20")
        return
    }

    // test for looked up player is "Floyd"
    if player == "Floyd" {
        fmt.Fprint(w, "10")
        return
    }
}

