/* Alta3 Research | RZFeeser
   TDD - Building tests to drive feature development (refactored)

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
        "fmt"
)

func TestGETPlayers2(t *testing.T) {
        t.Run("returns Pepper's score", func(t *testing.T) {
                request := newGetScoreRequest2("Pepper")
                response := httptest.NewRecorder()

                PlayerServer3(response, request)

                assertResponseBody2(t, response.Body.String(), "20") // we still have hardcoded values
        })

        t.Run("returns Floyd's score", func(t *testing.T) {
                request := newGetScoreRequest2("Floyd")
                response := httptest.NewRecorder()

                PlayerServer3(response, request)

                assertResponseBody2(t, response.Body.String(), "10") // we still have hardcoded values
        })
}

// new function added to help our testing function (added during refactoring)
// looks up players by name and returns their score
func newGetScoreRequest2(name string) *http.Request {
        req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
        return req
}

// new function added to help our testing function (added during refactoring)
// ensures the body contains the value we expect
func assertResponseBody2(t testing.TB, got, want string) {
        t.Helper()   // Helper marks the calling function as a test helper function
                     // When printing file and line information, that function will be skipped.
        if got != want {
                t.Errorf("response body is wrong, got %q want %q", got, want)
        }
}
// with original server.go (now renamed as server3.go)
// with original server_test.go (now renamed as server_test2.go)
// with original main.go (now renamed as main1.go)
/*
student@bchd:~/static/alta3golang/beyondBasics/tdd/v1$ go test -v
=== RUN   TestGETPlayers
=== RUN   TestGETPlayers/returns_Pepper's_score
=== RUN   TestGETPlayers/returns_Floyd's_score
--- PASS: TestGETPlayers (0.00s)
    --- PASS: TestGETPlayers/returns_Pepper's_score (0.00s)
    --- PASS: TestGETPlayers/returns_Floyd's_score (0.00s)
PASS
ok      github.com/prashantma/beyondBasics/tdd  0.006s
*/
