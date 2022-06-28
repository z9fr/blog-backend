FROM golang:1.18 AS builder

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /src/app .

CMD ["./app"]
