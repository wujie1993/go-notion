package notion

import (
	"context"
	"fmt"
)

// Block represents a Notion block
type Block struct {
	Object         string  `json:"object"`
	ID             string  `json:"id"`
	Parent         *Parent `json:"parent,omitempty"`
	CreatedTime    string  `json:"created_time,omitempty"`
	CreatedBy      *User   `json:"created_by,omitempty"`
	LastEditedTime string  `json:"last_edited_time,omitempty"`
	LastEditedBy   *User   `json:"last_edited_by,omitempty"`
	Archived       bool    `json:"archived,omitempty"`
	HasChildren    bool    `json:"has_children,omitempty"`
	Type           string  `json:"type"`

	// Block types
	Paragraph        *ParagraphBlock        `json:"paragraph,omitempty"`
	Heading1         *HeadingBlock          `json:"heading_1,omitempty"`
	Heading2         *HeadingBlock          `json:"heading_2,omitempty"`
	Heading3         *HeadingBlock          `json:"heading_3,omitempty"`
	BulletedListItem *ListItemBlock         `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ListItemBlock         `json:"numbered_list_item,omitempty"`
	Quote            *QuoteBlock            `json:"quote,omitempty"`
	ToDo             *ToDoBlock             `json:"to_do,omitempty"`
	Toggle           *ToggleBlock           `json:"toggle,omitempty"`
	Template         *TemplateBlock         `json:"template,omitempty"`
	Synced           *SyncedBlock           `json:"synced_block,omitempty"`
	ChildPage        *ChildPageBlock        `json:"child_page,omitempty"`
	ChildDatabase    *ChildDatabaseBlock    `json:"child_database,omitempty"`
	Equation         *EquationBlock         `json:"equation,omitempty"`
	Code             *CodeBlock             `json:"code,omitempty"`
	Callout          *CalloutBlock          `json:"callout,omitempty"`
	Divider          map[string]interface{} `json:"divider,omitempty"`
	Breadcrumb       map[string]interface{} `json:"breadcrumb,omitempty"`
	TableOfContents  *TableOfContentsBlock  `json:"table_of_contents,omitempty"`
	ColumnList       *ColumnListBlock       `json:"column_list,omitempty"`
	Column           *ColumnBlock           `json:"column,omitempty"`
	LinkPreview      *LinkPreviewBlock      `json:"link_preview,omitempty"`
	Table            *TableBlock            `json:"table,omitempty"`
	TableRow         *TableRowBlock         `json:"table_row,omitempty"`
	Embed            *EmbedBlock            `json:"embed,omitempty"`
	Bookmark         *BookmarkBlock         `json:"bookmark,omitempty"`
	Image            *FileBlock             `json:"image,omitempty"`
	Video            *FileBlock             `json:"video,omitempty"`
	File             *FileBlock             `json:"file,omitempty"`
	PDF              *FileBlock             `json:"pdf,omitempty"`
	Audio            *FileBlock             `json:"audio,omitempty"`
	LinkToPage       *LinkToPageBlock       `json:"link_to_page,omitempty"`
	Unsupported      map[string]interface{} `json:"unsupported,omitempty"`
}

// ParagraphBlock represents a paragraph block
type ParagraphBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// HeadingBlock represents a heading block
type HeadingBlock struct {
	RichText     []RichText `json:"rich_text"`
	Color        string     `json:"color,omitempty"`
	IsToggleable bool       `json:"is_toggleable,omitempty"`
	Children     []Block    `json:"children,omitempty"`
}

// ListItemBlock represents a list item block
type ListItemBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// QuoteBlock represents a quote block
type QuoteBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// ToDoBlock represents a to-do block
type ToDoBlock struct {
	RichText []RichText `json:"rich_text"`
	Checked  bool       `json:"checked"`
	Color    string     `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// ToggleBlock represents a toggle block
type ToggleBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// TemplateBlock represents a template block
type TemplateBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
}

// SyncedBlock represents a synced block
type SyncedBlock struct {
	SyncedFrom *SyncedFrom `json:"synced_from,omitempty"`
	Children   []Block     `json:"children,omitempty"`
}

// SyncedFrom represents a synced block reference
type SyncedFrom struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id,omitempty"`
}

// ChildPageBlock represents a child page block
type ChildPageBlock struct {
	Title string `json:"title"`
}

// ChildDatabaseBlock represents a child database block
type ChildDatabaseBlock struct {
	Title string `json:"title"`
}

// EquationBlock represents an equation block
type EquationBlock struct {
	Expression string `json:"expression"`
}

// CodeBlock represents a code block
type CodeBlock struct {
	Caption  []RichText `json:"caption,omitempty"`
	RichText []RichText `json:"rich_text"`
	Language string     `json:"language,omitempty"`
}

