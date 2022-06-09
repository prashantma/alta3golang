/* Alta3 Research | RZFeeser
   CHALLENGE 01 - testing an empty string */

package main

import (
    "testing"   // used for testing
    "regexp"    // regular expression library
)

func TestHello1(t *testing.T) {
    testName := "RZFeeser"
    desiredResult := regexp.MustCompile("Hello there "+testName)

    result, err := Hello("RZFeeser")
    
    if !desiredResult.MatchString(result) || err != nil {
        t.Fatalf("Wanted %v, but got, %v, or got %v", desiredResult, result, err)
    }
    
    result, err = Hello("")
    
    // we are testing the ability to produce an error when entering "ENTER" instead of a name
    // this should return an error for err
    if result != "" || err == nil {
        t.Fatalf("Test failed to produce correct failure for empty string. Got, %v", result)
    }
}
/*
=== RUN   TestHello1
--- PASS: TestHello1 (0.00s)
=== RUN   TestHello
--- PASS: TestHello (0.00s)
PASS
ok      github.com/prashantma/beyondBasics/testing      0.003s
*/
