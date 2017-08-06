# GoLang bindings for the Telegraph API [![discord](https://discordapp.com/api/guilds/208605007744860163/widget.png)](https://discord.gg/QJ8z5BN)
> This project is just to provide a wrapper around the API without any additional features.

[![License](https://img.shields.io/npm/l/express.svg?maxAge=2592000)](LICENSE.md)
[![Build Status](https://travis-ci.org/toby3d/go-telegraph.svg)](https://travis-ci.org/toby3d/go-telegraph)
[![GoDoc](https://godoc.org/github.com/toby3d/go-telegraph?status.svg)](https://godoc.org/github.com/toby3d/go-telegraph)
[![Go Report](https://goreportcard.com/badge/github.com/toby3d/go-telegraph)](https://goreportcard.com/report/github.com/toby3d/go-telegraph)
[![Patreon](https://img.shields.io/badge/support-patreon-E6461A.svg?maxAge=2592000)](https://www.patreon.com/toby3d)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go)

All [methods](https://toby3d.github.io/go-telegraph/#available-methods) and [types](https://toby3d.github.io/go-telegraph/#available-types) available and this library (possibly) is ready for use in production. Yaay!

## Start using telegraph
Download and install it:  
`$ go get -u github.com/toby3d/go-telegraph`

Import it in your code:  
`import "github.com/toby3d/go-telegraph"`

## Example
This is an example of "quick start", which shows **how to create a new account** for future pages, as well as **creating a [first simple page](http://telegra.ph/My-super-awesome-page-12-25)** with text, picture, video and signature:
```go
package main

import (
    "log"

    telegraph "github.com/toby3d/go-telegraph"
)

// Example content. Be sure to wrap every media in a <figure> tag, okay?
// Be easy, bro.
const data = `
    <figure>
        <img src="/file/6a5b15e7eb4d7329ca7af.jpg"/>
    </figure>
    <p><i>Hello</i>, my name is <b>Page</b>, <u>look at me</u>!</p>
    <figure>
        <iframe src="https://youtu.be/fzQ6gRAEoy0"></iframe>
        <figcaption>
            Yes, you can embed youtube, vimeo and twitter widgets too!
        </figcaption>
    </figure>
`

func checkError(err error) {
    if err != nil {
        log.Fatalln(err.Error())
    }
}

func main() {
    // Create new Telegraph account. Author name/link can be epmty.
    // So secure. Much anonymously. Wow.
    newAccount := &telegraph.Account{
        ShortName:  "toby3d", // required
        AuthorName: "Maxim Lebedev",
        AuthorURL:  "https://t.me/toby3d",
    }
    acc, err := telegraph.CreateAccount(newAccount)
    checkError(err)

    // Boom!.. And your text will be understandable for Telegraph. MAGIC.
    content, err := telegraph.ContentFormat(data)
    checkError(err)

    newPage := &telegraph.Page{
        Title:   "My super-awesome page",
        Content: content,

        // Not necessarily, but, hey, it's just an example.
        AuthorName: acc.AuthorName,
        AuthorURL:  acc.AuthorURL,
    }

    page, err := acc.CreatePage(newPage, false)
    checkError(err)

    log.Println("Kaboom! Page created, look what happened:", page.URL)
}
```

## Need help?
- [Open new issue](https://github.com/toby3d/go-telegraph/issues/new)
- [Discuss in Discord](https://discord.gg/QJ8z5BN)