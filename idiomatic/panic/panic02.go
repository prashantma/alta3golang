/* Alta3 Research | RZFeeser
   Raising a panic() for an error we are not prepared to handle ourselves. */

package main

import "os"

func main() {

    _, err := os.Create("/carolDanvers/msmarvel")
    
    // if we have an unexpected error
    if err != nil {
        panic(err)  // fast fail and dump the error to the screen
    }
}
/*
panic: open /carolDanvers/msmarvel: no such file or directory

goroutine 1 [running]:
main.main()
        /home/student/static/alta3golang/idiomatic/panic/panic02.go:14 +0x51
exit status 2
*/
