/* Alta3 Research | RZFeeser
   TDD - Feature Design Example
   
   The solution to our test should be fast, cut corners, and narrow in scope. Our goal is to create a solution that turns a RED test to a GREEN test.
   
   While code is testing RED, you can feel free to revert the HEAD if you get stuck / lost. You won't lose much.

   Testing GREEN? Time to refactor (code can be CLEANED UP). */

// still part of the main package, just located within a new file
package main
  
import (
        "fmt"
        "net/http"
)

// PlayerServer currently returns Hello, world given _any_ request.
func PlayerServer1(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "20") // always return 20 (this is what our test demands)
}

