//go:build !windows

package client

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	shmipc "github.com/cloudwego/shmipc-go"
)

// ShmIPC 是基于共享内存的本机 IPC 网络实现。
type ShmIPC struct{}

// DialContext 连接到指定名称对应的 shmipc 监听器。
func (ShmIPC) DialContext(ctx context.Context, address string) (net.Conn, error) {
	path, err := shmIPCAddress(address)
	if err != nil {
		return nil, err
	}
	conf := shmipc.DefaultSessionManagerConfig()
	conf.Network = "unix"
	conf.Address = path
	conf.ShareMemoryPathPrefix = filepath.Join(os.TempDir(), "emptydea-core", "shmipc-"+shmIPCFileName(address))
	conf.QueuePath = conf.ShareMemoryPathPrefix + "_queue"
	if deadline, ok := ctx.Deadline(); ok {
		conf.ConnectionWriteTimeout = time.Until(deadline)
	}
	manager, err := shmipc.NewSessionManager(conf)
	if err != nil {
		return nil, err
	}
	stream, err := manager.GetStream()
	if err != nil {
		_ = manager.Close()
		return nil, err
	}
	return &shmIPCConn{Stream: stream, manager: manager}, nil
}

func shmIPCAddress(address string) (string, error) {
	if address == "" {
		return "", fmt.Errorf("shmipc: empty address")
	}
	if filepath.IsAbs(address) {
		return address, nil
	}
	return filepath.Join(os.TempDir(), "emptydea-core", shmIPCFileName(address)), nil
}

func shmIPCFileName(address string) string {
	replacer := strings.NewReplacer("/", "_", "\\", "_", ":", "_")
	name := replacer.Replace(address)
	if filepath.Ext(name) == "" {
		name += ".sock"
	}
	return name
}

type shmIPCConn struct {
	*shmipc.Stream
	manager *shmipc.SessionManager
	once    sync.Once
	err     error
}

func (c *shmIPCConn) Close() error {
	c.once.Do(func() {
		c.err = errors.Join(c.Stream.Close(), c.manager.Close())
	})
	return c.err
}

func init() {
	RegisterNetwork("shmipc", func() Network { return ShmIPC{} })
}
