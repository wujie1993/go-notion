package notion

import (
	"context"
	"fmt"
)

// Database represents a Notion database
type Database struct {
	Object         string                      `json:"object"`
	ID             string                      `json:"id"`
	CreatedTime    string                      `json:"created_time"`
	CreatedBy      *User                       `json:"created_by"`
	LastEditedTime string                      `json:"last_edited_time"`
	LastEditedBy   *User                       `json:"last_edited_by"`
	Title          []RichText                  `json:"title"`
	Description    []RichText                  `json:"description,omitempty"`
	Icon           *Icon                       `json:"icon,omitempty"`
	Cover          *Cover                      `json:"cover,omitempty"`
	Properties     map[string]DatabaseProperty `json:"properties"`
	Parent         *Parent                     `json:"parent"`
	URL            string                      `json:"url"`
	Archived       bool                        `json:"archived"`
	IsInline       bool                        `json:"is_inline"`
	PublicURL      string                      `json:"public_url,omitempty"`
}

// DatabaseProperty represents a database property
type DatabaseProperty struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Title          map[string]interface{} `json:"title,omitempty"`
	RichText       map[string]interface{} `json:"rich_text,omitempty"`
	Number         *NumberProperty        `json:"number,omitempty"`
	Select         *SelectProperty        `json:"select,omitempty"`
	MultiSelect    *MultiSelectProperty   `json:"multi_select,omitempty"`
	Date           map[string]interface{} `json:"date,omitempty"`
	People         map[string]interface{} `json:"people,omitempty"`
	Files          map[string]interface{} `json:"files,omitempty"`
	Checkbox       map[string]interface{} `json:"checkbox,omitempty"`
	URL            map[string]interface{} `json:"url,omitempty"`
	Email          map[string]interface{} `json:"email,omitempty"`
	PhoneNumber    map[string]interface{} `json:"phone_number,omitempty"`
	Formula        *FormulaProperty       `json:"formula,omitempty"`
	Relation       *RelationProperty      `json:"relation,omitempty"`
	Rollup         *RollupProperty        `json:"rollup,omitempty"`
	CreatedTime    map[string]interface{} `json:"created_time,omitempty"`
	CreatedBy      map[string]interface{} `json:"created_by,omitempty"`
	LastEditedTime map[string]interface{} `json:"last_edited_time,omitempty"`
	LastEditedBy   map[string]interface{} `json:"last_edited_by,omitempty"`
	Status         *StatusProperty        `json:"status,omitempty"`
}

// NumberProperty represents a number property configuration
type NumberProperty struct {
	Format string `json:"format"`
}

// SelectProperty represents a select property configuration
type SelectProperty struct {
	Options []SelectOption `json:"options"`
}

// MultiSelectProperty represents a multi-select property configuration
type MultiSelectProperty struct {
	Options []SelectOption `json:"options"`
}

// SelectOption represents a select option
type SelectOption struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

// FormulaProperty represents a formula property configuration
type FormulaProperty struct {
	Expression string `json:"expression"`
}

// RelationProperty represents a relation property configuration
type RelationProperty struct {
	DatabaseID     string                 `json:"database_id"`
	Type           string                 `json:"type,omitempty"`
	SingleProperty map[string]interface{} `json:"single_property,omitempty"`
	DualProperty   *DualProperty          `json:"dual_property,omitempty"`
}

// DualProperty represents a dual relation property
type DualProperty struct {
	SyncedPropertyName string `json:"synced_property_name"`
	SyncedPropertyID   string `json:"synced_property_id"`
}

// RollupProperty represents a rollup property configuration
type RollupProperty struct {
	RelationPropertyName string `json:"relation_property_name"`
	RelationPropertyID   string `json:"relation_property_id"`
	RollupPropertyName   string `json:"rollup_property_name"`
	RollupPropertyID     string `json:"rollup_property_id"`
	Function             string `json:"function"`
}

