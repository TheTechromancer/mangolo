package utils

import (
    "os"
    "bufio"
)

func Reader(filename string) <- chan []rune {
    // yields lines from STDIN or file

    out_chan := make(chan []rune)

    if filename == "__stdin__" {
        go func () {
            defer close(out_chan)
            stdinScanner := bufio.NewScanner(os.Stdin)
            for stdinScanner.Scan() {
                out_chan <- []rune(stdinScanner.Text())
            }
        }()

    } else {

        go func() {
            f, _ := os.Open(filename)
            defer close(out_chan)
            defer f.Close()
            fileScanner := bufio.NewScanner(f)
            for fileScanner.Scan() {
                out_chan <- []rune(fileScanner.Text())
            }
        }()
    }

    return out_chan

}