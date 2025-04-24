package main

import (
	"fmt"
	"log"
	"net"

	"github.com/satyanurhutama/realtime-stock-price/pkg/api"
	"github.com/satyanurhutama/realtime-stock-price/pkg/config"
	"github.com/satyanurhutama/realtime-stock-price/pkg/repository"
	"github.com/satyanurhutama/realtime-stock-price/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadEnv()
	repo := repository.AlphaVantageRepository{APIKey: config.GetEnv("ALPHA_VANTAGE_API_KEY")}
	service := service.StockPriceService{Repository: repo}

	grpcServer := grpc.NewServer()

	api.RegisterStockPriceServiceServer(grpcServer, &service)

	reflection.Register(grpcServer)

	// Start listening
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
