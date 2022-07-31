package fetch

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func FetchCh(url string, ch chan<- timeStruct) {
	var t timeStruct
	defer wg.Done()
	t.Start = time.Now()
	resp, err := http.Get(url)
	t.End = time.Now()
	t.Url = url
	t.Status = "success"
	t.Duration = time.Since(t.Start).Seconds()
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

func FetchSync(url string) timeStruct {
	var res timeStruct
	res.Url = url
	res.Start = time.Now()
	resp, err := http.Get(url)
	res.End = time.Now()

	secs := time.Since(res.Start).Seconds()
	res.Duration = secs
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		res.Status = "failure"
	}
	res.Bytes = nbytes

	return res

}
