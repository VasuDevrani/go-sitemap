# SITEMAP 

- Crawl sitemap of any website
- getSeoData of each sitemap URL
- non-blocking process with goroutines
- to generate sitemaps use [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator)

## Usage
``` go get github.com/vasudevrani/sitemap ```
```go
package main

import (
	"fmt"

	"github.com/vasudevrani/sitemap/app/sitemap"
)

func main() {
	stm := sitemap.NewSiteMap("sitemap_url", concurrency_number)
	results := stm.ScrapeSitemap()
	for _, res := range results {
		fmt.Println(res)
	}
}
```

## Example-View
<img width="400" alt="image" src="https://user-images.githubusercontent.com/101383635/234504479-30c06b87-c201-4737-8a06-9e614256cd85.png">
<img width="400" alt="image" src="https://user-images.githubusercontent.com/101383635/234504540-015315df-91bf-4f2d-9f51-dd9632bb9818.png">
