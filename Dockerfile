FROM golang:1.15-alpine3.13

WORKDIR /app
COPY . /app

RUN go build -v -a -ldflags '-extldflags "-static"' .
RUN chmod +x /app/k3s-dashboard-go

FROM alpine:3.13

EXPOSE 8080

WORKDIR /app

COPY --chown=0:0 --from=builder /app/k3s-dashboard-go /app
COPY static /app/static

ENTRYPOINT [ "./k3s-dashboard-go", "start"]
