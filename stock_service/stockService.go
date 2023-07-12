package stock_service

import (
	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedStockServiceServer
}

func (s *Server) GetStockDetails(ctx context.Context, req *StockRequest) (*StockResponse, error) {
	// response := StockResponse{Price: 10.0}
	response := StockResponse{Price: float32(CalculateIndicator(req))}
	return &response, nil
}
