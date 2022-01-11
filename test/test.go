package main

import (
	"fmt"

	"go-reloaded/functions"
)

func main() {
	fmt.Println(functions.Hex("1E"))
	fmt.Println(functions.Bin("10"))
	fmt.Println(functions.Up("hello"))
	fmt.Println(functions.Low("HELLo"))
	fmt.Println(functions.Cap("hElLo"))
}