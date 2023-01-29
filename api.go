package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/yuanyu90221/price-fetcher/types"
)

type JSONAPIServer struct {
	listendAddr string
	svc         PriceFectcher
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func NewJSONAPIServer(listenAddr string, svc PriceFectcher) *JSONAPIServer {
	return &JSONAPIServer{
		listendAddr: listenAddr,
		svc:         svc,
	}
}
func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listendAddr, nil)
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(100000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticket := r.URL.Query().Get("ticket")
	price, err := s.svc.FetchPrice(ctx, ticket)
	if err != nil {
		return err
	}
	priceResp := types.PriceResponse{
		Price:  price,
		Ticket: ticket,
	}

	return writeJSON(w, http.StatusOK, &priceResp)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
