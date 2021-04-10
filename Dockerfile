FROM golang:1.16-alpine3.13 as builder

WORKDIR /app
COPY . /app

RUN apk add --no-cache git make build-base
RUN go build -v -a -ldflags '-extldflags "-static"' .
RUN chmod +x /app/k3s-dashboard-go

FROM alpine:3.13

EXPOSE 8080

WORKDIR /app

COPY --chown=0:0 --from=builder /app/k3s-dashboard-go /app
COPY --chown=0:0 --from=builder /app/static /app/static

ENTRYPOINT [ "./k3s-dashboard-go", "start"]
