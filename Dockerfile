FROM golang:1.13-alpine AS builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o server .

FROM alpine
RUN mkdir app
COPY --from=builder /go/src/app/server ./app

CMD "./app/server"