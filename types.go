package notion

import (
	"time"

	"github.com/google/uuid"
)

// Object types
const (
	ObjectTypeBlock    = "block"
	ObjectTypePage     = "page"
	ObjectTypeDatabase = "database"
	ObjectTypeUser     = "user"
	ObjectTypeList     = "list"
)

// Common types

// Object represents a common Notion object
type Object struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

// User represents a Notion user
type User struct {
	Object    string `json:"object"`
	ID        string `json:"id"`
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Person    *struct {
		Email string `json:"email,omitempty"`
	} `json:"person,omitempty"`
	Bot *struct {
		Owner         *User  `json:"owner,omitempty"`
		WorkspaceName string `json:"workspace_name,omitempty"`
	} `json:"bot,omitempty"`
}

// Parent represents a parent object
type Parent struct {
	Type       string `json:"type"`
	PageID     string `json:"page_id,omitempty"`
	DatabaseID string `json:"database_id,omitempty"`
	BlockID    string `json:"block_id,omitempty"`
	Workspace  bool   `json:"workspace,omitempty"`
}

// Icon represents an icon
type Icon struct {
	Type     string `json:"type"`
	Emoji    string `json:"emoji,omitempty"`
	External *File  `json:"external,omitempty"`
	File     *File  `json:"file,omitempty"`
}

// Cover represents a cover image
type Cover struct {
	Type     string `json:"type"`
	External *File  `json:"external,omitempty"`
	File     *File  `json:"file,omitempty"`
}

// File represents a file
type File struct {
	URL        string     `json:"url"`
	ExpiryTime *time.Time `json:"expiry_time,omitempty"`
}

// RichText represents rich text content
type RichText struct {
	Type        string       `json:"type"`
	Text        *Text        `json:"text,omitempty"`
	Mention     *Mention     `json:"mention,omitempty"`
	Equation    *Equation    `json:"equation,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	PlainText   string       `json:"plain_text,omitempty"`
	Href        string       `json:"href,omitempty"`
}

// Text represents text content
type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

// Link represents a link
type Link struct {
	URL string `json:"url"`
}

// Mention represents a mention
type Mention struct {
	Type            string  `json:"type"`
	User            *User   `json:"user,omitempty"`
	Page            *Object `json:"page,omitempty"`
	Database        *Object `json:"database,omitempty"`
	Date            *Date   `json:"date,omitempty"`
	LinkPreview     *Link   `json:"link_preview,omitempty"`
	TemplateMention *struct {
		Type                string `json:"type"`
		TemplateMentionDate string `json:"template_mention_date,omitempty"`
		TemplateMentionUser string `json:"template_mention_user,omitempty"`
	} `json:"template_mention,omitempty"`
}

// Equation represents an equation
type Equation struct {
	Expression string `json:"expression"`
}

// Annotations represents text annotations
type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

// Date represents a date or date range
type Date struct {
	Start    string `json:"start"`
	End      string `json:"end,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`
}

// ListResponse represents a paginated list response
type ListResponse struct {
	Object     string `json:"object"`
	NextCursor string `json:"next_cursor,omitempty"`
	HasMore    bool   `json:"has_more"`
}

// Color constants
const (
	ColorDefault          = "default"
	ColorGray             = "gray"
	ColorBrown            = "brown"
	ColorOrange           = "orange"
	ColorYellow           = "yellow"
	ColorGreen            = "green"
	ColorBlue             = "blue"
	ColorPurple           = "purple"
	ColorPink             = "pink"
	ColorRed              = "red"
	ColorGrayBackground   = "gray_background"
	ColorBrownBackground  = "brown_background"
	ColorOrangeBackground = "orange_background"
	ColorYellowBackground = "yellow_background"
	ColorGreenBackground  = "green_background"
	ColorBlueBackground   = "blue_background"
	ColorPurpleBackground = "purple_background"
	ColorPinkBackground   = "pink_background"
	ColorRedBackground    = "red_background"
)

// GenerateID generates a new UUID for Notion objects
func GenerateID() string {
	return uuid.New().String()
}
