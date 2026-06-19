//go:build windows

package client

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/Microsoft/go-winio"
)

func dialPipeContext(ctx context.Context, address string) (net.Conn, error) {
	path, err := pipeAddress(address)
	if err != nil {
		return nil, err
	}
	return winio.DialPipeContext(ctx, path)
}

func pipeAddress(address string) (string, error) {
	if address == "" {
		return "", fmt.Errorf("pipe: empty address")
	}
	if strings.HasPrefix(address, `\\.\pipe\`) {
		return address, nil
	}
	return `\\.\pipe\emptydea-core-` + sanitizePipeName(address), nil
}

func sanitizePipeName(address string) string {
	replacer := strings.NewReplacer(`\`, "_", `/`, "_", ":", "_")
	return replacer.Replace(address)
}