// StatusProperty represents a status property configuration
type StatusProperty struct {
	Options []StatusOption `json:"options"`
	Groups  []StatusGroup  `json:"groups"`
}

// StatusOption represents a status option
type StatusOption struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

// StatusGroup represents a status group
type StatusGroup struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name"`
	Color     string   `json:"color,omitempty"`
	OptionIDs []string `json:"option_ids"`
}

// DatabasesListResponse represents a list of databases
type DatabasesListResponse struct {
	ListResponse
	Results []Database `json:"results"`
}

// CreateDatabaseRequest represents a request to create a database
type CreateDatabaseRequest struct {
	Parent      *Parent                     `json:"parent"`
	Title       []RichText                  `json:"title,omitempty"`
	Description []RichText                  `json:"description,omitempty"`
	Icon        *Icon                       `json:"icon,omitempty"`
	Cover       *Cover                      `json:"cover,omitempty"`
	Properties  map[string]DatabaseProperty `json:"properties"`
	IsInline    bool                        `json:"is_inline,omitempty"`
}

// UpdateDatabaseRequest represents a request to update a database
type UpdateDatabaseRequest struct {
	Title       []RichText                  `json:"title,omitempty"`
	Description []RichText                  `json:"description,omitempty"`
	Icon        *Icon                       `json:"icon,omitempty"`
	Cover       *Cover                      `json:"cover,omitempty"`
	Properties  map[string]DatabaseProperty `json:"properties,omitempty"`
	Archived    *bool                       `json:"archived,omitempty"`
}

// QueryDatabaseRequest represents a request to query a database
type QueryDatabaseRequest struct {
	Filter      *Filter `json:"filter,omitempty"`
	Sorts       []Sort  `json:"sorts,omitempty"`
	StartCursor string  `json:"start_cursor,omitempty"`
	PageSize    int     `json:"page_size,omitempty"`
}

// Filter represents a database filter
type Filter struct {
	Property    string             `json:"property,omitempty"`
	Type        string             `json:"type,omitempty"`
	RichText    *TextFilter        `json:"rich_text,omitempty"`
	Number      *NumberFilter      `json:"number,omitempty"`
	Checkbox    *CheckboxFilter    `json:"checkbox,omitempty"`
	Select      *SelectFilter      `json:"select,omitempty"`
	MultiSelect *MultiSelectFilter `json:"multi_select,omitempty"`
	Date        *DateFilter        `json:"date,omitempty"`
	People      *PeopleFilter      `json:"people,omitempty"`
	Files       *FilesFilter       `json:"files,omitempty"`
	Relation    *RelationFilter    `json:"relation,omitempty"`
	Formula     *FormulaFilter     `json:"formula,omitempty"`
	Or          []Filter           `json:"or,omitempty"`
	And         []Filter           `json:"and,omitempty"`
}

// TextFilter represents a text filter
type TextFilter struct {
	Equals         string `json:"equals,omitempty"`
	DoesNotEqual   string `json:"does_not_equal,omitempty"`
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	StartsWith     string `json:"starts_with,omitempty"`
	EndsWith       string `json:"ends_with,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// NumberFilter represents a number filter
type NumberFilter struct {
	Equals               *float64 `json:"equals,omitempty"`
	DoesNotEqual         *float64 `json:"does_not_equal,omitempty"`
	GreaterThan          *float64 `json:"greater_than,omitempty"`
	LessThan             *float64 `json:"less_than,omitempty"`
	GreaterThanOrEqualTo *float64 `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqualTo    *float64 `json:"less_than_or_equal_to,omitempty"`
	IsEmpty              bool     `json:"is_empty,omitempty"`
	IsNotEmpty           bool     `json:"is_not_empty,omitempty"`
}

// CheckboxFilter represents a checkbox filter
type CheckboxFilter struct {
	Equals       *bool `json:"equals,omitempty"`
	DoesNotEqual *bool `json:"does_not_equal,omitempty"`
}

