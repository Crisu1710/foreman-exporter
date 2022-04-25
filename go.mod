module example.com/foreman

go 1.17

replace (
	example.com/mymetrics => ./cmd/mymetrics
	example.com/mycollector => ./cmd/mycollector
	example.com/mytime => ./cmd/mytime
)

require (
	example.com/mymetrics v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.12.1
	example.com/mycollector v0.0.0-00010101000000-000000000000 // indirect
	example.com/mytime v0.0.0-00010101000000-000000000000 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)
