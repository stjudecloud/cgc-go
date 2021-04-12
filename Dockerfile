FROM golang:alpine as builder
WORKDIR $GOPATH/src/cgc

RUN apk add --no-cache ca-certificates

COPY . .

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "10001" \    
    "user"

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /go/bin/cgc

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/cgc /go/bin/cgc

USER user:user
ENTRYPOINT ["/go/bin/cgc"]
CMD ["--help"]

