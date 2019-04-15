FROM golang:1.12.1-alpine3.9 as build
RUN apk add --no-cache git curl

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY ./src /go/src/app/

RUN go build -o perf2blob
################################################################################################################

FROM ubuntu:18.04
RUN apt-get update && apt-get install -y git curl gcc make bison flex elfutils libelf-dev libdw-dev libaudit-dev xz-utils

RUN mkdir /linuxtools
WORKDIR /linuxtools

RUN curl https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.9.125.tar.xz -o linux.tar.xz

RUN ls -l 
RUN tar -xJf linux.tar.xz
RUN cd /linuxtools/linux-4.9.125/tools/perf && ls -l && make O=/tmp/ && ls /tmp/

RUN uname -r
RUN mkdir -p /app
COPY --from=build /go/src/app/perf2blob /app/
WORKDIR /app

CMD ["./perf2blob"]
