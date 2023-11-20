package main

import (
	"encoding/json"
	"fmt"
	"go-fetch/fetch"
	"go-fetch/utils"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type fetchUrlsRequest struct {
	Urls []string `json:"urls"`
}

type fetchUrlsResponse struct {
	ChUrls    []fetch.TimeStruct `json:"ChUrls"`
	ChTimed   float64
	SyncTimed float64
	SyncUrls  fetch.Syncd `json:"SyncUrls"`
}

type fetchUrlsAttemptsRequest struct {
	Urls     []string `json:"urls"`
	Attempts int      `json:"attempts"`
}

type fetchUrlsAttemptsResponse struct {
	TotalTime float64
	ChAvg     float64
	SyncAvg   float64
}

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

type fetchValidateRequest struct {
	Urls []string
}

type fetchValidateResponse struct {
	Valid       bool
	InValidUrls []string
}

type urls struct {
	Urls []string `json:"urls"`
}

type attemptTimes struct {
	chTime   float64
	syncTime float64
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/fetch", fetchUrls).Methods("POST")
	router.HandleFunc("/fetch/attempts", fetchUrlsAttempts).Methods("POST")
	router.HandleFunc("/fetch/validate", fetchUrlsValidate).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func fetchUrlsValidate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading request body: %v", err)
		return
	}

	f := fetchValidateRequest{}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &f); err != nil {
		log.Fatal("Error binding urls", err)
	}
	res := fetchValidateResponse{
		Valid: true,
	}
	for _, url := range f.Urls {
		resp, err := fetch.GetUrl(url)
		if err != nil || resp == nil {
			res.Valid = false
			res.InValidUrls = append(res.InValidUrls, url)
		}
	}

	writeResponse(w, res)
}

// fetchUrls fetches the provided urls
func fetchUrls(w http.ResponseWriter, r *http.Request) {
	var urls urls

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading request body: %v", err)
		return
	}

	// Close the request body after reading (important!)
	defer r.Body.Close()

	if err := json.Unmarshal(body, &urls); err != nil {
		log.Fatal("Error binding urls", err)
	}
	fmt.Printf("fetched urls: [%d]\n", len(urls.Urls))

	if len(urls.Urls) == 0 {
		w.WriteHeader(400)
		m := generateException("No urls provided in request")
		w.Write(m)
		return
	}

	startCh := time.Now()
	chTimedUrls := fetch.GetUrlsGo(urls.Urls)
	elpsdTimeCh := time.Now().Sub(startCh).Seconds()
	// chTimedData, err := json.Marshal(chTimedUrls)
	startSync := time.Now()
	syncTimedUrls := fetch.GetUrlsSync(urls.Urls)
	elpsdTimeSync := time.Now().Sub(startSync).Seconds()
	// syncTimedData, err := json.Marshal(syncTimedUrls)

	if err != nil {
		panic(err)
	}

	f := fetchUrlsResponse{
		ChUrls:    chTimedUrls,
		ChTimed:   elpsdTimeCh,
		SyncTimed: elpsdTimeSync,
		SyncUrls:  syncTimedUrls,
	}

	writeResponse(w, f)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from goFetch \n")
}

func fetchUrlsAttempts(w http.ResponseWriter, r *http.Request) {
	// Definitely shouldn't parse in endpoint handler

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading request body: %v", err)
		return
	}
	// Close the request body after reading (important!)
	defer r.Body.Close()

	var start = time.Now()
	f := fetchUrlsAttemptsRequest{}

	if err := json.Unmarshal(body, &f); err != nil {
		log.Fatal("Error binding urls", err)
	}

	fmt.Printf("Running %d iterations on %s", f.Attempts, f.Urls)

	if len(f.Urls) == 0 {
		w.WriteHeader(400)
		m := generateException("No urls provided in request.")
		w.Write(m)
		return
	}

	if f.Attempts < 1 {
		w.WriteHeader(400)
		m := generateException("Attempts must be greater than 0.")
		w.Write(m)
		return
	}

	var attemptsTimed []attemptTimes

	for i := 0; i < f.Attempts; i++ {
		startCh := time.Now()
		fetch.GetUrlsGo(f.Urls)
		elpsdTimeCh := time.Now().Sub(startCh).Seconds()
		startSync := time.Now()
		fetch.GetUrlsSync(f.Urls)
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

	a := fetchUrlsAttemptsResponse{
		TotalTime: duration,
		ChAvg:     chAvg,
		SyncAvg:   syncAvg,
	}

	writeResponse(w, a)
}

func writeResponse(w http.ResponseWriter, content interface{}) {
	j, err := json.Marshal(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshaling data: %v", err)
		return
	}

	w.Write(j)
}

type Exception struct {
	Message string
	Code    int
}

func generateException(message string) []byte {
	e := Exception{
		Message: message,
		Code:    400,
	}
	b, err := json.Marshal(e)
	if err != nil {
		panic("Error generating exception")
	}
	return b
}
