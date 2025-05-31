package module01

import (
	"fmt"
)

func Reverse(str string) string {
	reversed := ""

	//rune is Go terminology for a single Unicode code point.
	runes := []rune(str)

	fmt.Printf("Runes of %s (unicode): %v\n", str, runes)

	for i := len(runes) - 1; i >= 0; i-- {
		reversed += string(runes[i])
	}

	return reversed
}
