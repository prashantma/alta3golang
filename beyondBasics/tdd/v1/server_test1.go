/* Alta3 Research | RZFeeser
   TDD - Building tests to drive feature development.
   This test should be written BEFORE the feature is implimented.
   
   Functions need testing. Before writing the function (PlayerServer), we'll define 
   a test.

   Here we expect to get a plyers score (Pepper), which we expect to be 20.
   
   A narrow test such as this is "OK". */


package main

import (
        "net/http"
        "net/http/httptest"
        "testing"
)


// test the function PlayerServer
func TestGETPlayers1(t *testing.T) {
        // 
        t.Run("returns Pepper's score", func(t *testing.T) {
                
                // create a request           method,         path,              body
                request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
                response := httptest.NewRecorder()   // from the testing package, this "spy" lets us inspect what is in a response

                PlayerServer1(response, request)      // this is the function we are trying to test

                got := response.Body.String()        // what did we get back?
                want := "20"                         // this is what we expect to get back

                if got != want {
                        t.Errorf("got %q, want %q", got, want)
                }
        })
        t.Run("returns Floyd's score", func(t *testing.T) {
                request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
                response := httptest.NewRecorder()

                PlayerServer1(response, request)

                got := response.Body.String()
                want := "10"

                if got != want {
                    t.Errorf("got %q, want %q", got, want)
                }
        })
}
// with original server.go (now renamed as server1.go)
// with original server_test.go (now renamed as server_test1.go)
// with original main.go (now renamed as main1.go)
/*
student@bchd:~/static/alta3golang/beyondBasics/tdd/v1$ go test -v
=== RUN   TestGETPlayers
=== RUN   TestGETPlayers/returns_Pepper's_score
--- PASS: TestGETPlayers (0.00s)
    --- PASS: TestGETPlayers/returns_Pepper's_score (0.00s)
PASS
ok      github.com/prashantma/beyondBasics/tdd  0.008s
*/
/*
student@bchd:~/static/alta3golang/beyondBasics/tdd/v1$ go test -v
=== RUN   TestGETPlayers
=== RUN   TestGETPlayers/returns_Pepper's_score
=== RUN   TestGETPlayers/returns_Floyd's_score
    server_test.go:50: got "20", want "10"
--- FAIL: TestGETPlayers (0.00s)
    --- PASS: TestGETPlayers/returns_Pepper's_score (0.00s)
    --- FAIL: TestGETPlayers/returns_Floyd's_score (0.00s)
FAIL
exit status 1
FAIL    github.com/prashantma/beyondBasics/tdd  0.007s
*/

// with new server.go (now renamed as server2.go)
// with original server_test.go (now renamed as server_test1.go)
// with original main.go (now remamed as main1.go)
/*
student@bchd:~/static/alta3golang/beyondBasics/tdd/v1$ go test -v
=== RUN   TestGETPlayers
=== RUN   TestGETPlayers/returns_Pepper's_score
=== RUN   TestGETPlayers/returns_Floyd's_score
--- PASS: TestGETPlayers (0.00s)
    --- PASS: TestGETPlayers/returns_Pepper's_score (0.00s)
    --- PASS: TestGETPlayers/returns_Floyd's_score (0.00s)
PASS
ok      github.com/prashantma/beyondBasics/tdd  0.005s
*/

