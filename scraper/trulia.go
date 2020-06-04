package scraper

import (
	"github.com/Z-M-Huang/RealEstateScraper/utils"
	"github.com/yterajima/go-sitemap"
)

//Trulia https://www.trulia.com/
type Trulia struct {
	URLs []string
}

//Scrape scraping process
func (Trulia) Scrape() {
	index, err := sitemap.Get("https://www.trulia.com/sitemaps/xml/public_records/index.xml", nil)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
}
