package main

import (
	"context"
	"fmt"
	"time"
)

type PriceFectcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct {
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticket string) (float64, error) {
	return MockPriceFetcher(ctx, ticket)
}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 1_000.0,
	"XRP": 0.5,
}

func MockPriceFetcher(ctx context.Context, ticket string) (float64, error) {
	time.Sleep(100 * time.Millisecond)
	price, ok := priceMocks[ticket]
	if !ok {
		return price, fmt.Errorf("the given ticket %s is not supported", ticket)
	}

	return price, nil
}
