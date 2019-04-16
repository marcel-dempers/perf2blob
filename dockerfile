#FROM golang:1.12.1-alpine3.9 as build
FROM golang:1.12.4-stretch as build
RUN apt-get update && apt-get install -y git curl

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

RUN go get github.com/Azure/azure-storage-blob-go/azblob
COPY ./src /go/src/app/

RUN GOOS=linux GOARCH=amd64 go build -o perf2blob
################################################################################################################

#See my dockerfiles repo for perf images:
#https://github.com/marcel-dempers/my-desktop/tree/master/dockerfiles/perf
FROM aimvector/perf:4.9.125 
RUN apt-get update && apt-get install -y ca-certificates
RUN mkdir -p /app
COPY --from=build /go/src/app/perf2blob /app/
WORKDIR /app
ENTRYPOINT ["./perf2blob"]
