package sitemap

type Sitemap interface {
	NewSiteMap(URL string, c int) interface{}
}