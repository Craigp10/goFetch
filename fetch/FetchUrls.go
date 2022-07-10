package fetch

import (
	"fmt"
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

var wg sync.WaitGroup

func FetchUrls(urls []string) []timeStruct {
	var res []timeStruct
	start := time.Now()
	ch := make(chan timeStruct)
	for _, url := range urls {
		wg.Add(1)
		go Fetch(url, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for k := range ch {
		fmt.Printf(" Url: %s \n Start Time: %s \n End Time: %s \n Duration: %f \n Length: %d\n", k.Url, k.Start, k.End, k.Duration, k.Bytes)
		fmt.Println(" ****************** ")
		res = append(res, k)
	}

	// for range os.Args[1:] {
	// 	fmt.Println(<-ch)
	// }
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	return res
}

// func fetchAll() {

// }

// ./go-fetching https://godoc.org https://golang.org http://gopl.io https://www.reddit.com/r/learnprogramming/
