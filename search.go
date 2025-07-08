package notion

import (
	"context"
	"fmt"
)

// SearchRequest represents a search request
type SearchRequest struct {
	Query       string        `json:"query,omitempty"`
	Sort        *SearchSort   `json:"sort,omitempty"`
	Filter      *SearchFilter `json:"filter,omitempty"`
	StartCursor string        `json:"start_cursor,omitempty"`
	PageSize    int           `json:"page_size,omitempty"`
}

// SearchSort represents search sorting options
type SearchSort struct {
	Direction string `json:"direction"`
	Timestamp string `json:"timestamp"`
}

// SearchFilter represents search filtering options
type SearchFilter struct {
	Value    string `json:"value"`
	Property string `json:"property"`
}

// SearchResponse represents a search response
type SearchResponse struct {
	ListResponse
	Results []SearchResult `json:"results"`
}

// SearchResult represents a search result item
type SearchResult struct {
	Object   string    `json:"object"`
	Page     *Page     `json:"page,omitempty"`
	Database *Database `json:"database,omitempty"`
}

// Search performs a search across pages and databases
func (c *Client) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	resp, err := c.makeRequest(ctx, "POST", "/search", req)
	if err != nil {
		return nil, err
	}

	var searchResp SearchResponse
	if err := parseResponse(resp, &searchResp); err != nil {
		return nil, fmt.Errorf("failed to parse search response: %w", err)
	}

	return &searchResp, nil
}
