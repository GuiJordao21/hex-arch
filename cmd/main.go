package main

import (
	"fmt"

	"github.com/GuiJordao21/hex-arch/internal/adapters/core/arithmetic"
)

func main() {
	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(3, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
