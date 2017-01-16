# GoLang bindings for the Telegraph API [![discord](https://discordapp.com/api/guilds/208605007744860163/widget.png)](https://discord.gg/fM4QqmA)
> This project is just to provide a wrapper around the API without any additional features.

[![License](https://img.shields.io/npm/l/express.svg?maxAge=2592000)](LICENSE.md)
[![Build Status](https://travis-ci.org/toby3d/telegraph.svg)](https://travis-ci.org/toby3d/telegraph)
[![GoDoc](https://godoc.org/github.com/toby3d/telegraph?status.svg)](https://godoc.org/github.com/toby3d/telegraph)
[![Go Report](https://goreportcard.com/badge/github.com/toby3d/telegraph)](https://goreportcard.com/report/github.com/toby3d/telegraph)
[![Patreon](https://img.shields.io/badge/support-patreon-E6461A.svg?maxAge=2592000)](https://www.patreon.com/toby3d)

All [methods](https://toby3d.github.io/telegraph/#available-methods) and [types](https://toby3d.github.io/telegraph/#available-types) available and this library (possibly) is ready for use in production. Yaay!

## Start using telegraph
Download and install it:  
`$ go get -u github.com/toby3d/telegraph`

Import it in your code:  
`import "github.com/toby3d/telegraph"`

## Example
This is an example of "quick start", which shows **how to create a new account** for future pages, as well as **creating a [first simple page](http://telegra.ph/My-super-awesome-page-12-25)** with text, picture, video and signature:
```go
package main

import (
    "github.com/toby3d/telegraph"
    "log"
)

// Example content. Be sure to wrap every media in a <figure> tag, okay? Be easy, bro.
const data = `<figure><img src="/file/6a5b15e7eb4d7329ca7af.jpg"/></figure><p><i>Hello</i>, my name is
<b>Page</b>, <u>look at me</u>!</p><figure><iframe src="https://youtu.be/fzQ6gRAEoy0"></iframe>
<figcaption>Yes, you can embed youtube, vimeo and twitter widgets too!</figcaption></figure>`

func main() {
    // Create new Telegraph account. Author name/link can be epmty.
    // So secure. Much anonymously. Wow.
    acc, err := telegraph.CreateAccount(
        "toby3d", // required for assign all new pages (invisible for others)
        "Maxim Lebedev",
        "https://telegram.me/toby3d",
    )
    if err != nil {
        log.Fatal(err.Error())
    }

    // Boom!.. And your text will be understandable for Telegraph. MAGIC.
    content, _ := telegraph.ContentFormat(data)
    
    newPage := &telegraph.Page{
        Title:   "My super-awesome page",
        Content: content,

        // Not necessarily, but, hey, it's just an example.
        AuthorName: acc.AuthorName,
        AuthorURL:  acc.AuthorURL,
    }

    if page, err := acc.CreatePage(newPage, false); err != nil {
        log.Print(err.Error())
    }
	
    log.Println("Kaboom! Page created, look what happened:", page.URL)
}
```

## Requirements
- [fasthttp](https://github.com/valyala/fasthttp)
- [net/html](https://golang.org/x/net/html)

## Documentation
- [Contributors](CONTRIBUTORS.md) and Patrons!
- [License](LICENSE.md)
- [ToDo](https://github.com/toby3d/telegraph/projects/1)
