package notion

import (
	"context"
	"fmt"
)

// UsersListResponse represents a list of users
type UsersListResponse struct {
	ListResponse
	Results []User `json:"results"`
}

// GetUser retrieves a user by ID
func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	resp, err := c.makeRequest(ctx, "GET", "/users/"+userID, nil)
	if err != nil {
		return nil, err
	}

	var user User
	if err := parseResponse(resp, &user); err != nil {
		return nil, fmt.Errorf("failed to parse user response: %w", err)
	}

	return &user, nil
}

// ListUsers retrieves a list of all users
func (c *Client) ListUsers(ctx context.Context, startCursor string, pageSize int) (*UsersListResponse, error) {
	path := "/users"

	if startCursor != "" || pageSize > 0 {
		path += "?"
		if startCursor != "" {
			path += "start_cursor=" + startCursor
		}
		if pageSize > 0 {
			if startCursor != "" {
				path += "&"
			}
			path += fmt.Sprintf("page_size=%d", pageSize)
		}
	}

	resp, err := c.makeRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var users UsersListResponse
	if err := parseResponse(resp, &users); err != nil {
		return nil, fmt.Errorf("failed to parse users response: %w", err)
	}

	return &users, nil
}

// GetMe retrieves the current bot user
func (c *Client) GetMe(ctx context.Context) (*User, error) {
	resp, err := c.makeRequest(ctx, "GET", "/users/me", nil)
	if err != nil {
		return nil, err
	}

	var user User
	if err := parseResponse(resp, &user); err != nil {
		return nil, fmt.Errorf("failed to parse user response: %w", err)
	}

	return &user, nil
}
