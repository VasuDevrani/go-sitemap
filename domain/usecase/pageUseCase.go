package usecase

import (
	"net/http"

	ps "github.com/vasudevrani/sitemap/app/parser"
	"github.com/vasudevrani/sitemap/app/api"
	"github.com/vasudevrani/sitemap/domain/model"
)

func ScrapePage(url string, token chan struct{}, parser ps.Parser) (model.SeoData, error) {
	res, err := crawlPage(url, token)
	if err != nil {
		return model.SeoData{}, err
	}
	data, err := parser.GetSeoData(res)
	if err != nil {
		return model.SeoData{}, err
	}
	return data, nil
}

func crawlPage(url string, tokens chan struct{}) (*http.Response, error) {
	tokens <- struct{}{}
	resp, err := api.MakeRequest(url)
	<-tokens
	if err != nil {
		return nil, err
	}
	return resp, err
}