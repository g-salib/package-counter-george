FROM golang:1.18-alpine AS build
WORKDIR /go/src/app
RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add git bash build-base pkgconf sudo


#RUN apk add --no-cache librdkafka-dev
#RUN git clone --branch v1.3.0 https://github.com/edenhill/librdkafka.git && cd librdkafka && ./configure --prefix /usr && make && make install
 
ENV GO111MODULE=on
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
 
COPY . .
# RUN go get ./main
RUN GOOS=linux go build -tags musl -a -o main .
FROM alpine:edge
RUN apk update && apk upgrade
RUN apk -U add ca-certificates
RUN apk update && apk upgrade
#RUN apk add --no-cache librdkafka-dev
COPY --from=build /go/src/app/main /
ENTRYPOINT ["./main"]