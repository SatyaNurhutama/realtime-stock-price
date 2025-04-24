package service

import (
	"context"
	"time"

	"github.com/satyanurhutama/realtime-stock-price/pkg/api"
	"github.com/satyanurhutama/realtime-stock-price/pkg/repository"
)

type StockPriceService struct {
	api.UnimplementedStockPriceServiceServer
	Repository repository.AlphaVantageRepository
}

func (s *StockPriceService) GetStockPrice(ctx context.Context, req *api.StockPriceRequest) (*api.StockPriceResponse, error) {
	price, time, err := s.Repository.GetStockPrice(req.Symbol)
	if err != nil {
		return nil, err
	}

	return &api.StockPriceResponse{
		Symbol:    req.Symbol,
		Price:     price,
		Timestamp: time,
	}, nil
}

func (s *StockPriceService) GetStockPriceStream(req *api.StockPriceRequest, stream api.StockPriceService_GetStockPriceStreamServer) error {
	for {
		price, timestamp, err := s.Repository.GetStockPrice(req.Symbol)
		if err != nil {
			return err
		}

		response := &api.StockPriceResponse{
			Symbol:    req.Symbol,
			Price:     price,
			Timestamp: timestamp,
		}

		if err := stream.Send(response); err != nil {
			return err
		}

		time.Sleep(60 * time.Second)
	}
}
