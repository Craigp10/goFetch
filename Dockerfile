# syntax=docker/dockerfile:1


FROM golang:1.18-alpine
ADD . /go/src/practice-gin
# ADD file:8e81116368669ed3dd361bc898d61bff249f524139a239fdaf3ec46869a39921 in / 

# Create a working directory in our image
WORKDIR /go/src/practice-gin
RUN go get practice-gin
RUN go install
# Copies our go mod and go sum files into our working directory
COPY go.mod ./
COPY go.sum ./

# Download all go dependicies
RUN go mod download

# Copy source code binary into working directory
COPY *.go ./

# Build the app image
RUN go build -o /practice-gin

# Ports exposed on container
EXPOSE 8080

# Enter go binary to run application
ENTRYPOINT ["/go/bin/practice-gin"]
# ENTRYPOINT ["/practice-gin"]

#Build docker image
# docker build --tag practice-gin .

#Run docker container w/ image in detached mode exposing port 8080
# docker run -p 8080:8080 -it practice-gin
# docker run -p 8080:8080 -d practice-gin