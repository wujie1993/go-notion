package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/wujie1993/go-notion"
)

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("NOTION_API_KEY")
	if apiKey == "" {
		log.Fatal("NOTION_API_KEY environment variable is required")
	}

	// Create a new Notion client
	client := notion.NewClient(apiKey)
	ctx := context.Background()

	// Example 1: Get current user
	fmt.Println("=== Getting current user ===")
	user, err := client.GetMe(ctx)
	if err != nil {
		log.Printf("Error getting current user: %v", err)
	} else {
		fmt.Printf("Current user: %s (%s)\n", user.Name, user.ID)
	}

	// Example 2: Search for pages and databases
	fmt.Println("\n=== Searching for content ===")
	searchReq := &notion.SearchRequest{
		Query:    "example",
		PageSize: 5,
	}

	searchResp, err := client.Search(ctx, searchReq)
	if err != nil {
		log.Printf("Error searching: %v", err)
	} else {
		fmt.Printf("Found %d results\n", len(searchResp.Results))
		for i, result := range searchResp.Results {
			if result.Page != nil {
				fmt.Printf("  %d. Page: %s (ID: %s)\n", i+1, getPageTitle(result.Page), result.Page.ID)
			} else if result.Database != nil {
				fmt.Printf("  %d. Database: %s (ID: %s)\n", i+1, getDatabaseTitle(result.Database), result.Database.ID)
			}
		}
	}

	// Example 3: Create a simple page (uncomment and modify as needed)
	/*
		fmt.Println("\n=== Creating a new page ===")
		createPageReq := &notion.CreatePageRequest{
			Parent: notion.NewPageParent("YOUR_PARENT_PAGE_ID"), // Replace with actual page ID
			Properties: map[string]notion.PageProperty{
				"title": notion.NewTitleProperty([]notion.RichText{
					notion.NewText("Example Page from Go"),
				}),
			},
			Children: []notion.Block{
				*notion.NewParagraphBlock([]notion.RichText{
					notion.NewText("This page was created using the Go Notion API client!"),
				}),
				*notion.NewHeading1Block([]notion.RichText{
					notion.NewText("Features"),
				}),
				*notion.NewBulletedListItemBlock([]notion.RichText{
					notion.NewText("Easy to use Go client"),
				}),
				*notion.NewBulletedListItemBlock([]notion.RichText{
					notion.NewText("Support for all major Notion API features"),
				}),
				*notion.NewBulletedListItemBlock([]notion.RichText{
					notion.NewText("Type-safe and well-documented"),
				}),
			},
		}

		page, err := client.CreatePage(ctx, createPageReq)
		if err != nil {
			log.Printf("Error creating page: %v", err)
		} else {
			fmt.Printf("Created page: %s (ID: %s)\n", getPageTitle(page), page.ID)
			fmt.Printf("Page URL: %s\n", page.URL)
		}
	*/

	fmt.Println("\n=== Example completed ===")
}

// Helper function to get page title
func getPageTitle(page *notion.Page) string {
	for _, prop := range page.Properties {
		if prop.Type == "title" && len(prop.Title) > 0 {
			return prop.Title[0].PlainText
		}
	}
	return "Untitled"
}

// Helper function to get database title
func getDatabaseTitle(database *notion.Database) string {
	if len(database.Title) > 0 {
		return database.Title[0].PlainText
	}
	return "Untitled Database"
}
