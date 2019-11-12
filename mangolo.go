package main

import (

    "os"
    "fmt"

    "mangolo/utils"
    "mangolo/pkg/manglers"
)

func main() {

    options := utils.ParseArgs()
    reader := utils.Reader(*options.Input)

    mangler_chain := manglers.MakeManglerChain(reader, options)
    
    if *options.Input == "__stdin__" {
        os.Stderr.WriteString("[+] Reading from STDIN\n")
    }

    // print every mangled word
    iteration := 0
    for s := range mangler_chain {
        fmt.Println(string(s))
        if (*options.OutPipe == true) && (iteration % 100 == 0) {
            fmt.Fprintf(os.Stderr, "\r[+] %d words written", iteration)
        }
        iteration += 1
    }
    if (*options.OutPipe == true) {
        fmt.Fprintf(os.Stderr, "\r[+] %d words written\n", iteration)
    }

}