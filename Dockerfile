##################################
# Building go binary
##################################
FROM golang:1.18-alpine3.15 as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go build -buildvcs=false . 

##################################
# Creating final image from binary
##################################
FROM alpine:3.15
COPY --from=builder /build/restaurant-scrapper .
RUN addgroup -g 1000 spok && adduser -u 1000 -G spok -D spok
USER spok

ENTRYPOINT ["./restaurant-scrapper"]