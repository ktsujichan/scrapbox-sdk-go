# Scrapbox SDK for Go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/ktsujichan/scrapbox-sdk-go/issues)

Scrapbox API client library written in Golang.

## Install
```
go get -u github.com/ktsujichan/scrapbox-sdk-go
```

## How to Use
```golang
package main

import (
	"context"
	"fmt"

	"github.com/ktsujichan/scrapbox-sdk-go"
)

func main() {
	c, _ := scrapbox.NewClient()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pages, _ := c.ListPages(ctx, "projectName", nil)
	fmt.Println(pages)
}
```
