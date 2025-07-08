# go-notion

A comprehensive Go client library for the Notion API, inspired by the official [Notion SDK for JavaScript](https://github.com/makenotion/notion-sdk-js).

## Features

- 🚀 **Complete API Coverage**: Support for all major Notion API endpoints
- 🔒 **Type Safe**: Fully typed with Go structs for all Notion objects
- 🏗️ **Easy to Use**: Simple and intuitive API design
- 🧪 **Well Tested**: Comprehensive test suite
- 📚 **Rich Documentation**: Extensive examples and documentation
- 🛠️ **Helper Functions**: Convenient utilities for creating common objects

## Installation

```bash
go get github.com/wujie1993/go-notion
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/wujie1993/go-notion"
)

func main() {
    // Create a new client
    client := notion.NewClient(os.Getenv("NOTION_API_KEY"))
    ctx := context.Background()

    // Get current user
    user, err := client.GetMe(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Hello, %s!", user.Name)
}
```

## API Reference

### Client Configuration

```go
// Basic client
client := notion.NewClient("your-api-key")

// Client with custom options
client := notion.NewClient("your-api-key",
    notion.WithHTTPClient(customHTTPClient),
    notion.WithBaseURL("https://api.notion.com/v1"),
    notion.WithVersion("2022-06-28"),
)
```

### Pages

```go
// Get a page
page, err := client.GetPage(ctx, "page-id")

// Create a page
page, err := client.CreatePage(ctx, &notion.CreatePageRequest{
    Parent: notion.NewPageParent("parent-page-id"),
    Properties: map[string]notion.PageProperty{
        "title": notion.NewTitleProperty([]notion.RichText{
            notion.NewText("My New Page"),
        }),
    },
})

// Update a page
page, err := client.UpdatePage(ctx, "page-id", &notion.UpdatePageRequest{
    Properties: map[string]notion.PageProperty{
        "title": notion.NewTitleProperty([]notion.RichText{
            notion.NewText("Updated Title"),
        }),
    },
})
```

### Databases

```go
// Get a database
database, err := client.GetDatabase(ctx, "database-id")

// Query a database
pages, err := client.QueryDatabase(ctx, "database-id", &notion.QueryDatabaseRequest{
    Filter: &notion.Filter{
        Property: "Status",
        Select: &notion.SelectFilter{
            Equals: "Done",
        },
    },
    Sorts: []notion.Sort{
        {
            Property:  "Created",
            Direction: notion.SortDirectionDescending,
        },
    },
})

// Create a database
database, err := client.CreateDatabase(ctx, &notion.CreateDatabaseRequest{
    Parent: notion.NewPageParent("parent-page-id"),
    Title: []notion.RichText{
        notion.NewText("My Database"),
    },
    Properties: map[string]notion.DatabaseProperty{
        "Name": {
            Type:  notion.PropertyTypeTitle,
            Title: map[string]interface{}{},
        },
        "Status": {
            Type: notion.PropertyTypeSelect,
            Select: &notion.SelectProperty{
                Options: []notion.SelectOption{
                    {Name: "Not started", Color: notion.ColorRed},
                    {Name: "In progress", Color: notion.ColorYellow},
                    {Name: "Done", Color: notion.ColorGreen},
                },
            },
        },
    },
})
```

### Blocks

```go
// Get block children
blocks, err := client.GetBlockChildren(ctx, "block-id", "", 0)

// Append blocks to a page
blocks, err := client.AppendBlockChildren(ctx, "page-id", &notion.AppendBlockChildrenRequest{
    Children: []notion.Block{
        *notion.NewParagraphBlock([]notion.RichText{
            notion.NewText("Hello, World!"),
        }),
        *notion.NewHeading1Block([]notion.RichText{
            notion.NewText("This is a heading"),
        }),
        *notion.NewBulletedListItemBlock([]notion.RichText{
            notion.NewText("First item"),
        }),
        *notion.NewNumberedListItemBlock([]notion.RichText{
            notion.NewText("Numbered item"),
        }),
        *notion.NewToDoBlock([]notion.RichText{
            notion.NewText("Task to complete"),
        }, false),
        *notion.NewCodeBlock([]notion.RichText{
            notion.NewText("console.log('Hello, World!');"),
        }, "javascript"),
        *notion.NewDividerBlock(),
    },
})

// Update a block
block, err := client.UpdateBlock(ctx, "block-id", &notion.UpdateBlockRequest{
    Paragraph: &notion.ParagraphBlock{
        RichText: []notion.RichText{
            notion.NewText("Updated content"),
        },
    },
})

// Delete a block
block, err := client.DeleteBlock(ctx, "block-id")
```

### Users

```go
// Get current user
user, err := client.GetMe(ctx)

// Get a user by ID
user, err := client.GetUser(ctx, "user-id")

// List all users
users, err := client.ListUsers(ctx, "", 0)
```

### Search

```go
// Search for pages and databases
results, err := client.Search(ctx, &notion.SearchRequest{
    Query: "meeting notes",
    Filter: &notion.SearchFilter{
        Value:    "page",
        Property: "object",
    },
    Sort: &notion.SearchSort{
        Direction: notion.SortDirectionDescending,
        Timestamp: "last_edited_time",
    },
})
```

## Helper Functions

The library provides many helper functions to make working with Notion objects easier:

### Rich Text Helpers

```go
// Plain text
text := notion.NewText("Hello, World!")

// Text with link
text := notion.NewTextWithLink("Notion", "https://notion.so")

// Text with annotations
text := notion.NewAnnotatedText("Bold text", &notion.Annotations{
    Bold: true,
    Color: notion.ColorRed,
})
```

### Block Helpers

```go
// Various block types
paragraph := notion.NewParagraphBlock([]notion.RichText{notion.NewText("Content")})
heading1 := notion.NewHeading1Block([]notion.RichText{notion.NewText("Heading")})
heading2 := notion.NewHeading2Block([]notion.RichText{notion.NewText("Heading")})
heading3 := notion.NewHeading3Block([]notion.RichText{notion.NewText("Heading")})
bulletList := notion.NewBulletedListItemBlock([]notion.RichText{notion.NewText("Item")})
numberList := notion.NewNumberedListItemBlock([]notion.RichText{notion.NewText("Item")})
todo := notion.NewToDoBlock([]notion.RichText{notion.NewText("Task")}, false)
code := notion.NewCodeBlock([]notion.RichText{notion.NewText("code")}, "go")
quote := notion.NewQuoteBlock([]notion.RichText{notion.NewText("Quote")})
callout := notion.NewCalloutBlock([]notion.RichText{notion.NewText("Note")}, notion.NewEmojiIcon("💡"))
divider := notion.NewDividerBlock()
```

### Property Helpers

```go
// Page properties
title := notion.NewTitleProperty([]notion.RichText{notion.NewText("Title")})
richText := notion.NewRichTextProperty([]notion.RichText{notion.NewText("Content")})
number := notion.NewNumberProperty(42.5)
checkbox := notion.NewCheckboxProperty(true)
url := notion.NewURLProperty("https://example.com")
email := notion.NewEmailProperty("test@example.com")
phone := notion.NewPhoneNumberProperty("+1234567890")
select := notion.NewSelectProperty(notion.SelectOption{Name: "Option", Color: notion.ColorBlue})
multiSelect := notion.NewMultiSelectProperty([]notion.SelectOption{
    {Name: "Tag1", Color: notion.ColorRed},
    {Name: "Tag2", Color: notion.ColorGreen},
})
date := notion.NewDateProperty(notion.Date{Start: "2023-12-01"})
people := notion.NewPeopleProperty([]notion.User{{ID: "user-id"}})
relation := notion.NewRelationProperty([]notion.Relation{{ID: "related-page-id"}})
```

### Parent Helpers

```go
// Parent objects
pageParent := notion.NewPageParent("page-id")
databaseParent := notion.NewDatabaseParent("database-id")
workspaceParent := notion.NewWorkspaceParent()
```

### Icon and Cover Helpers

```go
// Icons
emojiIcon := notion.NewEmojiIcon("📄")
fileIcon := notion.NewExternalFileIcon("https://example.com/icon.png")

// Covers
cover := notion.NewExternalCover("https://example.com/cover.jpg")
```

## Examples

Check out the [examples](./examples) directory for more comprehensive examples:

- [Basic Usage](./examples/basic/main.go) - Getting started with the basics

## Authentication

To use the Notion API, you need to create an integration and get an API key:

1. Go to [https://www.notion.so/my-integrations](https://www.notion.so/my-integrations)
2. Click "New integration"
3. Give your integration a name and select capabilities
4. Copy the "Internal Integration Token"
5. Share the pages/databases you want to access with your integration

Set the API key as an environment variable:

```bash
export NOTION_API_KEY="your-integration-token"
```

## Error Handling

The client returns detailed error information:

```go
page, err := client.GetPage(ctx, "invalid-id")
if err != nil {
    if notionErr, ok := err.(*notion.Error); ok {
        log.Printf("Notion API error: %s (status: %d, code: %s)", 
            notionErr.Message, notionErr.Status, notionErr.Code)
    } else {
        log.Printf("Other error: %v", err)
    }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Related Projects

- [Official Notion SDK for JavaScript](https://github.com/makenotion/notion-sdk-js)
- [Notion API Documentation](https://developers.notion.com/)
