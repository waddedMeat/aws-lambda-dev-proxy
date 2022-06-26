FROM golang:1.18 as BUILD

WORKDIR /go/src

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags="-w -s" -o /go/bin/aws-lambda-proxy

FROM scratch
COPY --from=BUILD /go/bin/aws-lambda-proxy /go/bin/aws-lambda-proxy

ENTRYPOINT ["/go/bin/aws-lambda-proxy"]
