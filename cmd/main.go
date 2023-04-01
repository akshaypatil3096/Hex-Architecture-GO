package main

import (
	"fmt"

	"github.com/akshaypatil3096/Hex-Architecture-GO/internal/adapters/framework/right/db"
	"github.com/akshaypatil3096/Hex-Architecture-GO/internal/ports"
)

func main() {
	prepare()
}

func prepare() {
	fmt.Println("HI")
	var (
		err error

		// ports
		dbaseAdapter ports.DbPort
		core         ports.ArithmeticPort
		appAdapter   ports.APIPort
		gRPCAdapter  ports.GRPCPort
	)

	db.NewAdapter("")
}
