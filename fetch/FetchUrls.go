package fetch

import (
	"reflect"
	"sync"
	"time"
)

type TimeStruct struct {
	Url      string
	Start    time.Time
	End      time.Time
	Duration float64
	Status   string
	Bytes    int64
}

type FetchSyncd struct {
	Urls     []TimeStruct
	Start    time.Time
	End      time.Time
	Duration float64
	Status   string
}

var wg sync.WaitGroup

const (
	Failure string = "failure"
	Success        = "success"
)

// FetchUrlsGo fetches the provided urls and 'syncs' the response via go routines
func FetchUrlsGo(urls []string) []TimeStruct {
	ch := make(chan TimeStruct)
	for _, url := range urls {
		wg.Add(1)
		go FetchSyncCh(url, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	res := ChanToSlice(ch).([]TimeStruct)
	// fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	return res
}

// ChanToSlice converts a channel to a slice of any type
func ChanToSlice(ch interface{}) interface{} {
	chv := reflect.ValueOf(ch)
	slv := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(ch).Elem()), 0, 0)
	for {
		v, ok := chv.Recv()
		if !ok {
			return slv.Interface()
		}
		slv = reflect.Append(slv, v)
	}
}

// FetchUrlsSync fetches the provided urls and 'syncs' the responses
func FetchUrlsSync(urls []string) FetchSyncd {
	var res FetchSyncd
	var urlsResp []TimeStruct
	res.Start = time.Now()

	for _, url := range urls {
		iter := FetchSync(url)
		urlsResp = append(urlsResp, iter)
	}
	res.End = time.Now()
	res.Duration = time.Since(res.Start).Seconds()
	res.Urls = urlsResp
	res.Status = Success
	return res
}

// Call from command line -- No longer implemented -- Only can use when running program from command line
// ./go-fetching https://godoc.org https://golang.org http://gopl.io https://www.reddit.com/r/learnprogramming/
