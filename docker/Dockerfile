# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/wombatlabs/coinsecd

WORKDIR /go/src/github.com/wombatlabs/coinsecd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache coinsecd dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o coinsecd .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/wombatlabs/coinsecd/coinsecd /app/
COPY --from=build /go/src/github.com/wombatlabs/coinsecd/infrastructure/config/sample-coinsecd.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
