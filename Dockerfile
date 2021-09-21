FROM golang:1.16.7-alpine3.14 

ENV TZ=Asia/Seoul

RUN apk add --no-cache make g++ cmake bash git zlib-dev curl-dev proj-dev geos-dev
