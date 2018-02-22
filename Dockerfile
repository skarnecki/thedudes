FROM golang:1.9.2

RUN apt-get update
RUN apt-get install -y unzip jq python-pip

ARG BUILD_NUMBER

ADD . $GOPATH/src/github.com/skarnecki/thedudes
RUN go install -ldflags="-X main.Version=0.0.$BUILD_NUMBER" github.com/skarnecki/thedudes

ENTRYPOINT $GOPATH/bin/thedudes
EXPOSE 8080