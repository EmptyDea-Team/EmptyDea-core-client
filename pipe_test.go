package client

import (
	"context"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	api_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestPipeAddress(t *testing.T) {
	address, err := PipeAddress("test-listener")
	if err != nil {
		t.Fatalf("PipeAddress() error = %v", err)
	}
	if runtime.GOOS == "windows" {
		if address != `\\.\pipe\emptydea-core-test-listener` {
			t.Fatalf("PipeAddress() = %q", address)
		}
		return
	}
	if !filepath.IsAbs(address) {
		t.Fatalf("PipeAddress() = %q, want absolute path", address)
	}
	if filepath.Base(address) != "test-listener.sock" {
		t.Fatalf("PipeAddress() = %q, want test-listener.sock basename", address)
	}
}

func TestDialerPipe(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("windows named pipe runtime test is covered by cross compilation")
	}
	address, err := PipeAddress("test-client-listener")
	if err != nil {
		t.Fatalf("PipeAddress() error = %v", err)
	}
	if err = os.MkdirAll(filepath.Dir(address), 0o700); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}
	_ = os.Remove(address)
	listener, err := net.Listen("unix", address)
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer func() {
		_ = listener.Close()
		_ = os.Remove(address)
	}()

	grpcServer := grpc.NewServer()
	api_pb.RegisterFrameServiceServer(grpcServer, frameService{})
	go func() {
		_ = grpcServer.Serve(listener)
	}()
	defer grpcServer.Stop()

	c, err := Dialer{}.DialContext(context.Background(), "pipe", "test-client-listener")
	if err != nil {
		t.Fatalf("DialContext() error = %v", err)
	}
	defer c.Close()

	resp, err := c.Frame().Ping(context.Background())
	if err != nil {
		t.Fatalf("Ping() error = %v", err)
	}
	if !resp {
		t.Fatal("Ping() success = false")
	}
}

type frameService struct {
	api_pb.UnimplementedFrameServiceServer
}

func (frameService) Ping(context.Context, *emptypb.Empty) (*api_pb.PingResponse, error) {
	return &api_pb.PingResponse{Success: true}, nil
}
