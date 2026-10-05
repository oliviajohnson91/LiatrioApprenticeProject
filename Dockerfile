FROM golang:1.27.1 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server .

FROM alpine:latest

COPY --from=builder /build/server /server

EXPOSE 3000

ENTRYPOINT ["/server"]