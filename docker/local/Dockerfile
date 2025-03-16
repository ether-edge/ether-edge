FROM golang:1.20-alpine AS builder

RUN apk add make git

WORKDIR /ether-edge

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /ether-edge

COPY --from=builder /ether-edge/ether-edge ./
COPY ./docker/local/ether-edge.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478 5001

ENTRYPOINT ["./ether-edge.sh"]
