FROM golang:1.14
MAINTAINER Venkata Krishnan

ENV SOURCES /go/src/github.com/venkatsvpr/go-backend
COPY . ${SOURCES}
RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT  go-backend