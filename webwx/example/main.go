package main

import (
	"fmt"
	"github.com/wuxc/gowebwx/webwx"
)

func main() {
	w := webwx.NewWebWx()
	w.OnOnline(func() {
		fmt.Println("Online!")
	})
	w.OnOffline(func() {
		fmt.Println("Offline!")
	})
	w.Start(0)
}
