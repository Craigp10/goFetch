package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"practice-gin/fetch"
	"practice-gin/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type timeRequest struct {
	url        string
	start_time time.Time
	end_time   time.Time
	secs       float64
	run_time   time.Duration
	status     string
	bytes      int64
}

type fetchAllResp struct {
	ChTimed          string        `json:"chTimed"`
	SyncTimed        string        `json:"syncTimed"`
	SingleThreadTime string        `json:"singleThreadTime"`
	Stats            []timeRequest `json:"stats"`
}

type urls struct {
	Urls []string `json:"urls"`
}

type attemptTimes struct {
	chTime   float64
	syncTime float64
}

type Body struct {
	// json tag to de-serialize json body
	Urls     []string `json:"urls"`
	Attempts int      `json:"attempts"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go and Gin!")
	})

	r.POST("/fetchUrls", func(c *gin.Context) {
		var urls urls

		if err := c.BindJSON(&urls); err != nil {
			log.Fatal("Error binding urls", err)
		}

		startCh := time.Now()
		chTimedUrls := fetch.FetchUrlsGo(urls.Urls)
		elpsdTimeCh := time.Now().Sub(startCh).Seconds()
		chTimedData, err := json.Marshal(chTimedUrls)
		startSync := time.Now()
		syncTimedUrls := fetch.FetchUrlsSync(urls.Urls)
		elpsdTimeSync := time.Now().Sub(startSync).Seconds()
		syncTimedData, err := json.Marshal(syncTimedUrls)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"chUrls": string(chTimedData), "ChTimed": elpsdTimeCh, "SyncTimed": elpsdTimeSync, "syncUrls": string(syncTimedData)})

	})

	r.POST("/fetchUrlsAttempts", func(c *gin.Context) {

		var body Body
		var urls []string
		var attempts int
		var start = time.Now()
		if err := c.BindJSON(&body); err != nil {
			log.Fatal("Error binding body", err)
		}

		attempts = body.Attempts
		urls = body.Urls

		fmt.Printf("Running %d iterations on %s", attempts, urls)

		var attemptsTimed []attemptTimes

		for i := 0; i < attempts; i++ {
			startCh := time.Now()
			fetch.FetchUrlsGo(urls)
			elpsdTimeCh := time.Now().Sub(startCh).Seconds()
			startSync := time.Now()
			fetch.FetchUrlsSync(urls)
			elpsdTimeSync := time.Now().Sub(startSync).Seconds()
			attemptsTimed = append(attemptsTimed, attemptTimes{
				chTime:   elpsdTimeCh,
				syncTime: elpsdTimeSync,
			})
		}

		var chValues, syncValues []float64

		for j := 0; j < len(attemptsTimed); j++ {
			chValues = append(chValues, attemptsTimed[j].chTime)
			syncValues = append(syncValues, attemptsTimed[j].syncTime)
		}

		chAvg := utils.Average(chValues)
		syncAvg := utils.Average(syncValues)
		duration := time.Since(start).Seconds()
		c.JSON(http.StatusOK, gin.H{"totalTime": duration, "chAvg": chAvg, "syncAvg": syncAvg})
	})
	r.Run()
}
