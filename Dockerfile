FROM golang:1.19-alpine as build-env
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
FROM scratch
COPY --from=build-env /go/bin/app /go/bin/app
WORKDIR /go/bin


ENTRYPOINT ["/go/bin/app"]