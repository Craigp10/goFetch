# Welcome to goFetch

goFetch is an Obserability program aimed to provide insight and later provide analytics on network requests.

Currently in version 1 where the user can see the runtime of fetching provided URLs through both synchronous and asynchronous approaches.

goFetch was created to gain exposure by building an application with specific go packages and docker with the intent to eventually be scaled up to a network request observability plugin.

The continued development of goFetch will branch into improved containerization, deeper networking programming, along with concurrency on the web.

To run this application follow these steps:

1. Install go if not already installed [here](https://go.dev/doc/install)

2. clone this program `git clone https://github.com/Craigp10/goFetch.git`

3. Ensure go packages are installed and up to date - execute `go mod tidy`

4. Run locally. The program can be ran locally through a few approaches. by following either of the steps below:
   1. Go server
      a. execute in terminal `go run server.go`
   2. Docker container
      a. Ensure docker is installed locally [here](https://docs.docker.com/engine/install/)
      b. Build Docker image - execute `docker build --tag go-fetch .`
      c. Run docker container with image - execute `docker run -p 8080:8080 -d --name go_fetch go-fetch`
      d. Ensure container named `go_fetch` is running - execute `docker ps`
   3. Running binary
      a. Build the binary `go build .`
      b. Move the binary to your PATH, most likely in the following path `/Users/<userprofile>/bin/`
      c. Execute the following to run the server on port 8080 `go-fetch run --port 8080`

5. Request server by following the examples below...

1. 
## API Endpoints

Path: `http://localhost:8080/fetch/attempts`
Verb: `POST`
Body:

```
{
	"urls": [
	"https://godoc.org",
	"https://golang.org",
	"http://gopl.io",
	"https://google.com",
	"https://www.google.com/search?q=golang",
	"https://github.com/Craigp10/goFetch",
	"https://www.reddit.com/r/learnprogramming"
	],
	"attempts":10
}
```

Response:

```
{
	"TotalTime": 38.939043247,
	"ChAvg": 1.1763240082,
	"SyncAvg": 1.8566166617,
	"MutexAvg": 0.8609441541999999
}
```

Path: `http://localhost:8080/fetch`
Verb: `POST`
Body:

```
{
	"urls":[
		"https://godoc.org",
		"https://golang.org",
		"http://gopl.io",
		"https://google.com",
		"https://www.google.com/search?q=golang",
		"https://github.com/Craigp10/goFetch",
		"https://www.reddit.com/r/learnprogramming"
	]
}
```

Response:

```
{
	"ChUrls": [
		{
			"Url": "https://github.com/Craigp10/goFetch",
			"Start": "2024-07-07T17:51:26.43687-04:00",
			"End": "2024-07-07T17:51:26.462207-04:00",
			"Duration": 0.025337902,
			"Status": "success",
			"Bytes": 332567
		},
		{
			"Url": "https://godoc.org",
			"Start": "2024-07-07T17:51:26.436779-04:00",
			"End": "2024-07-07T17:51:26.554127-04:00",
			"Duration": 0.117348435,
			"Status": "success",
			"Bytes": 32378
		},
		{
			"Url": "https://golang.org",
			"Start": "2024-07-07T17:51:26.43661-04:00",
			"End": "2024-07-07T17:51:26.554501-04:00",
			"Duration": 0.11789157,
			"Status": "success",
			"Bytes": 61860
		},
		{
			"Url": "https://google.com",
			"Start": "2024-07-07T17:51:26.436843-04:00",
			"End": "2024-07-07T17:51:26.581429-04:00",
			"Duration": 0.144585927,
			"Status": "success",
			"Bytes": 20317
		},
		{
			"Url": "http://gopl.io",
			"Start": "2024-07-07T17:51:26.436595-04:00",
			"End": "2024-07-07T17:51:26.591182-04:00",
			"Duration": 0.154587354,
			"Status": "success",
			"Bytes": 4154
		},
		{
			"Url": "https://www.google.com/search?q=golang",
			"Start": "2024-07-07T17:51:26.436946-04:00",
			"End": "2024-07-07T17:51:26.799753-04:00",
			"Duration": 0.362807295,
			"Status": "success",
			"Bytes": 120626
		},
		{
			"Url": "https://www.reddit.com/r/learnprogramming",
			"Start": "2024-07-07T17:51:26.436938-04:00",
			"End": "2024-07-07T17:51:26.997144-04:00",
			"Duration": 0.560205162,
			"Status": "success",
			"Bytes": 538060
		}
	],
	"ChTimed": 0.56066923,
	"SyncTimed": 1.588066613,
	"MutexTimed": 0.719711609
}

```

Path: `http://localhost:8080/fetch/validate`
Verb: `POST`
Body:

```
{
	"urls": [
		"https://godoc.org",
		"https://golang.org",
		"http://gopl.io",
		"https://google.com",
		"https://www.google.com/search?q=golang",
		"https://github.com/Craigp10/goFetch",
		"https://www.reddit.com/r/learnprogramming",
		"http://meemeals.org"
	]
}
```

Response:

```
{
	"Valid": false,
	"InValidUrls": [
		"http://meemeals.org"
	]
}
```
