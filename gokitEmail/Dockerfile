FROM golang:latest
MAINTAINER David Wakeling
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/henryhudson/EnterpriseComputingCA/tree/master/gokitEmail
EXPOSE 8888
CMD [ "go", "run", "/app/main.go" ]