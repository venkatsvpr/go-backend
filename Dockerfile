FROM golang:1.14
MAINTAINER Venkata Krishnan

ADD VERSION .
ENV SOURCES /go/src/github.com/venkatsvpr/go-backend
COPY . ${SOURCES}
RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8086
EXPOSE 8086

ENTRYPOINT  go-backend