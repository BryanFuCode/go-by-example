package main

import (
	"fmt"
	"os"

	"github.com/BryanFuCode/go-by-example/simpledict/v5/translator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	translator.QueryCaiyun(word)
	translator.QueryHuoshan(word)
}
