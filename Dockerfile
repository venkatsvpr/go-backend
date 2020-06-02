FROM golang:1.14
MAINTAINER Venkata Krishnan

ENV SOURCES /go/src/github.com/venkatsvpr/go-backend/src/
COPY src/ ${SOURCES}/src
RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

