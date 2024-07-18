# -- multistage docker build: stage #1: build stage
FROM golang:1.22.0-alpine AS build

RUN mkdir -p /go/src/github.com/wombatlabs/coinsecd

WORKDIR /go/src/github.com/wombatlabs/coinsecd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache coinsecd dependencies
RUN go mod download

COPY . .
RUN mkdir -p /coinsec/bin/
RUN go build $FLAGS -o /coinsec/bin/ ./cmd/...

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /root/

RUN apk add --no-cache ca-certificates tini

COPY --from=build /coinsec/bin/* /usr/bin/

ENTRYPOINT [ "/usr/bin/coinsecd" ]
