//go:build !windows

package client

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func dialPipeContext(ctx context.Context, address string) (net.Conn, error) {
	path, err := pipeAddress(address)
	if err != nil {
		return nil, err
	}
	var d net.Dialer
	return d.DialContext(ctx, "unix", path)
}

func pipeAddress(address string) (string, error) {
	if address == "" {
		return "", fmt.Errorf("pipe: empty address")
	}
	if filepath.IsAbs(address) {
		return address, nil
	}
	return filepath.Join(os.TempDir(), "emptydea-core", pipeFileName(address)), nil
}

func pipeFileName(address string) string {
	replacer := strings.NewReplacer("/", "_", "\\", "_", ":", "_")
	name := replacer.Replace(address)
	if filepath.Ext(name) == "" {
		name += ".sock"
	}
	return name
}