// SelectFilter represents a select filter
type SelectFilter struct {
	Equals       string `json:"equals,omitempty"`
	DoesNotEqual string `json:"does_not_equal,omitempty"`
	IsEmpty      bool   `json:"is_empty,omitempty"`
	IsNotEmpty   bool   `json:"is_not_empty,omitempty"`
}

// MultiSelectFilter represents a multi-select filter
type MultiSelectFilter struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// DateFilter represents a date filter
type DateFilter struct {
	Equals     string `json:"equals,omitempty"`
	Before     string `json:"before,omitempty"`
	After      string `json:"after,omitempty"`
	OnOrBefore string `json:"on_or_before,omitempty"`
	OnOrAfter  string `json:"on_or_after,omitempty"`
	PastWeek   bool   `json:"past_week,omitempty"`
	PastMonth  bool   `json:"past_month,omitempty"`
	PastYear   bool   `json:"past_year,omitempty"`
	NextWeek   bool   `json:"next_week,omitempty"`
	NextMonth  bool   `json:"next_month,omitempty"`
	NextYear   bool   `json:"next_year,omitempty"`
	IsEmpty    bool   `json:"is_empty,omitempty"`
	IsNotEmpty bool   `json:"is_not_empty,omitempty"`
}

// PeopleFilter represents a people filter
type PeopleFilter struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// FilesFilter represents a files filter
type FilesFilter struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
}

// RelationFilter represents a relation filter
type RelationFilter struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// FormulaFilter represents a formula filter
type FormulaFilter struct {
	String   *TextFilter     `json:"string,omitempty"`
	Checkbox *CheckboxFilter `json:"checkbox,omitempty"`
	Number   *NumberFilter   `json:"number,omitempty"`
	Date     *DateFilter     `json:"date,omitempty"`
}

// Sort represents a sort configuration
type Sort struct {
	Property  string `json:"property,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Direction string `json:"direction"`
}

// Sort directions
const (
	SortDirectionAscending  = "ascending"
	SortDirectionDescending = "descending"
)

// GetDatabase retrieves a database by ID
func (c *Client) GetDatabase(ctx context.Context, databaseID string) (*Database, error) {
	resp, err := c.makeRequest(ctx, "GET", "/databases/"+databaseID, nil)
	if err != nil {
		return nil, err
	}

	var database Database
	if err := parseResponse(resp, &database); err != nil {
		return nil, fmt.Errorf("failed to parse database response: %w", err)
	}

	return &database, nil
}

// CreateDatabase creates a new database
func (c *Client) CreateDatabase(ctx context.Context, req *CreateDatabaseRequest) (*Database, error) {
	resp, err := c.makeRequest(ctx, "POST", "/databases", req)
	if err != nil {
		return nil, err
	}

	var database Database
	if err := parseResponse(resp, &database); err != nil {
		return nil, fmt.Errorf("failed to parse database response: %w", err)
	}

	return &database, nil
}

// UpdateDatabase updates an existing database
func (c *Client) UpdateDatabase(ctx context.Context, databaseID string, req *UpdateDatabaseRequest) (*Database, error) {
	resp, err := c.makeRequest(ctx, "PATCH", "/databases/"+databaseID, req)
	if err != nil {
		return nil, err
	}

	var database Database
	if err := parseResponse(resp, &database); err != nil {
		return nil, fmt.Errorf("failed to parse database response: %w", err)
	}

	return &database, nil
}

// QueryDatabase queries a database with filters and sorts
func (c *Client) QueryDatabase(ctx context.Context, databaseID string, req *QueryDatabaseRequest) (*PagesListResponse, error) {
	resp, err := c.makeRequest(ctx, "POST", "/databases/"+databaseID+"/query", req)
	if err != nil {
		return nil, err
	}

	var pages PagesListResponse
	if err := parseResponse(resp, &pages); err != nil {
		return nil, fmt.Errorf("failed to parse pages response: %w", err)
	}

	return &pages, nil
}
