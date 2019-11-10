package main

import (

    "os"
    "fmt"
    "flag"

    "./utils"
    "./stretcher"
)

func main() {

    wordlist := flag.String("wordlist", "", "Wordlist to parse.")
    mangle := flag.Bool("mangle", false, "Mangle the wordlist")
    flag.Parse()

    // print help if there are no arguments
    if len(os.Args) < 2 {
        flag.PrintDefaults()
        os.Exit(2)
    }

    manglers := []stretcher.Mangler{}
    if *mangle == true {
        manglers = append(manglers, stretcher.Cap)
    }

    reader := utils.Reader(*wordlist)

    for s := range stretcher.CreateManglerChain(reader, manglers) {
        fmt.Println(string(s))
    }

}