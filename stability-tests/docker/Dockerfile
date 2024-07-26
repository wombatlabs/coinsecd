ARG COINSECD_IMAGE
ARG COINSECMINER_IMAGE

FROM ${COINSECD_IMAGE} as coinsecd
FROM ${COINSECMINER_IMAGE} as coinsecminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/wombatlabs/coinsecd

WORKDIR /go/src/github.com/wombatlabs/coinsecd

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=coinsecd /app/ /app/
COPY --from=coinsecminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/wombatlabs/coinsecd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]