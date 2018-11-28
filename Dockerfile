FROM golang:1.11.2-alpine as builder
RUN apk add --no-cache git
RUN mkdir /build
COPY main.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 go get github.com/elastic/beats/filebeat
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o filebeat .

FROM ubuntu:bionic
RUN apt-get update
RUN apt-get install -y nano
COPY filebeat.yml /filebeat/filebeat.yml
COPY --from=builder /build/filebeat /filebeat/
WORKDIR /filebeat
ENTRYPOINT [ "/filebeat/filebeat" ]
CMD [ "-c", "filebeat.yml" ]
