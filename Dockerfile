# syntax=docker/dockerfile:1


FROM golang:1.18-alpine
ADD . /go/src/go-fetch
# ADD file:8e81116368669ed3dd361bc898d61bff249f524139a239fdaf3ec46869a39921 in / 

# Create a working directory in our image
WORKDIR /go/src/go-fetch
RUN go get go-fetch
RUN go install
# Copies our go mod and go sum files into our working directory
COPY go.mod ./
COPY go.sum ./

# Download all go dependicies
RUN go mod download

# Copy source code binary into working directory
COPY *.go ./

# Build the app image
RUN go build -o /go-fetch

# Ports exposed on container
EXPOSE 8080

# Enter go binary to run application
ENTRYPOINT ["/go/bin/go-fetch"]
# ENTRYPOINT ["/go-fetch"]

#Build docker image
# docker build --tag go-fetch .

#Run docker container w/ image in detached mode exposing port 8080
# docker run -p 8080:8080 -it go-fetch
# docker run -p 8080:8080 -d go-fetch