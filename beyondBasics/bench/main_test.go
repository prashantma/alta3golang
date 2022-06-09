/* Alta3 Research | RZFeeser
   Mixing tests with our benchmark tests. */
   
package main

import (
    "fmt"
    "testing"
)

// additional test (non benchmark)
func TestCalculate1(t *testing.T) {
    fmt.Println("Test Calculate")
    expected := 4
    result := Calculate(2)
    if expected != result {
        t.Error("Failed")
    }
}

// the only benchmark test
func BenchmarkCalculate1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Calculate(2)
    }
}

// additional test (non benchmark)
func TestOther1(t *testing.T) {
    fmt.Println("Testing something else")
    fmt.Println("This shouldn't run with -run=calc")
}
/*
student@bchd:~/static/alta3golang/beyondBasics/bench$ go test -run=Calculate -bench=.
Test Calculate
goos: linux
goarch: amd64
pkg: github.com/prashantma/beyondBasics/bench
cpu: Intel(R) Xeon(R) CPU E5-2650 v3 @ 2.30GHz
BenchmarkCalculate1-2           1000000000               0.4703 ns/op
PASS
ok      github.com/prashantma/beyondBasics/bench        0.537s
*/
