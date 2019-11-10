package stretcher

import (
    "unicode"
)


func Cap(in_chan <- chan []rune) <- chan []rune {
    // Essentially a generator

    out_chan := make(chan []rune)

    go func() {
        defer close(out_chan)
        for s := range in_chan {
            for capped := range _cap(s) {
                out_chan <- capped
            }
        }
    }()

    return out_chan
}


func _cap(s []rune) <- chan []rune {

    out_chan := make(chan []rune)

    go func() {
        defer close(out_chan)
        if len(s) == 1 {
            if unicode.IsLetter(s[0]) {
                out_chan <- []rune{unicode.ToUpper(s[0])}
                out_chan <- []rune{unicode.ToLower(s[0])}
            } else {
                out_chan <- s
            }
        } else {

            mid_point := len(s) / 2

            for right_half := range _cap(s[mid_point:]) {
                for left_half := range _cap(s[:mid_point]) {
                    out_chan <- append(left_half, right_half...)
                }
            }
        }
    }()

    return out_chan
}