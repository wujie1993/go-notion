package notion

// Helper functions for creating common rich text and property values

// NewText creates a new RichText with plain text
func NewText(content string) RichText {
	return RichText{
		Type: "text",
		Text: &Text{
			Content: content,
		},
		PlainText: content,
	}
}

// NewTextWithLink creates a new RichText with a link
func NewTextWithLink(content, url string) RichText {
	return RichText{
		Type: "text",
		Text: &Text{
			Content: content,
			Link: &Link{
				URL: url,
			},
		},
		PlainText: content,
		Href:      url,
	}
}

// NewAnnotatedText creates a new RichText with annotations
func NewAnnotatedText(content string, annotations *Annotations) RichText {
	return RichText{
		Type: "text",
		Text: &Text{
			Content: content,
		},
		Annotations: annotations,
		PlainText:   content,
	}
}

// NewParagraphBlock creates a new paragraph block
func NewParagraphBlock(richText []RichText) *Block {
	return &Block{
		Type: "paragraph",
		Paragraph: &ParagraphBlock{
			RichText: richText,
		},
	}
}

// NewHeading1Block creates a new heading 1 block
func NewHeading1Block(richText []RichText) *Block {
	return &Block{
		Type: "heading_1",
		Heading1: &HeadingBlock{
			RichText: richText,
		},
	}
}

// NewHeading2Block creates a new heading 2 block
func NewHeading2Block(richText []RichText) *Block {
	return &Block{
		Type: "heading_2",
		Heading2: &HeadingBlock{
			RichText: richText,
		},
	}
}

// NewHeading3Block creates a new heading 3 block
func NewHeading3Block(richText []RichText) *Block {
	return &Block{
		Type: "heading_3",
		Heading3: &HeadingBlock{
			RichText: richText,
		},
	}
}

// NewBulletedListItemBlock creates a new bulleted list item block
func NewBulletedListItemBlock(richText []RichText) *Block {
	return &Block{
		Type: "bulleted_list_item",
		BulletedListItem: &ListItemBlock{
			RichText: richText,
		},
	}
}

// NewNumberedListItemBlock creates a new numbered list item block
func NewNumberedListItemBlock(richText []RichText) *Block {
	return &Block{
		Type: "numbered_list_item",
		NumberedListItem: &ListItemBlock{
			RichText: richText,
		},
	}
}

// NewToDoBlock creates a new to-do block
func NewToDoBlock(richText []RichText, checked bool) *Block {
	return &Block{
		Type: "to_do",
		ToDo: &ToDoBlock{
			RichText: richText,
			Checked:  checked,
		},
	}
}

// NewCodeBlock creates a new code block
func NewCodeBlock(richText []RichText, language string) *Block {
	return &Block{
		Type: "code",
		Code: &CodeBlock{
			RichText: richText,
			Language: language,
		},
	}
}

// NewQuoteBlock creates a new quote block
func NewQuoteBlock(richText []RichText) *Block {
	return &Block{
		Type: "quote",
		Quote: &QuoteBlock{
			RichText: richText,
		},
	}
}

// NewCalloutBlock creates a new callout block
func NewCalloutBlock(richText []RichText, icon *Icon) *Block {
	return &Block{
		Type: "callout",
		Callout: &CalloutBlock{
			RichText: richText,
			Icon:     icon,
		},
	}
}

// NewDividerBlock creates a new divider block
func NewDividerBlock() *Block {
	return &Block{
		Type:    "divider",
		Divider: map[string]interface{}{},
	}
}

// NewTitleProperty creates a new title property
func NewTitleProperty(title []RichText) PageProperty {
	return PageProperty{
		Type:  "title",
		Title: title,
	}
}

// NewRichTextProperty creates a new rich text property
func NewRichTextProperty(richText []RichText) PageProperty {
	return PageProperty{
		Type:     "rich_text",
		RichText: richText,
	}
}

// NewNumberProperty creates a new number property
func NewNumberProperty(number float64) PageProperty {
	return PageProperty{
		Type:   "number",
		Number: &number,
	}
}

// NewSelectProperty creates a new select property
func NewSelectProperty(option SelectOption) PageProperty {
	return PageProperty{
		Type:   "select",
		Select: &option,
	}
}

// NewMultiSelectProperty creates a new multi-select property
func NewMultiSelectProperty(options []SelectOption) PageProperty {
	return PageProperty{
		Type:        "multi_select",
		MultiSelect: options,
	}
}

// NewDateProperty creates a new date property
func NewDateProperty(date Date) PageProperty {
	return PageProperty{
		Type: "date",
		Date: &date,
	}
}

// NewCheckboxProperty creates a new checkbox property
func NewCheckboxProperty(checked bool) PageProperty {
	return PageProperty{
		Type:     "checkbox",
		Checkbox: checked,
	}
}

