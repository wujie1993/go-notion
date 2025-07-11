package notion

import (
	"testing"
)

func TestNewText(t *testing.T) {
	text := NewText("Hello, World!")

	if text.Type != "text" {
		t.Errorf("Expected type 'text', got '%s'", text.Type)
	}

	if text.Text.Content != "Hello, World!" {
		t.Errorf("Expected content 'Hello, World!', got '%s'", text.Text.Content)
	}

	if text.PlainText != "Hello, World!" {
		t.Errorf("Expected plain text 'Hello, World!', got '%s'", text.PlainText)
	}
}

func TestNewTextWithLink(t *testing.T) {
	text := NewTextWithLink("Notion", "https://notion.so")

	if text.Type != "text" {
		t.Errorf("Expected type 'text', got '%s'", text.Type)
	}

	if text.Text.Content != "Notion" {
		t.Errorf("Expected content 'Notion', got '%s'", text.Text.Content)
	}

	if text.Text.Link.URL != "https://notion.so" {
		t.Errorf("Expected URL 'https://notion.so', got '%s'", text.Text.Link.URL)
	}

	if text.Href != "https://notion.so" {
		t.Errorf("Expected href 'https://notion.so', got '%s'", text.Href)
	}
}

func TestNewParagraphBlock(t *testing.T) {
	richText := []RichText{NewText("This is a paragraph")}
	block := NewParagraphBlock(richText)

	if block.Type != "paragraph" {
		t.Errorf("Expected type 'paragraph', got '%s'", block.Type)
	}

	if len(block.Paragraph.RichText) != 1 {
		t.Errorf("Expected 1 rich text element, got %d", len(block.Paragraph.RichText))
	}

	if block.Paragraph.RichText[0].PlainText != "This is a paragraph" {
		t.Errorf("Expected plain text 'This is a paragraph', got '%s'", block.Paragraph.RichText[0].PlainText)
	}
}

func TestNewTableBlock(t *testing.T) {
	block := NewTableBlock(3, true, false)

	if block.Type != "table" {
		t.Errorf("Expected type 'table', got '%s'", block.Type)
	}

	if block.Table.TableWidth != 3 {
		t.Errorf("Expected table width 3, got %d", block.Table.TableWidth)
	}

	if !block.Table.HasColumnHeader {
		t.Errorf("Expected has column header to be true")
	}

	if block.Table.HasRowHeader {
		t.Errorf("Expected has row header to be false")
	}

	if len(block.Table.Children) != 0 {
		t.Errorf("Expected empty children, got %d", len(block.Table.Children))
	}
}

func TestNewTableRowBlock(t *testing.T) {
	cells := [][]RichText{
		{NewText("Cell 1")},
		{NewText("Cell 2")},
		{NewText("Cell 3")},
	}
	block := NewTableRowBlock(cells)

	if block.Type != "table_row" {
		t.Errorf("Expected type 'table_row', got '%s'", block.Type)
	}

	if len(block.TableRow.Cells) != 3 {
		t.Errorf("Expected 3 cells, got %d", len(block.TableRow.Cells))
	}

	if block.TableRow.Cells[0][0].PlainText != "Cell 1" {
		t.Errorf("Expected first cell to be 'Cell 1', got '%s'", block.TableRow.Cells[0][0].PlainText)
	}
}

func TestNewHeadingBlocks(t *testing.T) {
	richText := []RichText{NewText("Heading")}

	h1 := NewHeading1Block(richText)
	if h1.Type != "heading_1" {
		t.Errorf("Expected type 'heading_1', got '%s'", h1.Type)
	}

	h2 := NewHeading2Block(richText)
	if h2.Type != "heading_2" {
		t.Errorf("Expected type 'heading_2', got '%s'", h2.Type)
	}

	h3 := NewHeading3Block(richText)
	if h3.Type != "heading_3" {
		t.Errorf("Expected type 'heading_3', got '%s'", h3.Type)
	}
}

func TestNewListItemBlocks(t *testing.T) {
	richText := []RichText{NewText("List item")}

	bulleted := NewBulletedListItemBlock(richText)
	if bulleted.Type != "bulleted_list_item" {
		t.Errorf("Expected type 'bulleted_list_item', got '%s'", bulleted.Type)
	}

	numbered := NewNumberedListItemBlock(richText)
	if numbered.Type != "numbered_list_item" {
		t.Errorf("Expected type 'numbered_list_item', got '%s'", numbered.Type)
	}
}

func TestNewToDoBlock(t *testing.T) {
	richText := []RichText{NewText("Task")}
	block := NewToDoBlock(richText, true)

	if block.Type != "to_do" {
		t.Errorf("Expected type 'to_do', got '%s'", block.Type)
	}

	if !block.ToDo.Checked {
		t.Error("Expected checked to be true")
	}
}

