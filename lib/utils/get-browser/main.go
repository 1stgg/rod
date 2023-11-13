// Package main ...
package main

import (
	"fmt"

	"github.com/1stgg/rod/lib/launcher"
	"github.com/1stgg/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
