package fetch

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Fetch(url string, ch chan<- timeStruct) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	end := time.Now()
	var t timeStruct
	t.Url = url
	t.Start = start
	t.End = end
	secs := time.Since(start).Seconds()
	t.Duration = secs
	if err != nil {
		t.Status = "failure"
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Status = "failure"
		return
	}
	t.Bytes = nbytes
	ch <- t
	return
}
