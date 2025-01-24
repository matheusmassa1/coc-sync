package cocapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

func NewClient(client *http.Client, apiKey string) *Client {
	return &Client{
		client: client,
		apiKey: apiKey,
		baseURL: "https://api.clashofclans.com/v1",
	}
}

func (c *Client) DoGet(ctx context.Context, path string, response interface{}) error {
	fullURL := fmt.Sprintf("%s%s", c.baseURL, path)
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if ctx.Err() != nil {
		return fmt.Errorf("request context cancelled: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, body)
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}
