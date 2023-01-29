package main

import (
	"context"
	"log"
)

type metricsService struct {
	next PriceFectcher
}

func NewMetricsService(next PriceFectcher) PriceFectcher {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticket string) (price float64, err error) {
	log.Println("pushing metrics to prometheus")
	return s.next.FetchPrice(ctx, ticket)
}
