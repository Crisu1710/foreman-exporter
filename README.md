# Foreman Prometheus Exporter (!!! under development !!!)

This is a Prometheus exporter for [Foreman](https://www.theforeman.org).

## Prerequisites

* [Go](https://golang.org/doc/)

## Installation

### From sources

```bash
$ git clone https://github.com/Crisu1710/foreman-exporter.git
$ cd foreman-exporter
```

Build the binary:
```bash
$ RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o foreman-exporter .
```

## Usage

To run the exporter, type the following commands:


```bash
$ export FOREMAN_HOST="foreman.example.com" #Set Foremen host
$ echo -n USER:PASSWORD | base64
$ export FOREMAN_PW="Basic THE_OUTPUT_FROM_ECHO" #Password for the API User
```
Binary:
```bash
$ ./foreman_exporter
```

Sources:
```bash
$ go run main.go
```

## Using K8s

```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: foreman-exporter-deploy
  labels:
    app.kubernetes.io/name: foreman-exporter
    app.kubernetes.io/version: "0.1.0"
spec:
  replicas: 1
  selector:
    matchLabels:
      app: foreman-exporter
  template:
    metadata:
      labels:
        app: foreman-exporter
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "2112"
    spec:
      containers:
        - name: foreman-exporter
          image: ghcr.io/crisu1710/foreman_exporter:0.1.0
          env:
            - name: FOREMAN_HOST
              value: foreman.example.com
            - name: FOREMAN_PW
              value: "Basic THE_OUTPUT_FROM_ECHO" #echo -n USER:PASSWORD | base64
          ports:
            - containerPort: 2112
              name: metrics
```

## Available Prometheus metrics

|            Metric name            | Description                                                     |
|:---------------------------------:|-----------------------------------------------------------------|
|     foreman_puppet_last_report    | Timestamp of the last puppet run of each host                   |   


