package fetch

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// FetchSyncCh manages the sync via channel
func FetchSyncCh(url string, ch chan<- TimeStruct) {
	var t TimeStruct
	defer wg.Done()
	defer func() { ch <- t }()
	t.Start = time.Now()
	t.Url = url
	resp, err := fetchUrl(url)
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

// FetchSync manages the 'sync' for a provided url
func FetchSync(url string) TimeStruct {
	var res TimeStruct
	res.Url = url
	res.Start = time.Now()
	resp, err := fetchUrl(url)
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
func fetchUrl(url string) (*http.Response, error) {
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
