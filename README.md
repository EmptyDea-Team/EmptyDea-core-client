# EmptyDea-core-client

`EmptyDea-core-client` 是 EmptyDeaCore 的 gRPC 客户端封装模块。它依赖 `EmptyDea-core-api` 的 `pb` 类型，对外提供更接近 core 源码风格的调用入口。

这个模块不依赖 `EmptyDea-core` 源码，也不直接导入 mousetunnel。客户端侧只处理 gRPC 请求、响应和必要的本地轻量类型。

## 目录结构

```text
client.go                    客户端总入口
dial.go                      gRPC 连接辅助函数
frame.go                     框架层连接控制
convertutil/                 pb 基础值解包辅助函数
resources_control/           资源层客户端封装
resources_control/uqholder/  UQHolder getter 风格封装
game_interface/              游戏交互层客户端封装
game_interface/item_stack_operation/
                             ItemStackTransaction 操作类型
game_interface/item_stack_transaction/
                             批量提交物品栈操作的客户端封装
```

## 基本用法

```go
package main

import (
	"context"
	"log"

	client "github.com/EmptyDea-Team/EmptyDea-core-client"
)

func main() {
	ctx := context.Background()

	c, conn, err := client.Dial(ctx, "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ok, err := c.Frame().Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ping:", ok)
}
```

## 启动连接

`FrameConfig` 来自 `EmptyDea-core-api/define`，在客户端中以类型别名导出：

```go
_, err := c.Frame().StartConnection(ctx, client.FrameConfig{
	AuthServer:     "http://127.0.0.1:8080",
	UserToken:      "token",
	ServerCode:     "server-code",
	ServerPassword: "password",
})
```

## 访问服务

```go
resources := c.Resources()
gameInterface := c.GameInterface()

_ = resources
_ = gameInterface
```

`Resources()` 下提供 packet、inventory、container、constant packet、UQHolder 等资源能力。

`GameInterface()` 下提供命令、方块放置、容器交互、物品转移、ItemStackTransaction、PlayerKit 等游戏交互能力。

## 本地开发

当前模块直接依赖 `github.com/EmptyDea-Team/EmptyDea-core-api`。本地联调时如需改用 sibling 模块，可临时在本机添加 replace。

验证：

```bash
cd EmptyDea-core-client
go test ./...
```
