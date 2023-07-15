package fetch

import (
	"reflect"
	"sync"
	"time"
)

type timeStruct struct {
	Url      string
	Start    time.Time
	End      time.Time
	Duration float64
	Status   string
	Bytes    int64
}

type fetchSync struct {
	Urls     []timeStruct
	Start    time.Time
	End      time.Time
	Duration float64
	// Status   string
}

var wg sync.WaitGroup

func FetchUrlsGo(urls []string) []timeStruct {
	// var res []timeStruct
	// start := time.Now()
	ch := make(chan timeStruct)
	for _, url := range urls {
		wg.Add(1)
		go FetchCh(url, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	// for k := range ch {
	// 	// fmt.Printf(" Url: %s \n Start Time: %s \n End Time: %s \n Duration: %f \n Length: %d\n", k.Url, k.Start, k.End, k.Duration, k.Bytes)
	// 	// fmt.Println(" ****************** ")
	// 	res = append(res, k)
	// }
	res := ChanToSlice(ch).([]timeStruct)
	// fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	return res
}

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

func FetchUrlsSync(urls []string) fetchSync {
	var res fetchSync
	var urlsResp []timeStruct
	res.Start = time.Now()

	for _, url := range urls {
		iter := FetchSync(url)
		urlsResp = append(urlsResp, iter)
	}
	res.End = time.Now()
	res.Duration = time.Since(res.Start).Seconds()
	res.Urls = urlsResp
	return res
}

// Call from command line -- No longer implemented -- Only can use when running program from command line
// ./go-fetching https://godoc.org https://golang.org http://gopl.io https://www.reddit.com/r/learnprogramming/
