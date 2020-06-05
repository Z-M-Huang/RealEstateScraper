package scraper

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Z-M-Huang/RealEstateScraper/db"
	"github.com/Z-M-Huang/RealEstateScraper/utils"
	"github.com/Z-M-Huang/sitemap-parser"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/gocolly/colly/v2/queue"
)

//Trulia https://www.trulia.com/
type Trulia struct {
	elements   []sitemap.Element
	gzSitemaps []string

	q *queue.Queue
}

var (
	locker sync.Mutex
)

//Start process
func (t *Trulia) Start() {
	count := 0
	t.q, _ = queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)

	c := colly.NewCollector(
		colly.UserAgent("Real Estate Agent"),
		colly.AllowURLRevisit(),
	)
	rp, err := proxy.RoundRobinProxySwitcher(os.Getenv("TRULIA_PROXY"), os.Getenv("TRULIA_PROXY"), os.Getenv("TRULIA_PROXY"), os.Getenv("TRULIA_PROXY"), os.Getenv("TRULIA_PROXY"),
		os.Getenv("TRULIA_PROXY"), os.Getenv("TRULIA_PROXY"), os.Getenv("TRULIA_PROXY"), os.Getenv("TRULIA_PROXY"), os.Getenv("TRULIA_PROXY"))
	if err != nil {
		utils.Logger.Fatal(err.Error())
	}
	c.SetProxyFunc(rp)
	c.OnHTML("script", func(e *colly.HTMLElement) {
		if e.Attr("id") == "__NEXT_DATA__" {
			appData := &truliaRawObject{}
			err := json.Unmarshal([]byte(e.Text), &appData)
			if err != nil {
				utils.Logger.Error(err.Error())
				return
			}

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
				URL: appData.Props.AsPath,
			}
			if err = truliaEntity.Find(); err != nil {
				truliaEntity.State = strings.ToUpper(appData.Props.HomeDetails.Location.StateCode)
				truliaEntity.City = strings.ToUpper(appData.Props.HomeDetails.Location.City)
				truliaEntity.Data = b.Bytes()
			} else {
				truliaEntity.Data = b.Bytes()
			}
			err = truliaEntity.Save()
			if err != nil {
				utils.Logger.Error(err.Error())
			}
		}
	})
	c.OnRequest(func(r *colly.Request) {
		count++
		utils.Logger.Sugar().Infof("%d visiting %s", count, r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	for {
		fmt.Println(count)
		count = 0
		t.loadQueue()
		size, err := t.q.Size()
		if err != nil {
			return
		}
		if size > 0 {
			t.q.Run(c)
			c.Wait()
		} else {
			return
		}
	}
}

func (Trulia) getRedisTimeUpdatedKey(url string) string {
	return fmt.Sprintf("TRULIA_UPDATE_%s", url)
}

func (t *Trulia) loadQueue() {
	for {
		size, err := t.q.Size()
		if err != nil || size > 200 {
			return
		}
		var e sitemap.Element
		for {
			if len(t.elements) > 0 {
				e = t.elements[0]
				t.elements = t.elements[1:]
				break
			}

			if len(t.gzSitemaps) > 0 {
				gzSM := t.gzSitemaps[0]
				utils.Logger.Sugar().Info("Start loading urls from sitemap ", gzSM)
				sm, err := sitemap.GetSitemapGZ(gzSM)
				if err != nil {
					utils.Logger.Error(err.Error())
					return
				}
				utils.Logger.Sugar().Infof("%d urls loaded from sitemap", len(sm.Elements))
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
					continue
				}
			}
		} else {
			utils.RedisSet(key, e.LastMod, 24*time.Hour)
		}
		err = t.q.AddURL(e.Loc)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
	}
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
