package scraper

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/Z-M-Huang/RealEstateScraper/data"
	truliaData "github.com/Z-M-Huang/RealEstateScraper/scraper/data"
	"github.com/Z-M-Huang/RealEstateScraper/utils"
	"github.com/Z-M-Huang/sitemap-parser"
	"github.com/gocolly/colly/v2"
	"github.com/jinzhu/gorm"
	"github.com/temoto/robotstxt"
)

//Trulia https://www.trulia.com/
type Trulia struct {
	urls []string

	robots *robotstxt.RobotsData
}

var (
	locker sync.Mutex
)

const (
	robotsURL string = "https://www.trulia.com/robots.txt"
	agentName string = "Real Estate Agent"
)

//Scrape scraping process
func (t *Trulia) Scrape() {
	err := t.getRobotsTxt()
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	err = t.getURLsToScrap()
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	c := colly.NewCollector()
	c.OnHTML("script", func(e *colly.HTMLElement) {
		if e.Attr("id") == "__NEXT_DATA__" {
			appData := &truliaData.TruliaWebData{}
			err = json.Unmarshal([]byte(e.Text), &appData)
			if err != nil {
				utils.Logger.Error(err.Error())
				return
			}
			link := t.getNextLink()
			if link != "" {
				c.Visit(link)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		utils.Logger.Sugar().Infof("visiting %s", r.URL.String())
	})

	c.Visit(t.getNextLink())
}

func (t *Trulia) getURLsToScrap() error {

	//public records
	prIndex, err := sitemap.GetIndex("https://www.trulia.com/sitemaps/xml/public_records/index.xml")
	if err != nil {
		return err
	}

	prSitemaps, err := sitemap.GetSitemapGZ(prIndex.Elements[0].Loc)

	// prSitemaps, err := prIndex.GetSitemaps()
	// if err != nil {
	// 	return err
	// }

	// //recent updates
	// ruIndex, err := sitemap.GetIndex("https://www.trulia.com/sitemaps/xml/public_records/index.xml")
	// if err != nil {
	// 	return err
	// }
	// ruSitemaps, err := ruIndex.GetSitemaps()
	// if err != nil {
	// 	return err
	// }

	var allSavedTrulia []*data.TruliaEntity
	err = data.DoTransaction(func(tx *gorm.DB) error {
		if txDB := tx.Find(&allSavedTrulia); txDB.Error != nil {
			return txDB.Error
		}
		return nil
	})
	if err != nil {
		return err
	}

	//for _, sm := range prSitemaps {
	//for _, e := range sm.Elements {
	for _, e := range prSitemaps.Elements {
		found := false
		for _, record := range allSavedTrulia {
			if record.URL == e.Loc {
				found = true
				//check update time
				updateTime, err := time.Parse(time.RFC3339, e.LastMod)
				if err != nil {
					return err
				}
				if record.UpdatedAt.Before(updateTime) && t.robots.TestAgent(e.Loc, agentName) {
					t.urls = append(t.urls, e.Loc)
				}
				break
			}
		}
		if !found && t.robots.TestAgent(e.Loc, agentName) {
			t.urls = append(t.urls, e.Loc)
		}
	}
	//}

	// for _, sm := range ruSitemaps {
	// 	for _, e := range sm.Elements {
	// 		found := false
	// 		for _, record := range allSavedTrulia {
	// 			if record.URL == e.Loc {
	// 				found = true
	// 				//check update time
	// 				updateTime, err := time.Parse(time.RFC3339, e.LastMod)
	// 				if err != nil {
	// 					return err
	// 				}
	// 				if record.UpdatedAt.Before(updateTime) && t.robots.TestAgent(e.Loc, agentName) {
	// 					t.urls = append(t.urls, e.Loc)
	// 				}
	// 				break
	// 			}
	// 		}
	// 		if !found && t.robots.TestAgent(e.Loc, agentName) {
	// 			t.urls = append(t.urls, e.Loc)
	// 		}
	// 	}
	// }
	return nil
}

func (t *Trulia) getRobotsTxt() error {
	resp, err := http.Get(robotsURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	robots, err := robotstxt.FromResponse(resp)
	if err != nil {
		return err
	}
	t.robots = robots
	return nil
}

func (t *Trulia) getNextLink() string {
	locker.Lock()
	defer locker.Unlock()
	link := ""
	if len(t.urls) > 0 {
		link = t.urls[0]
		t.urls = t.urls[1:]
	}
	return link
}
