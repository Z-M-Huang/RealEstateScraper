package scraper

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Z-M-Huang/RealEstateScraper/db"
	"github.com/Z-M-Huang/RealEstateScraper/utils"
	"github.com/Z-M-Huang/sitemap-parser"
	"github.com/gocolly/colly/v2"
)

//Trulia https://www.trulia.com/
type Trulia struct {
	elements   []sitemap.Element
	gzSitemaps []string
}

var (
	locker sync.Mutex
)

//Start process
func (t *Trulia) Start() {
	c := colly.NewCollector(
		colly.UserAgent("Real Estate Agent"),
	)
	c.OnHTML("script", func(e *colly.HTMLElement) {
		if e.Attr("id") == "__NEXT_DATA__" {
			appData := &truliaRawObject{}
			err := json.Unmarshal([]byte(e.Text), &appData)
			if err != nil {
				utils.Logger.Error(err.Error())
				return
			}

			// key := t.getRedisKey(appData.Props.HomeDetails.Location.StateCode, appData.Props.HomeDetails.Location.City, appData.Props.AsPath)
			var b bytes.Buffer
			gz := gzip.NewWriter(&b)
			jsonBytes, err := json.Marshal(appData.Props)
			if err != nil {
				utils.Logger.Error(err.Error())
				return
			}
			_, err = gz.Write(jsonBytes)
			if err != nil {
				utils.Logger.Error(err.Error())
				return
			}
			gz.Close()
			truliaEntity := &db.Trulia{
				URL:   appData.Props.AsPath,
				State: strings.ToUpper(appData.Props.HomeDetails.Location.StateCode),
				City:  strings.ToUpper(appData.Props.HomeDetails.Location.City),
				Data:  b.Bytes(),
			}
			err = truliaEntity.Save()
			if err != nil {
				utils.Logger.Error(err.Error())
			}
			// utils.RedisSet(key, b.Bytes(), 24*time.Hour)
			link := t.getNextLink()
			if link != "" {
				c.Visit(link)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		utils.Logger.Sugar().Infof("visiting %s", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		link := t.getNextLink()
		if link != "" {
			c.Visit(link)
		}
	})

	c.Visit(t.getNextLink())
}

func (Trulia) getRedisTimeUpdatedKey(url string) string {
	return fmt.Sprintf("TRULIA_UPDATE_%s", url)
}

func (Trulia) getRedisKey(state, city, url string) string {
	state = strings.TrimSpace(strings.ToUpper(state))
	city = strings.TrimSpace(strings.ToUpper(city))
	return fmt.Sprintf("TRULIA_STATE_CITY_%s_%s_%s", state, city, url)
}

func (t *Trulia) getNextLink() string {
	locker.Lock()

	var e sitemap.Element

	for {
		if len(t.elements) > 0 {
			e = t.elements[0]
			t.elements = t.elements[1:]
			utils.Logger.Sugar().Infof("There are %d urls left in current sitemap", len(t.elements))
			break
		}

		if len(t.gzSitemaps) > 0 {
			gzSM := t.gzSitemaps[0]
			utils.Logger.Sugar().Info("Start loading urls from sitemap ", gzSM)
			sm, err := sitemap.GetSitemapGZ(gzSM)
			utils.Logger.Sugar().Infof("%d urls loaded from sitemap", len(sm.Elements))
			if err != nil {
				utils.Logger.Error(err.Error())

				locker.Unlock()
				return ""
			}
			t.gzSitemaps = t.gzSitemaps[1:]
			for _, r := range sm.Elements {
				t.elements = append(t.elements, r)
			}
			e = t.elements[0]
			t.elements = t.elements[1:]
			break
		}
		t.getIndexes()
	}

	relativePath := strings.Replace(e.Loc, "https://www.trulia.com", "", 1)
	key := t.getRedisTimeUpdatedKey(relativePath)
	if utils.RedisExist(key) {
		timeStr, err := utils.RedisGetString(key)
		if err != nil {
			utils.Logger.Error(err.Error())
		} else {
			redisTime, _ := time.Parse(time.RFC3339, timeStr)
			mapTime, _ := time.Parse(time.RFC3339, e.LastMod)
			if mapTime.After(redisTime) {
				utils.RedisSet(key, e.LastMod, 24*time.Hour)
			} else {
				locker.Unlock()
				return t.getNextLink()
			}
		}
	} else {
		utils.RedisSet(key, e.LastMod, 24*time.Hour)
	}

	locker.Unlock()
	return e.Loc
}

func (t *Trulia) getIndexes() error {
	//public records
	utils.Logger.Info("Getting indexes")
	prIndex, err := sitemap.GetIndex("https://www.trulia.com/sitemaps/xml/public_records/index.xml")
	if err != nil {
		utils.Logger.Error(err.Error())
		return err
	}
	for _, r := range prIndex.Elements {
		t.gzSitemaps = append(t.gzSitemaps, r.Loc)
	}

	//recent updates
	ruIndex, err := sitemap.GetIndex("https://www.trulia.com/sitemaps/xml/public_records/index.xml")
	if err != nil {
		utils.Logger.Error(err.Error())
		return err
	}
	utils.Logger.Info("Loading sitemaps...")
	for _, r := range ruIndex.Elements {
		t.gzSitemaps = append(t.gzSitemaps, r.Loc)
	}
	return nil
}
