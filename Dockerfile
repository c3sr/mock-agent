FROM golang:1.15

# Install any system dependencies up here
# go 15 is old enough that python3 is just 3.7 which is the highest version go-python works with
RUN apt-get update && apt-get install apt-utils python3 python3-dev -y

# Make our go directory structure
RUN  mkdir -p /go/src/github.com/c3sr/mock-agent
WORKDIR /go/src/github.com/c3sr/mock-agent

# Handle Go Depedencies and cache it as a layer
COPY go.mod go.mod
#COPY dlframework/go.mo? dlframework/
RUN go mod download

# Get the rest of the project in
COPY . .
RUN go install -tags=nogpu .

CMD ["/go/bin/mock-agent", "worker"]