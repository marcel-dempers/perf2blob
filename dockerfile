FROM golang:1.12.1-alpine3.9 as build
RUN apk add --no-cache git curl

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY ./src /go/src/app/

RUN go build -o perf2blob
################################################################################################################

FROM alpine:3.9
RUN mkdir -p /app
COPY --from=build /go/src/app/perf2blob /app/
WORKDIR /app

CMD ["./perf2blob"]
