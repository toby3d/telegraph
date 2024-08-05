package telegraph_test

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"source.toby3d.me/toby3d/telegraph"
	"source.toby3d.me/toby3d/telegraph/internal/util"
)

func Example() {
	client := http.DefaultClient

	account, err := telegraph.CreateAccount{
		AuthorURL:  nil,
		AuthorName: util.Must(telegraph.NewAuthorName("Anonymous")),
		ShortName:  *util.Must(telegraph.NewShortName("Sandbox")),
	}.Do(context.Background(), client)
	if err != nil {
		log.Fatalln("cannot create account:", err)
	}

	page, err := telegraph.CreatePage{
		AuthorURL:   nil,
		AccessToken: account.AccessToken,
		Title:       *util.Must(telegraph.NewTitle("Sample Page")),
		AuthorName:  &account.AuthorName,
		Content: []telegraph.Node{{
			Element: &telegraph.NodeElement{
				Tag:      telegraph.P,
				Children: []telegraph.Node{{Text: "Hello, World!"}},
			},
		}},
		ReturnContent: true,
	}.Do(context.Background(), client)
	if err != nil {
		log.Fatalln("cannot create page:", err)
	}

	fmt.Printf("'%s' by %s\n%s", page.Title, page.AuthorName, page.Content[0])
	// Output: 'Sample Page' by Anonymous
	// <p>Hello, World!</p>
}