package sitemap

import (
	parser "github.com/vasudevrani/sitemap/app/parser"
	"github.com/vasudevrani/sitemap/domain/model"
	uc "github.com/vasudevrani/sitemap/domain/usecase"
)

type siteMap struct {
	URL string
	parser parser.Parser
	concurrency int
}

func NewSiteMap(URL string, c int) *siteMap {
	p := parser.DefaultParser{}
	return &siteMap{URL, p, c}
}

func (stm siteMap) ScrapeSitemap() []model.SeoData {
	results := uc.ExtractSitemapURLs(stm.URL)
	res := uc.ScrapeUrls(results, stm.parser, stm.concurrency)
	return res
}