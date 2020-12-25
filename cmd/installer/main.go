package main

import (
	"github.com/mxschmitt/playwright-go"
	_ "github.com/starptech/fay/docs"
)

func main() {
	playwright.Install(&playwright.RunOptions{})
}
