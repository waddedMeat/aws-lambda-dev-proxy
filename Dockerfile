FROM golang:1.18 as BUILD

WORKDIR /go/src

COPY . .

RUN go get -d -v

RUN go build

#FROM scratch
#COPY --from=BUILD /go/src/aws-lambda-proxy /go/bin/aws-lambda-proxy

ENTRYPOINT ["/go/src/aws-lambda-proxy"]
