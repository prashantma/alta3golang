/* Alta3 Research | RZFeeser
   String Formatter - Lecture introduction */
   
package main

import (
    "fmt"		
)

var(
    a = 654                 // integer
    b = false               // boolean
    c = 2.651               // float
    d = 4 + 1i              // scientific notation
    e = "America"           // string
    f = 15.2 * 4525.321     // complex
)

func main(){	

    fmt.Printf("d for Integer: %d\n", a)
    fmt.Printf("6d for Integer: %6d\n", a)
    
    fmt.Printf("t for Boolean: %t\n", b)
    fmt.Printf("g for Float: %g\n", c)
    fmt.Printf("e for Scientific Notation: %e\n", d)
    fmt.Printf("E for Scientific Notation: %E\n", d)
    fmt.Printf("s for String: %s\n", e)
    fmt.Printf("G for Complex: %G\n", f)
    
    fmt.Printf("15s String: %15s\n", e)
    fmt.Printf("-10s String: %-10s\n",e)

    // create a formatted string, but don't print it out, just return the finished format
    t:= fmt.Sprintf("Print from right: %[3]d %[2]d %[1]d\n", 11, 22, 33)
    fmt.Println(t)	
}

