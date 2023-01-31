package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "github.com/etherwatt/gen"
	"github.com/etherwatt/internal/transport"
	"github.com/etherwatt/internal/usecases"
)

var ganacheURL = "http://localhost:7545"

// @title EtherWatt API
// @description EtherWatt server API

// @host      localhost:8080
// @BasePath /etherwatt/v1
func main() {

	// Connect to ethereum blockchain
	ctx := context.Background()

	client, err := ethclient.DialContext(ctx, ganacheURL)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()

	// Contract
	deployedContractAddress := "0xbcA30817507c3D27Cfd4353D007b9f09504A05f0"
	cAdd := common.HexToAddress(deployedContractAddress)
	c, err := contract.NewEtherwatt(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	uc := usecases.NewUsecases(client, c)

	t := transport.NewTrasport(uc)

	server := t.InitRoutes()

	server.Run()
}
