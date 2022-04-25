FROM golang:1.18-alpine3.15 as build

WORKDIR /tmp/foreman_exporter

RUN apk --no-cache add git alpine-sdk
COPY . .
RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o foreman_exporter ./

FROM scratch
LABEL name="foreman_exporter"

WORKDIR /root
COPY --from=build /tmp/foreman_exporter/foreman_exporter foreman_exporter

CMD ["./foreman_exporter"]