// CalloutBlock represents a callout block
type CalloutBlock struct {
	RichText []RichText `json:"rich_text"`
	Icon     *Icon      `json:"icon,omitempty"`
	Color    string     `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// TableOfContentsBlock represents a table of contents block
type TableOfContentsBlock struct {
	Color string `json:"color,omitempty"`
}

// ColumnListBlock represents a column list block
type ColumnListBlock struct {
	Children []Block `json:"children,omitempty"`
}

// ColumnBlock represents a column block
type ColumnBlock struct {
	Children []Block `json:"children,omitempty"`
}

// LinkPreviewBlock represents a link preview block
type LinkPreviewBlock struct {
	URL string `json:"url"`
}

// TableBlock represents a table block
type TableBlock struct {
	TableWidth      int     `json:"table_width"`
	HasColumnHeader bool    `json:"has_column_header"`
	HasRowHeader    bool    `json:"has_row_header"`
	Children        []Block `json:"children,omitempty"`
}

// TableRowBlock represents a table row block
type TableRowBlock struct {
	Cells [][]RichText `json:"cells"`
}

// EmbedBlock represents an embed block
type EmbedBlock struct {
	URL     string     `json:"url"`
	Caption []RichText `json:"caption,omitempty"`
}

// BookmarkBlock represents a bookmark block
type BookmarkBlock struct {
	URL     string     `json:"url"`
	Caption []RichText `json:"caption,omitempty"`
}

// FileBlock represents a file/image/video block
type FileBlock struct {
	Caption  []RichText `json:"caption,omitempty"`
	Type     string     `json:"type,omitempty"`
	File     *File      `json:"file,omitempty"`
	External *File      `json:"external,omitempty"`
}

// LinkToPageBlock represents a link to page block
type LinkToPageBlock struct {
	Type       string `json:"type"`
	PageID     string `json:"page_id,omitempty"`
	DatabaseID string `json:"database_id,omitempty"`
}

// BlocksListResponse represents a list of blocks
type BlocksListResponse struct {
	ListResponse
	Results []Block `json:"results"`
}

// AppendBlockChildrenRequest represents a request to append children to a block
type AppendBlockChildrenRequest struct {
	Children []Block `json:"children"`
}

// UpdateBlockRequest represents a request to update a block
type UpdateBlockRequest struct {
	Paragraph        *ParagraphBlock       `json:"paragraph,omitempty"`
	Heading1         *HeadingBlock         `json:"heading_1,omitempty"`
	Heading2         *HeadingBlock         `json:"heading_2,omitempty"`
	Heading3         *HeadingBlock         `json:"heading_3,omitempty"`
	BulletedListItem *ListItemBlock        `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ListItemBlock        `json:"numbered_list_item,omitempty"`
	Quote            *QuoteBlock           `json:"quote,omitempty"`
	ToDo             *ToDoBlock            `json:"to_do,omitempty"`
	Toggle           *ToggleBlock          `json:"toggle,omitempty"`
	Template         *TemplateBlock        `json:"template,omitempty"`
	Equation         *EquationBlock        `json:"equation,omitempty"`
	Code             *CodeBlock            `json:"code,omitempty"`
	Callout          *CalloutBlock         `json:"callout,omitempty"`
	TableOfContents  *TableOfContentsBlock `json:"table_of_contents,omitempty"`
	Embed            *EmbedBlock           `json:"embed,omitempty"`
	Bookmark         *BookmarkBlock        `json:"bookmark,omitempty"`
	Image            *FileBlock            `json:"image,omitempty"`
	Video            *FileBlock            `json:"video,omitempty"`
	File             *FileBlock            `json:"file,omitempty"`
	PDF              *FileBlock            `json:"pdf,omitempty"`
	Audio            *FileBlock            `json:"audio,omitempty"`
	Archived         *bool                 `json:"archived,omitempty"`
}

// GetBlock retrieves a block by ID
func (c *Client) GetBlock(ctx context.Context, blockID string) (*Block, error) {
	resp, err := c.makeRequest(ctx, "GET", "/blocks/"+blockID, nil)
	if err != nil {
		return nil, err
	}

	var block Block
	if err := parseResponse(resp, &block); err != nil {
		return nil, fmt.Errorf("failed to parse block response: %w", err)
	}

	return &block, nil
}

// UpdateBlock updates an existing block
func (c *Client) UpdateBlock(ctx context.Context, blockID string, req *UpdateBlockRequest) (*Block, error) {
	resp, err := c.makeRequest(ctx, "PATCH", "/blocks/"+blockID, req)
	if err != nil {
		return nil, err
	}

	var block Block
	if err := parseResponse(resp, &block); err != nil {
		return nil, fmt.Errorf("failed to parse block response: %w", err)
	}

	return &block, nil
}

// DeleteBlock deletes a block
func (c *Client) DeleteBlock(ctx context.Context, blockID string) (*Block, error) {
	resp, err := c.makeRequest(ctx, "DELETE", "/blocks/"+blockID, nil)
	if err != nil {
		return nil, err
	}

	var block Block
	if err := parseResponse(resp, &block); err != nil {
		return nil, fmt.Errorf("failed to parse block response: %w", err)
	}

	return &block, nil
}

// GetBlockChildren retrieves the children of a block
func (c *Client) GetBlockChildren(ctx context.Context, blockID string, startCursor string, pageSize int) (*BlocksListResponse, error) {
	path := "/blocks/" + blockID + "/children"

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

	var blocks BlocksListResponse
	if err := parseResponse(resp, &blocks); err != nil {
		return nil, fmt.Errorf("failed to parse blocks response: %w", err)
	}

	return &blocks, nil
}

// AppendBlockChildren appends new children to a block
func (c *Client) AppendBlockChildren(ctx context.Context, blockID string, req *AppendBlockChildrenRequest) (*BlocksListResponse, error) {
	resp, err := c.makeRequest(ctx, "PATCH", "/blocks/"+blockID+"/children", req)
	if err != nil {
		return nil, err
	}

	var blocks BlocksListResponse
	if err := parseResponse(resp, &blocks); err != nil {
		return nil, fmt.Errorf("failed to parse blocks response: %w", err)
	}

	return &blocks, nil
}
