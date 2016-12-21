# GoLang bindings for the Telegraph API
> This project is just to provide a wrapper around the API without any additional features.

[![License](https://img.shields.io/npm/l/express.svg?maxAge=2592000)](LICENSE.md)
[![Build Status](https://travis-ci.org/toby3d/telegraph.svg)](https://travis-ci.org/toby3d/telegraph)
[![GoDoc](https://godoc.org/github.com/toby3d/telegraph?status.svg)](https://godoc.org/github.com/toby3d/telegraph)
[![Go Report](https://goreportcard.com/badge/github.com/toby3d/telegraph)](https://goreportcard.com/report/github.com/toby3d/telegraph)
[![Patreon](https://img.shields.io/badge/support-patreon-E66500.svg?maxAge=2592000)](https://www.patreon.com/toby3d)
[![discord](https://discordapp.com/api/guilds/208605007744860163/widget.png)](https://discord.gg/fM4QqmA)

## Available features
### [Methods](http://telegra.ph/api#Available-methods)
- [x] createAccount
- [ ] createPage
- [x] editAccountInfo
- [ ] editPage
- [x] getAccountInfo
- [x] getPage
- [x] getPageList
- [x] getViews
- [x] revokeAccessToken

### [Types](http://telegra.ph/api#Available-types)
- [x] Account
- [ ] Node
- [ ] NodeElement
- [x] Page
- [x] PageList
- [x] PageViews

## Start using telegraph
Download and install it:  
`$ go get -u github.com/toby3d/telegraph`

Import it in your code:  
`import "github.com/toby3d/telegraph"`

## Requirements
- [fasthttp](https://github.com/valyala/fasthttp)
- [net/html](https://golang.org/x/net/html)

## Documentation
- [Contributors](CONTRIBUTORS.md) and Patrons!
- [License](LICENSE.md)
- [ToDo](https://github.com/toby3d/telegraph/projects/1)