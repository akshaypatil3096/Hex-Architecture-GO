package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akshaypatil3096/Hex-Architecture-GO/internal/adapters/app/api"
	"github.com/akshaypatil3096/Hex-Architecture-GO/internal/adapters/core/arithmetic"
	grpc "github.com/akshaypatil3096/Hex-Architecture-GO/internal/adapters/framework/left/gRPC"
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

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}

	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = grpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
