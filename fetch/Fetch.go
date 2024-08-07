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
	m        sync.Mutex
	Bytes    int64
}

func (s *Syncd) SetMutex(ts TimeStruct) {
	s.m.Lock()
	defer s.m.Unlock()

	s.Bytes += ts.Bytes
	s.Urls = append(s.Urls, ts)
}

type Status string

const (
	Failure Status = "failure"
	Success Status = "success"
)

// SyncCh manages the sync via channel
func SyncCh(url string, ch chan<- TimeStruct, wg *sync.WaitGroup) {
	defer wg.Done()

	var t TimeStruct
	t.Start = time.Now()
	t.Url = url

	resp, err := GetUrl(url)
	if err != nil {
		fmt.Printf("Error fetching URL [%s]: %v\n", url, err)
		t.Status = Failure
		ch <- t

		return
	}

	defer resp.Body.Close()

	t.Bytes, err = readResponseBody(resp)
	if err != nil {
		fmt.Printf("Error reading response body for URL [%s]: %v\n", url, err)
		t.Status = Failure
		ch <- t

		return
	}

	t.End = time.Now()
	t.Duration = time.Since(t.Start).Seconds()
	t.Status = Success
	ch <- t
}

func SyncMutex(url string, synced *Syncd, wg *sync.WaitGroup) {
	var t TimeStruct
	status := Failure
	defer func() {
		synced.SetMutex(t)

		t.Status = status
		t.Duration = time.Since(t.Start).Seconds()
		t.End = time.Now()
		wg.Done()
	}()

	t.Start = time.Now()
	t.Url = url

	resp, err := GetUrl(url)
	if err != nil {
		fmt.Printf("Unable to fetch url [ %s ], -- %v,\n", url, err)
		return
	}

	t.Bytes, err = readResponseBody(resp)
	if err != nil {
		fmt.Printf("Unable to read response body for url [ %s ], -- %v,\n", url, err)
		return
	}

	status = Success
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
	seconds := time.Since(res.Start).Seconds()
	res.Duration = seconds
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
func GetUrlsGoChan(urls []string) []TimeStruct {
	ch := make(chan TimeStruct)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go SyncCh(url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// res := chanToSlice(ch).([]TimeStruct)

	var res []TimeStruct
	for t := range ch {
		res = append(res, t)
	}

	return res
}

// GetUrlsGo fetches the provided urls and 'syncs' the response via go routines
func GetUrlsGoMutex(urls []string) *Syncd {
	s := &Syncd{}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go SyncMutex(url, s, &wg)
	}

	wg.Wait()

	return s
}

// chanToSlice converts a channel to a slice of any type. Not in use at the moment.
func chanToSlice(ch interface{}) interface{} {
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
func GetUrlsSync(urls []string) *Syncd {
	res := &Syncd{}
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
