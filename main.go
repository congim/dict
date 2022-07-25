package main

import (
	"fmt"

	"github.com/congim/dict/dict"
)

func main() {
	dic := dict.New()
	dic.AddWorld("DDDD", "fuck")
	fmt.Println(dic.Replace("ddddddddd Fuck", '*'))
}