func TestNewCodeBlock(t *testing.T) {
	richText := []RichText{NewText("console.log('Hello');")}
	block := NewCodeBlock(richText, "javascript")

	if block.Type != "code" {
		t.Errorf("Expected type 'code', got '%s'", block.Type)
	}

	if block.Code.Language != "javascript" {
		t.Errorf("Expected language 'javascript', got '%s'", block.Code.Language)
	}
}

func TestNewDividerBlock(t *testing.T) {
	block := NewDividerBlock()

	if block.Type != "divider" {
		t.Errorf("Expected type 'divider', got '%s'", block.Type)
	}

	if block.Divider == nil {
		t.Error("Expected divider to be non-nil")
	}
}

func TestNewProperties(t *testing.T) {
	// Test title property
	titleProp := NewTitleProperty([]RichText{NewText("Test Title")})
	if titleProp.Type != "title" {
		t.Errorf("Expected type 'title', got '%s'", titleProp.Type)
	}

	// Test number property
	numberProp := NewNumberProperty(42.5)
	if numberProp.Type != "number" {
		t.Errorf("Expected type 'number', got '%s'", numberProp.Type)
	}
	if *numberProp.Number != 42.5 {
		t.Errorf("Expected number 42.5, got %f", *numberProp.Number)
	}

	// Test checkbox property
	checkboxProp := NewCheckboxProperty(true)
	if checkboxProp.Type != "checkbox" {
		t.Errorf("Expected type 'checkbox', got '%s'", checkboxProp.Type)
	}
	if !checkboxProp.Checkbox {
		t.Error("Expected checkbox to be true")
	}

	// Test URL property
	urlProp := NewURLProperty("https://example.com")
	if urlProp.Type != "url" {
		t.Errorf("Expected type 'url', got '%s'", urlProp.Type)
	}
	if urlProp.URL != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got '%s'", urlProp.URL)
	}

	// Test email property
	emailProp := NewEmailProperty("test@example.com")
	if emailProp.Type != "email" {
		t.Errorf("Expected type 'email', got '%s'", emailProp.Type)
	}
	if emailProp.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", emailProp.Email)
	}
}

func TestNewParents(t *testing.T) {
	// Test page parent
	pageParent := NewPageParent("test-page-id")
	if pageParent.Type != "page_id" {
		t.Errorf("Expected type 'page_id', got '%s'", pageParent.Type)
	}
	if pageParent.PageID != "test-page-id" {
		t.Errorf("Expected page ID 'test-page-id', got '%s'", pageParent.PageID)
	}

	// Test database parent
	dbParent := NewDatabaseParent("test-db-id")
	if dbParent.Type != "database_id" {
		t.Errorf("Expected type 'database_id', got '%s'", dbParent.Type)
	}
	if dbParent.DatabaseID != "test-db-id" {
		t.Errorf("Expected database ID 'test-db-id', got '%s'", dbParent.DatabaseID)
	}

	// Test workspace parent
	workspaceParent := NewWorkspaceParent()
	if workspaceParent.Type != "workspace" {
		t.Errorf("Expected type 'workspace', got '%s'", workspaceParent.Type)
	}
	if !workspaceParent.Workspace {
		t.Error("Expected workspace to be true")
	}
}

func TestNewIcons(t *testing.T) {
	// Test emoji icon
	emojiIcon := NewEmojiIcon("📄")
	if emojiIcon.Type != "emoji" {
		t.Errorf("Expected type 'emoji', got '%s'", emojiIcon.Type)
	}
	if emojiIcon.Emoji != "📄" {
		t.Errorf("Expected emoji '📄', got '%s'", emojiIcon.Emoji)
	}

	// Test external file icon
	externalIcon := NewExternalFileIcon("https://example.com/icon.png")
	if externalIcon.Type != "external" {
		t.Errorf("Expected type 'external', got '%s'", externalIcon.Type)
	}
	if externalIcon.External.URL != "https://example.com/icon.png" {
		t.Errorf("Expected URL 'https://example.com/icon.png', got '%s'", externalIcon.External.URL)
	}
}

func TestNewCovers(t *testing.T) {
	cover := NewExternalCover("https://example.com/cover.jpg")
	if cover.Type != "external" {
		t.Errorf("Expected type 'external', got '%s'", cover.Type)
	}
	if cover.External.URL != "https://example.com/cover.jpg" {
		t.Errorf("Expected URL 'https://example.com/cover.jpg', got '%s'", cover.External.URL)
	}
}

func TestGenerateID(t *testing.T) {
	id1 := GenerateID()
	id2 := GenerateID()

	if id1 == id2 {
		t.Error("Expected different IDs, got the same")
	}

	if len(id1) == 0 {
		t.Error("Expected non-empty ID")
	}
}
