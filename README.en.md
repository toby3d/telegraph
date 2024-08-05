# telegraph [![Go Reference](https://pkg.go.dev/badge/source.toby3d.me/toby3d/telegraph/v2.svg)](https://pkg.go.dev/source.toby3d.me/toby3d/telegraph/v2)
![](assets/cover.jpg)

A simple package [with minimum official dependencies](v2/go.mod) to work with [Telegraph API](https://telegra.ph/api).

Download:
```bash
$ go get -u source.toby3d.me/toby3d/telegraph/v2
```

Import:
```
import "source.toby3d.me/toby3d/telegraph/v2"
```

Fill commands structs and execute it:
```go
package main

import (
  "context"
  "fmt"
  "log"
  "net/http"

  "source.toby3d.me/toby3d/telegraph/v2"
)

func Must[T any](v T, err error) T {
  if err != nil {
    panic(err)
  }

  return v
}

func main() {
  client := http.DefaultClient

  account, err := telegraph.CreateAccount{
    AuthorURL:  nil,
    AuthorName: Must(telegraph.NewAuthorName("Anonymous")),
    ShortName:  *Must(telegraph.NewShortName("Sandbox")),
  }.Do(context.Background(), client)
  if err != nil {
    log.Fatalln("cannot create account:", err)
  }

  page, err := telegraph.CreatePage{
    AuthorURL:   nil,
    AccessToken: account.AccessToken,
    Title:       *Must(telegraph.NewTitle("Sample Page")),
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
  // 'Sample Page' by Anonymous
  // <p>Hello, World!</p>
}
```

If you need help, [email me](mailto:support@toby3d.me?subject=Telegraph). If you want to help me, [send a donation](https://toby3d.me/en/pay).