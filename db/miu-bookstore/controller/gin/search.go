package gin

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miuer/ncepu-work/db/miu-bookstore/conf"

	se "github.com/miuer/ncepu-work/db/miu-bookstore/crawler/spider/engine"
	"github.com/miuer/ncepu-work/db/miu-bookstore/crawler/spider/model"
)

func novelSearch(ctx *gin.Context) {
	wd, exist := ctx.GetQuery("wd")
	if !exist || len(wd) == 0 {
		// ---------------------------
		return
	}

	start := time.Now().UnixNano()
	result := startSearch(wd)
	end := time.Now().UnixNano()

	searchTime := end - start
	searchCount := len(result)

	ctx.HTML(http.StatusOK, "search_index.html", gin.H{
		"list":        result,
		"novelName":   wd,
		"elapsedTime": searchTime,
		"count":       searchCount,
		"head":        "search_head",
	})
}

func startSearch(novelName string) []*model.SearchResult {
	// --------------------------check  redis

	group := sync.WaitGroup{}
	results := make([]*model.SearchResult, 0)

	for _, engine := range conf.EngineConf.Engine {
		var searchEngine se.EngineRuner
		group.Add(1)
		switch engine {
		case "bidu":
			searchEngine = se.NewBiduSearchEngine(func(result *model.SearchResult) {
				results = append(results, result)
			})
		}

		if searchEngine != nil {
			go searchEngine.EngineRun(novelName, &group)
		}
	}

	group.Wait()

	// ---------------- store in redis
	return results

}
