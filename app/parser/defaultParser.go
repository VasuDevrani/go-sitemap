package parser

import (
	"net/http"
	"github.com/vasudevrani/sitemap/domain/model"

	"github.com/PuerkitoBio/goquery"
)

type DefaultParser struct {}

func (d DefaultParser) GetSeoData(resp *http.Response) (model.SeoData, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return model.SeoData{}, err
	}
	result := model.SeoData{}
	result.URL = resp.Request.URL.String()
	result.StatusCode = resp.StatusCode
	result.Title = doc.Find("title").First().Text()
	result.H1 = doc.Find("h1").First().Text()
	result.MetaDescription, _ = doc.Find("meta[name^=description]").Attr("content")
	return result, nil
}