FROM golang:1.16-alpine AS build

WORKDIR /src/go-mon/
ADD app /src/go-mon/app
ADD config /src/go-mon/config
ADD monitor /src/go-mon/monitor
COPY go.* main.go /src/go-mon/
RUN CGO_ENABLED=0 go build -o /bin/go-mon

ENTRYPOINT ["/bin/go-mon"]
