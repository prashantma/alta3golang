package main

import (
    "testing"
)

func BenchmarkCalculate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Calculate(2)
    }
}
/*
student@bchd:~/static/alta3golang/beyondBasics/bench$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/prashantma/beyondBasics/bench
cpu: Intel(R) Xeon(R) CPU E5-2650 v3 @ 2.30GHz
BenchmarkCalculate-2    1000000000               0.4599 ns/op
PASS
ok      github.com/prashantma/beyondBasics/bench        0.530s
*/
