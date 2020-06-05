package main

import (
	"github.com/Z-M-Huang/RealEstateScraper/scraper"
	"github.com/Z-M-Huang/RealEstateScraper/utils"
)

func main() {
	utils.InitRedis()
	trulia := &scraper.Trulia{}
	trulia.Start()
}
