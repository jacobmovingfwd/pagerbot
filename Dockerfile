FROM golang:1.16.4-buster AS builder
WORKDIR /build/
RUN update-ca-certificates
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w' -o pagerbot .

FROM scratch AS runner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /pagerbot/
COPY --from=0 /build/pagerbot .
CMD ["/pagerbot/pagerbot"]  
