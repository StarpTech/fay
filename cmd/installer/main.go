package main

import (
	"github.com/mxschmitt/playwright-go"
	_ "github.com/starptech/fay/docs"
)

func main() {
	_ = playwright.Install(&playwright.RunOptions{})
}
