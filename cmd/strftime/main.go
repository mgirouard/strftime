package main

import (
	"fmt"
	"os"
	"time"

	sft "github.com/mgirouard/strftime"
)

func main() {
	now := time.Now()
	if len(os.Args) == 2 {
		fmt.Println(sft.Fmt(os.Args[1], now))
	} else {
		fns := []string{
			"a", "A", "b", "B", "c", "C", "d", "D", "e", "E", "F", "G", "g",
			"h", "H", "i", "j", "k", "l", "m", "M", "n", "O", "p", "P", "r",
			"R", "s", "S", "t", "T", "u", "U", "V", "w", "W", "x", "X", "y",
			"Y", "z", "Z", "+", "%",
		}
		for _, f := range fns {
			fmt.Println(sft.Fmt(`%%`+f+` = '%`+f+`'`, now))
		}
	}
}
