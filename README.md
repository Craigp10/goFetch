# Welcome to goFetch

goFetch is a quick program I made to gain some experience using multiple technologies, including go's sync package, docker, and http go framework gin.

The functionality of the app is to fetch the provided urls utilizing concurrency vs iteration, and provide metrics to display the advantages of each approach.

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
