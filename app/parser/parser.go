package parser

import (
	"net/http"
	"github.com/vasudevrani/go-sitemap/domain/model"
)

type Parser interface {
	GetSeoData(resp *http.Response) (model.SeoData, error)
}