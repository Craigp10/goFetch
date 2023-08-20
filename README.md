# Welcome to goFetch

goFetch is a program created to gain some exposure building an application specific go packages and docker.

The continued development of goFetch will branch into imporved containerization, deeper networking programming and concurrency on the web.

The current functionality of the app is to fetch the provided urls utilizing concurrency vs iteration, and provide metrics to display the advantages of each approach.

To run this application follow these steps:

1. Install go if not already installed [here](https://go.dev/doc/install)

2. clone this program `git clone https://github.com/Craigp10/goFetch.git`

3. Ensure go packages are installed and up to date - execute `go mod tidy`

4. Run local either of options below:
   1. Local go server - execute in terminal `go run server.go`
   2. Local docker container
      a. Ensure docker is installed locally [here]()
      b. Build Docker image - execute `docker build --tag go-fetch .`
      c. Run docker container with image - execute `docker run -p 8080:8080 -d --name go_fetch go-fetch`
      d. Ensure container named `go_fetch` is running - execute `docker ps`
5. Make request to server. Examples below...

## API Endpoints

Path: `http://localhost:8080/fetch/attempts`
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
	],
	"attempts":5
}
```

Response:

```
{
	"ChUrls": [
		{
			"Url": "https://google.com",
			"Start": "2023-08-20T12:55:27.592806-07:00",
			"End": "2023-08-20T12:55:27.846943-07:00",
			"Duration": 0.254140169,
			"Status": "success",
			"Bytes": 19528
		},
		{
			"Url": "http://meemeals.org",
			"Start": "2023-08-20T12:55:27.59307-07:00",
			"End": "0001-01-01T00:00:00Z",
			"Duration": 0,
			"Status": "failure",
			"Bytes": 0
		},
		{
			"Url": "https://github.com/Craigp10/goFetch",
			"Start": "2023-08-20T12:55:27.592912-07:00",
			"End": "2023-08-20T12:55:28.000656-07:00",
			"Duration": 0.407748547,
			"Status": "success",
			"Bytes": 197987
		},
		{
			"Url": "https://godoc.org",
			"Start": "2023-08-20T12:55:27.592845-07:00",
			"End": "2023-08-20T12:55:28.226335-07:00",
			"Duration": 0.633497149,
			"Status": "success",
			"Bytes": 31436
		},
		{
			"Url": "https://golang.org",
			"Start": "2023-08-20T12:55:27.592939-07:00",
			"End": "2023-08-20T12:55:28.231908-07:00",
			"Duration": 0.638976328,
			"Status": "success",
			"Bytes": 61870
		},
		{
			"Url": "http://gopl.io",
			"Start": "2023-08-20T12:55:27.59277-07:00",
			"End": "2023-08-20T12:55:28.247664-07:00",
			"Duration": 0.654901017,
			"Status": "success",
			"Bytes": 4154
		},
		{
			"Url": "https://www.google.com/search?q=golang",
			"Start": "2023-08-20T12:55:27.592903-07:00",
			"End": "2023-08-20T12:55:27.846742-07:00",
			"Duration": 0.253842009,
			"Status": "success",
			"Bytes": 92851
		},
		{
			"Url": "http://meemeals.com",
			"Start": "2023-08-20T12:55:27.592929-07:00",
			"End": "2023-08-20T12:55:28.295401-07:00",
			"Duration": 0.702479953,
			"Status": "success",
			"Bytes": 2444
		},
		{
			"Url": "https://www.reddit.com/r/learnprogramming",
			"Start": "2023-08-20T12:55:27.592892-07:00",
			"End": "2023-08-20T12:55:28.007489-07:00",
			"Duration": 0.414601957,
			"Status": "success",
			"Bytes": 394728
		}
	],
	"ChTimed": 1.156248495,
	"SyncTimed": 2.120280619,
	"SyncUrls": {
		"Urls": [
			{
				"Url": "https://godoc.org",
				"Start": "2023-08-20T12:55:28.748964-07:00",
				"End": "2023-08-20T12:55:28.889364-07:00",
				"Duration": 0.140401837,
				"Status": "success",
				"Bytes": 31436
			},
			{
				"Url": "https://golang.org",
				"Start": "2023-08-20T12:55:28.896474-07:00",
				"End": "2023-08-20T12:55:29.204132-07:00",
				"Duration": 0.307661943,
				"Status": "success",
				"Bytes": 61870
			},
			{
				"Url": "http://gopl.io",
				"Start": "2023-08-20T12:55:29.206925-07:00",
				"End": "2023-08-20T12:55:29.330066-07:00",
				"Duration": 0.123143285,
				"Status": "success",
				"Bytes": 4154
			},
			{
				"Url": "https://google.com",
				"Start": "2023-08-20T12:55:29.330547-07:00",
				"End": "2023-08-20T12:55:29.453982-07:00",
				"Duration": 0.123436427,
				"Status": "success",
				"Bytes": 19462
			},
			{
				"Url": "https://www.google.com/search?q=golang",
				"Start": "2023-08-20T12:55:29.456928-07:00",
				"End": "2023-08-20T12:55:29.596989-07:00",
				"Duration": 0.140063008,
				"Status": "success",
				"Bytes": 92851
			},
			{
				"Url": "https://github.com/Craigp10/goFetch",
				"Start": "2023-08-20T12:55:29.886662-07:00",
				"End": "2023-08-20T12:55:29.933538-07:00",
				"Duration": 0.046877705,
				"Status": "success",
				"Bytes": 197987
			},
			{
				"Url": "https://www.reddit.com/r/learnprogramming",
				"Start": "2023-08-20T12:55:29.96919-07:00",
				"End": "2023-08-20T12:55:30.251818-07:00",
				"Duration": 0.282631356,
				"Status": "success",
				"Bytes": 394730
			},
			{
				"Url": "http://meemeals.com",
				"Start": "2023-08-20T12:55:30.677269-07:00",
				"End": "2023-08-20T12:55:30.866401-07:00",
				"Duration": 0.189135262,
				"Status": "success",
				"Bytes": 2444
			},
			{
				"Url": "http://meemeals.org",
				"Start": "2023-08-20T12:55:30.866691-07:00",
				"End": "0001-01-01T00:00:00Z",
				"Duration": 0,
				"Status": "failure",
				"Bytes": 0
			}
		],
		"Start": "2023-08-20T12:55:28.748964-07:00",
		"End": "2023-08-20T12:55:30.869221-07:00",
		"Duration": 2.120279591,
		"Status": "success"
	}
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
	"TotalTime": 8.158262352,
	"ChAvg": 1.7709771129999998,
	"SyncAvg": 2.30806952
}

```

Path: `http://localhost:8080/fetch/validate`
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
	"Valid": false,
	"InValidUrls": [
		"http://meemeals.org"
	]
}
```
