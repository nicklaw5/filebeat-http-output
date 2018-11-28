FROM golang:1.11.2-alpine as builder
RUN mkdir /build
COPY main.go /build/
WORKDIR /build
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o test-server .

FROM alpine
COPY --from=builder /build/test-server /test-server/
WORKDIR /test-server
ENTRYPOINT [ "/test-server/test-server" ]
CMD [ "--host", "0.0.0.0", "--port", "8888" ]
