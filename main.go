package main

import (
	"github.com/Z-M-Huang/RealEstateScraper/data"
	"github.com/Z-M-Huang/RealEstateScraper/scraper"
)

func main() {
	data.InitDB()

	trulia := &scraper.Trulia{}
	trulia.Scrape()
}
