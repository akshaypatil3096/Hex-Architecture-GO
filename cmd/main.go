package main

import (
	"fmt"

	"github.com/akshaypatil3096/Hex-Architecture-GO/internal/adapters/core/arithmetic"
)

func main() {
	fmt.Println("HI")

	arithAdapter := arithmetic.NewAdapter()
	res, _ := arithAdapter.Addition(1, 2)

	// api.NewAdapter(arithAdapter)
	fmt.Println(res)
}
