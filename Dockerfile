FROM golang:alpine 

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o 
ENV PROJECT = ./src
ENV GO111MODULE = on

WORKDIR csv_normalization
CMD ["csv_normalization"]