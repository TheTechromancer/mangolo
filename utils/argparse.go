package utils

import (
    "os"
    "flag"
    "github.com/mattn/go-isatty"
)


type Options struct {
    Input   *string
    Capswap *bool
    OutPipe *bool
}


func ParseArgs() *Options {


    out_pipe := true
    /*
    if (fi.Mode() & os.ModeCharDevice) == 0 || (fi.Mode() & os.ModeNamedPipe) == 0 {
        out_pipe = true
    } else {
        out_pipe = false
    }*/
    if isatty.IsTerminal(os.Stdout.Fd()) {
        out_pipe = false
    }

    options := &Options {
        Input: flag.String("input", "__stdin__", "Input wordlist"),
        Capswap: flag.Bool("capswap", false, "Output all possible capital permutations"),
        OutPipe: &out_pipe,
    }
    
    flag.Parse()

    // print help if there are no arguments
    if len(os.Args) < 2 && *options.Input != "__stdin__" {
        flag.PrintDefaults()
        os.Exit(2)
    }

    return options
}