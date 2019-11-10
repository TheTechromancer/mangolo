package utils

import (
    "os"
    "bufio"
)

func Reader(filename string) <- chan []rune {
    // yields lines from filename

    out_chan := make(chan []rune)

    go func() {
        f, _ := os.Open(filename)
        defer f.Close()
        fileScanner := bufio.NewScanner(f)
        for fileScanner.Scan() {
            out_chan <- []rune(fileScanner.Text())
        }
        close(out_chan)
    }()

    return out_chan

}