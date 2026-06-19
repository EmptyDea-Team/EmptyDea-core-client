package client

import (
	_ "github.com/mostynb/go-grpc-compression/snappy"
	_ "github.com/mostynb/go-grpc-compression/zstd"
	_ "google.golang.org/grpc/encoding/gzip"
)
