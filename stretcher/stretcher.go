package stretcher


type Mangler func(channel <- chan []rune) <- chan []rune


func CreateManglerChain(in_chan <- chan []rune, manglers []Mangler) <- chan []rune{

    if len(manglers) == 0 {
        return in_chan
    } 

    out_chan := manglers[0](in_chan)

    if len(manglers) == 1 {
        return out_chan
    } else {
        return CreateManglerChain(out_chan, manglers[1:])
    }
}