// NewURLProperty creates a new URL property
func NewURLProperty(url string) PageProperty {
	return PageProperty{
		Type: "url",
		URL:  url,
	}
}

// NewEmailProperty creates a new email property
func NewEmailProperty(email string) PageProperty {
	return PageProperty{
		Type:  "email",
		Email: email,
	}
}

// NewPhoneNumberProperty creates a new phone number property
func NewPhoneNumberProperty(phoneNumber string) PageProperty {
	return PageProperty{
		Type:        "phone_number",
		PhoneNumber: phoneNumber,
	}
}

// NewPeopleProperty creates a new people property
func NewPeopleProperty(people []User) PageProperty {
	return PageProperty{
		Type:   "people",
		People: people,
	}
}

// NewRelationProperty creates a new relation property
func NewRelationProperty(relations []Relation) PageProperty {
	return PageProperty{
		Type:     "relation",
		Relation: relations,
	}
}

// NewPageParent creates a new page parent
func NewPageParent(pageID string) *Parent {
	return &Parent{
		Type:   "page_id",
		PageID: pageID,
	}
}

// NewDatabaseParent creates a new database parent
func NewDatabaseParent(databaseID string) *Parent {
	return &Parent{
		Type:       "database_id",
		DatabaseID: databaseID,
	}
}

// NewWorkspaceParent creates a new workspace parent
func NewWorkspaceParent() *Parent {
	return &Parent{
		Type:      "workspace",
		Workspace: true,
	}
}

// NewEmojiIcon creates a new emoji icon
func NewEmojiIcon(emoji string) *Icon {
	return &Icon{
		Type:  "emoji",
		Emoji: emoji,
	}
}

// NewExternalFileIcon creates a new external file icon
func NewExternalFileIcon(url string) *Icon {
	return &Icon{
		Type: "external",
		External: &File{
			URL: url,
		},
	}
}

// NewExternalCover creates a new external cover
func NewExternalCover(url string) *Cover {
	return &Cover{
		Type: "external",
		External: &File{
			URL: url,
		},
	}
}

// Property type constants
const (
	PropertyTypeTitle          = "title"
	PropertyTypeRichText       = "rich_text"
	PropertyTypeNumber         = "number"
	PropertyTypeSelect         = "select"
	PropertyTypeMultiSelect    = "multi_select"
	PropertyTypeDate           = "date"
	PropertyTypePeople         = "people"
	PropertyTypeFiles          = "files"
	PropertyTypeCheckbox       = "checkbox"
	PropertyTypeURL            = "url"
	PropertyTypeEmail          = "email"
	PropertyTypePhoneNumber    = "phone_number"
	PropertyTypeFormula        = "formula"
	PropertyTypeRelation       = "relation"
	PropertyTypeRollup         = "rollup"
	PropertyTypeCreatedTime    = "created_time"
	PropertyTypeCreatedBy      = "created_by"
	PropertyTypeLastEditedTime = "last_edited_time"
	PropertyTypeLastEditedBy   = "last_edited_by"
	PropertyTypeStatus         = "status"
)

// Block type constants
const (
	BlockTypeParagraph        = "paragraph"
	BlockTypeHeading1         = "heading_1"
	BlockTypeHeading2         = "heading_2"
	BlockTypeHeading3         = "heading_3"
	BlockTypeBulletedListItem = "bulleted_list_item"
	BlockTypeNumberedListItem = "numbered_list_item"
	BlockTypeQuote            = "quote"
	BlockTypeToDo             = "to_do"
	BlockTypeToggle           = "toggle"
	BlockTypeTemplate         = "template"
	BlockTypeSynced           = "synced_block"
	BlockTypeChildPage        = "child_page"
	BlockTypeChildDatabase    = "child_database"
	BlockTypeEquation         = "equation"
	BlockTypeCode             = "code"
	BlockTypeCallout          = "callout"
	BlockTypeDivider          = "divider"
	BlockTypeBreadcrumb       = "breadcrumb"
	BlockTypeTableOfContents  = "table_of_contents"
	BlockTypeColumnList       = "column_list"
	BlockTypeColumn           = "column"
	BlockTypeLinkPreview      = "link_preview"
	BlockTypeTable            = "table"
	BlockTypeTableRow         = "table_row"
	BlockTypeEmbed            = "embed"
	BlockTypeBookmark         = "bookmark"
	BlockTypeImage            = "image"
	BlockTypeVideo            = "video"
	BlockTypeFile             = "file"
	BlockTypePDF              = "pdf"
	BlockTypeAudio            = "audio"
	BlockTypeLinkToPage       = "link_to_page"
	BlockTypeUnsupported      = "unsupported"
)
