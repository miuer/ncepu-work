package engine

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/miuer/ncepu-work/db/miu-bookstore/conf"
	"github.com/miuer/ncepu-work/db/miu-bookstore/crawler/spider/model"
)

type biduSearchEngine struct {
	parseRule       string
	searchRule      string
	domain          string
	parseResultFunc func(searchResult *model.SearchResult)
}

func NewBiduSearchEngine(parseResultFunc func(searchResult *model.SearchResult)) *biduSearchEngine {
	return &biduSearchEngine{
		parseRule:       "#content_left h3 .t a",
		searchRule:      "intitle:%s 阅读 小说",
		domain:          "https://www.baidu.com/s?wd=%s&ie=utf-8&rn=15&vf_bl=1",
		parseResultFunc: parseResultFunc,
	}
}

func (bidu *biduSearchEngine) EngineRun(novelName string, group *sync.WaitGroup) {
	defer group.Done()

	searchKey := url.QueryEscape(fmt.Sprintf(bidu.searchRule, novelName))
	requestURL := fmt.Sprintf(bidu.domain, searchKey)

	log.Println(requestURL)

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnHTML(bidu.parseRule, func(element *colly.HTMLElement) {

		group.Add(1)
		log.Println("123456")
		go bidu.extractData(element, group)
	})

	err := c.Visit(requestURL)
	if err != nil {
		fmt.Println(err)
	}

}

func (bidu *biduSearchEngine) extractData(element *colly.HTMLElement, group *sync.WaitGroup) {
	defer group.Done()

	href := element.Attr("href")
	title := element.Text

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnResponse(func(resp *colly.Response) {
		realURL := resp.Request.URL.String()
		isContain := strings.Contains(realURL, "baidu")
		if isContain {
			return
		}

		host := resp.Request.URL.Host

		_, ok := conf.EngineConf.IgnoreDomain[host]
		if ok {
			return
		}

		result := &model.SearchResult{Title: title, Href: realURL, Host: host}

		bidu.parseResultFunc(result)

	})

	err := c.Visit(href)
	if err != nil {
		fmt.Println(err)
	}
}
