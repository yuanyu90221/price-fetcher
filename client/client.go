package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yuanyu90221/price-fetcher/types"
)

type Client struct {
	endPoint string
}

func New(endPoint string) *Client {
	return &Client{
		endPoint: endPoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticket string) (*types.PriceResponse, error) {
	endPoint := fmt.Sprintf("%s?ticket=%s", c.endPoint, ticket)

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code: %s", httpErr["error"])
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}

	return priceResp, nil
}
