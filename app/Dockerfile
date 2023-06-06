FROM golang:alpine AS builder
RUN apk update && apk upgrade
#install certificates in Alpine image to establish secure communications
RUN apk --no-cache add ca-certificates make
#library linking CGO_ENABLED=0 GOOS=linux
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

#scratch for building minimal images - final docker images does not contain more than your own executable
FROM scratch
# This line will copy all certificates to final image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#copy compiled code (go binary) from bulder stage
COPY --from=builder /dist/main /
COPY .env /
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT ["/main"]
