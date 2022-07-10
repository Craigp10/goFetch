package main

import (
	"encoding/json"
	"log"
	"net/http"
	"practice-gin/fetch"
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
	SingleThreadTime string        `json:"singleThreadTime"`
	Stats            []timeRequest `json:"stats"`
}

type urls struct {
	Urls []string `json:"urls"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go and Gin!")
	})

	r.POST("/fetchUrls", func(c *gin.Context) {
		start := time.Now()
		var urls urls

		if err := c.BindJSON(&urls); err != nil {
			log.Fatal("ERROR", err)
		}

		res := fetch.FetchUrls(urls.Urls)
		elpsdTime := time.Now().Sub(start).Seconds()
		data, err := json.Marshal(res)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"testing": true, "urls": string(data), "ChTimed": elpsdTime})

	})
	r.Run()
}
