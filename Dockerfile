ARG APP_NAME=foreman-exporter

FROM golang:1.18-alpine as build

ARG APP_NAME
WORKDIR /app

RUN apk --no-cache add git alpine-sdk
COPY . .
RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./$APP_NAME

FROM scratch

ARG APP_NAME

COPY --from=build /app/$APP_NAME /app

ENTRYPOINT ["/app"]
