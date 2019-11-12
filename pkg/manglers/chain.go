package manglers

import (
    "mangolo/utils"
)

type Mangler func(in_chan <- chan []rune) <- chan []rune


func MakeManglerChain(in_chan <- chan []rune, options *utils.Options) <- chan []rune {
    // Given an options object, constructs chain of manglers

    mangler_chain := []Mangler{}
    if *options.Capswap == true {
        mangler_chain = append(mangler_chain, Capswap)
    }
    return make_mangler_channel(in_chan, mangler_chain)
}


func make_mangler_channel(in_chan <- chan []rune, mangler_chain []Mangler) <- chan []rune{
    // Given an array of manglers, chains them together and returns an output channel 

    if len(mangler_chain) == 0 {
        return in_chan
    } 

    out_chan := mangler_chain[0](in_chan)

    if len(mangler_chain) == 1 {
        return out_chan
    } else {
        return make_mangler_channel(out_chan, mangler_chain[1:])
    }
}