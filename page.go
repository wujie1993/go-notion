package notion

import (
	"context"
	"fmt"
)

// Page represents a Notion page
type Page struct {
	Object         string                  `json:"object"`
	ID             string                  `json:"id"`
	CreatedTime    string                  `json:"created_time"`
	CreatedBy      *User                   `json:"created_by"`
	LastEditedTime string                  `json:"last_edited_time"`
	LastEditedBy   *User                   `json:"last_edited_by"`
	Archived       bool                    `json:"archived"`
	Icon           *Icon                   `json:"icon,omitempty"`
	Cover          *Cover                  `json:"cover,omitempty"`
	Properties     map[string]PageProperty `json:"properties"`
	Parent         *Parent                 `json:"parent"`
	URL            string                  `json:"url"`
	PublicURL      string                  `json:"public_url,omitempty"`
}

// PageProperty represents a page property value
type PageProperty struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	Title          []RichText     `json:"title,omitempty"`
	RichText       []RichText     `json:"rich_text,omitempty"`
	Number         *float64       `json:"number,omitempty"`
	Select         *SelectOption  `json:"select,omitempty"`
	MultiSelect    []SelectOption `json:"multi_select,omitempty"`
	Date           *Date          `json:"date,omitempty"`
	Formula        *Formula       `json:"formula,omitempty"`
	Relation       []Relation     `json:"relation,omitempty"`
	Rollup         *Rollup        `json:"rollup,omitempty"`
	People         []User         `json:"people,omitempty"`
	Files          []File         `json:"files,omitempty"`
	Checkbox       bool           `json:"checkbox,omitempty"`
	URL            string         `json:"url,omitempty"`
	Email          string         `json:"email,omitempty"`
	PhoneNumber    string         `json:"phone_number,omitempty"`
	CreatedTime    string         `json:"created_time,omitempty"`
	CreatedBy      *User          `json:"created_by,omitempty"`
	LastEditedTime string         `json:"last_edited_time,omitempty"`
	LastEditedBy   *User          `json:"last_edited_by,omitempty"`
	Status         *StatusOption  `json:"status,omitempty"`
}

// Formula represents a formula result
type Formula struct {
	Type    string   `json:"type"`
	String  string   `json:"string,omitempty"`
	Number  *float64 `json:"number,omitempty"`
	Boolean *bool    `json:"boolean,omitempty"`
	Date    *Date    `json:"date,omitempty"`
}

// Relation represents a relation to another page
type Relation struct {
	ID string `json:"id"`
}

// Rollup represents a rollup result
type Rollup struct {
	Type       string        `json:"type"`
	Number     *float64      `json:"number,omitempty"`
	Date       *Date         `json:"date,omitempty"`
	Array      []RollupValue `json:"array,omitempty"`
	Incomplete bool          `json:"incomplete,omitempty"`
}

// RollupValue represents a value in a rollup array
type RollupValue struct {
	Type           string         `json:"type"`
	Title          []RichText     `json:"title,omitempty"`
	RichText       []RichText     `json:"rich_text,omitempty"`
	Number         *float64       `json:"number,omitempty"`
	Select         *SelectOption  `json:"select,omitempty"`
	MultiSelect    []SelectOption `json:"multi_select,omitempty"`
	Date           *Date          `json:"date,omitempty"`
	Formula        *Formula       `json:"formula,omitempty"`
	Relation       []Relation     `json:"relation,omitempty"`
	Rollup         *Rollup        `json:"rollup,omitempty"`
	People         []User         `json:"people,omitempty"`
	Files          []File         `json:"files,omitempty"`
	Checkbox       bool           `json:"checkbox,omitempty"`
	URL            string         `json:"url,omitempty"`
	Email          string         `json:"email,omitempty"`
	PhoneNumber    string         `json:"phone_number,omitempty"`
	CreatedTime    string         `json:"created_time,omitempty"`
	CreatedBy      *User          `json:"created_by,omitempty"`
	LastEditedTime string         `json:"last_edited_time,omitempty"`
	LastEditedBy   *User          `json:"last_edited_by,omitempty"`
}

// PagesListResponse represents a list of pages
type PagesListResponse struct {
	ListResponse
	Results []Page `json:"results"`
}

// CreatePageRequest represents a request to create a page
type CreatePageRequest struct {
	Parent     *Parent                 `json:"parent"`
	Properties map[string]PageProperty `json:"properties,omitempty"`
	Children   []Block                 `json:"children,omitempty"`
	Icon       *Icon                   `json:"icon,omitempty"`
	Cover      *Cover                  `json:"cover,omitempty"`
}

// UpdatePageRequest represents a request to update a page
type UpdatePageRequest struct {
	Properties map[string]PageProperty `json:"properties,omitempty"`
	Archived   *bool                   `json:"archived,omitempty"`
	Icon       *Icon                   `json:"icon,omitempty"`
	Cover      *Cover                  `json:"cover,omitempty"`
}

// GetPage retrieves a page by ID
func (c *Client) GetPage(ctx context.Context, pageID string) (*Page, error) {
	resp, err := c.makeRequest(ctx, "GET", "/pages/"+pageID, nil)
	if err != nil {
		return nil, err
	}

	var page Page
	if err := parseResponse(resp, &page); err != nil {
		return nil, fmt.Errorf("failed to parse page response: %w", err)
	}

	return &page, nil
}

// CreatePage creates a new page
func (c *Client) CreatePage(ctx context.Context, req *CreatePageRequest) (*Page, error) {
	resp, err := c.makeRequest(ctx, "POST", "/pages", req)
	if err != nil {
		return nil, err
	}

	var page Page
	if err := parseResponse(resp, &page); err != nil {
		return nil, fmt.Errorf("failed to parse page response: %w", err)
	}

	return &page, nil
}

// UpdatePage updates an existing page
func (c *Client) UpdatePage(ctx context.Context, pageID string, req *UpdatePageRequest) (*Page, error) {
	resp, err := c.makeRequest(ctx, "PATCH", "/pages/"+pageID, req)
	if err != nil {
		return nil, err
	}

	var page Page
	if err := parseResponse(resp, &page); err != nil {
		return nil, fmt.Errorf("failed to parse page response: %w", err)
	}

	return &page, nil
}
