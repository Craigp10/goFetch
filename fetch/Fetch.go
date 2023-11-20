package fetch

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sync"
	"time"
)

type TimeStruct struct {
	Url      string
	Start    time.Time
	End      time.Time
	Duration float64
	Status   Status
	Bytes    int64
}

type Syncd struct {
	Urls     []TimeStruct
	Start    time.Time
	End      time.Time
	Duration float64
	Status   Status
}

var wg sync.WaitGroup

type Status string

const (
	Failure Status = "failure"
	Success Status = "success"
)

// SyncCh manages the sync via channel
func SyncCh(url string, ch chan<- TimeStruct) {
	var t TimeStruct
	defer wg.Done()
	defer func() { ch <- t }()
	t.Start = time.Now()
	t.Url = url
	resp, err := GetUrl(url)
	if err != nil {
		fmt.Printf("Unable to fetch url [ %s ], -- %v,\n", url, err)
		t.Status = Failure
		return
	}
	t.End = time.Now()
	t.Duration = time.Since(t.Start).Seconds()
	t.Bytes, err = readResponseBody(resp)
	if err != nil {
		fmt.Printf("Unable to read response body for url [ %s ], -- %v,\n", url, err)
		t.Status = Failure
		return
	}
	t.Status = Success

	return
}

// Sync manages the 'sync' for a provided url
func Sync(url string) TimeStruct {
	var res TimeStruct
	res.Url = url
	res.Start = time.Now()
	resp, err := GetUrl(url)
	if err != nil {
		fmt.Printf("Unable to fetch url [ %s ], -- %v,\n", url, err)
		res.Status = Failure
		return res
	}
	res.End = time.Now()
	secs := time.Since(res.Start).Seconds()
	res.Duration = secs
	res.Bytes, err = readResponseBody(resp)
	if err != nil {
		res.Status = Failure
		return res
	}
	res.Status = Success
	return res

}

// fetchUrl performs a GET request on the provided url
func GetUrl(url string) (*http.Response, error) {
	return http.Get(url)
}

// readResponseBody assumes successful response argument and copies the response body into an int64 buffer
func readResponseBody(resp *http.Response) (int64, error) { // Maybe byte[] instead of int?
	var buf int64
	defer resp.Body.Close()
	buf, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

// GetUrlsGo fetches the provided urls and 'syncs' the response via go routines
func GetUrlsGo(urls []string) []TimeStruct {
	ch := make(chan TimeStruct)
	for _, url := range urls {
		wg.Add(1)
		go SyncCh(url, ch)
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

// GetUrlsSync fetches the provided urls and 'syncs' the responses
func GetUrlsSync(urls []string) Syncd {
	var res Syncd
	var urlsResp []TimeStruct
	res.Start = time.Now()

	for _, url := range urls {
		iter := Sync(url)
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
