package parser

import (
	"net/http"
	"github.com/vasudevrani/sitemap/domain/model"
)

type Parser interface {
	GetSeoData(resp *http.Response) (model.SeoData, error)
}