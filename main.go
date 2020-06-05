package main

import (
	"github.com/Z-M-Huang/RealEstateScraper/db"
	"github.com/Z-M-Huang/RealEstateScraper/scraper"
	"github.com/Z-M-Huang/RealEstateScraper/utils"
)

func main() {
	db.InitDB()
	utils.InitRedis()
	trulia := &scraper.Trulia{}
	trulia.Start()
}
