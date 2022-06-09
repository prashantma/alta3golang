/* Alta3 Research | RZFeeser
   exec.Command().Output() - Capturing output from commands   */


package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {

    out, err := exec.Command("ping", "google.com", "-c", "4").Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
}

