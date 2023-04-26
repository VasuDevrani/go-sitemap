package main

import (
	"fmt"

	"github.com/vasudevrani/go-sitemap/app/sitemap"
)

func main() {
	stm := sitemap.NewSiteMap("https://www.quicksprout.com/sitemap.xml", 10)
	results := stm.ScrapeSitemap()
	for _, res := range results {
		fmt.Println(res)
	}
}