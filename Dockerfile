FROM node:14.16.1-alpine3.13 as nodebuilder
#FROM node:14.16.1-buster-slim as nodebuilder

WORKDIR /app
COPY ./static /app

#RUN npm install
RUN apk add --no-cache git make build-base
RUN yarn install
RUN yarn build

FROM golang:1.16-alpine3.13 as gobuilder

WORKDIR /app
COPY . /app

RUN apk add --no-cache git make build-base
RUN go build -v -a -ldflags '-extldflags "-static"' .
RUN chmod +x /app/k3s-dashboard-go

FROM alpine:3.13

EXPOSE 8080

WORKDIR /app

COPY --chown=0:0 --from=gobuilder /app/k3s-dashboard-go /app
COPY --chown=0:0 --from=nodebuilder /app/build /app/static

ENTRYPOINT [ "./k3s-dashboard-go", "start"]
