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

func TestGETPlayers3(t *testing.T) {
    server := &PlayerServer{} // create a new instance of PlayerServer

    t.Run("returns Pepper's score", func(t *testing.T) {
        request := newGetScoreRequest3("Pepper")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)    // call the method ServeHTTP

        assertResponseBody3(t, response.Body.String(), "20")
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request := newGetScoreRequest3("Floyd")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertResponseBody3(t, response.Body.String(), "10")
    })
}

// new function added to help our testing function (added during refactoring)
// looks up players by name and returns their score
func newGetScoreRequest3(name string) *http.Request {
        req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
        return req
}

// new function added to help our testing function (added during refactoring)
// ensures the body contains the value we expect
func assertResponseBody3(t testing.TB, got, want string) {
        t.Helper()   // Helper marks the calling function as a test helper function
                     // When printing file and line information, that function will be skipped.
        if got != want {
                t.Errorf("response body is wrong, got %q want %q", got, want)
        }
}
// when "PlayStore" is not passed in test
// with original main.go (now renamed as main1.go)
/*
student@bchd:~/static/alta3golang/beyondBasics/tdd/v1$ go test -v
=== RUN   TestGETPlayers
=== RUN   TestGETPlayers/returns_Pepper's_score
--- FAIL: TestGETPlayers (0.00s)
    --- FAIL: TestGETPlayers/returns_Pepper's_score (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
        panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x58c42c]

goroutine 21 [running]:
testing.tRunner.func1.2({0x5abda0, 0x72cbd0})
        /usr/local/go/src/testing/testing.go:1389 +0x24e
testing.tRunner.func1()
        /usr/local/go/src/testing/testing.go:1392 +0x39f
panic({0x5abda0, 0x72cbd0})
        /usr/local/go/src/runtime/panic.go:838 +0x207
github.com/prashantma/beyondBasics/tdd.(*PlayerServer).ServeHTTP(0xc000080db0?, {0x61e048?, 0xc0000849c0}, 0x6e29b8?)
        /home/student/static/alta3golang/beyondBasics/tdd/v1/server.go:32 +0xac
github.com/prashantma/beyondBasics/tdd.TestGETPlayers.func1(0x0?)
        /home/student/static/alta3golang/beyondBasics/tdd/v1/server_test.go:30 +0xca
testing.tRunner(0xc000083ba0, 0xc000080dc0)
        /usr/local/go/src/testing/testing.go:1439 +0x102
created by testing.(*T).Run
        /usr/local/go/src/testing/testing.go:1486 +0x35f
exit status 2
FAIL    github.com/prashantma/beyondBasics/tdd  0.008s
*/
