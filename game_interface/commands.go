package game_interface

import (
	"context"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// Commands 是远程命令收发实现。
type Commands struct {
	client game_interface_pb.CommandsServiceClient
}

// SendSettingsCommand 向租赁服发送 Sizukana 命令且无视返回值。
func (c *Commands) SendSettingsCommand(ctx context.Context, command string, dimensional bool) error {
	_, err := c.client.SendSettingsCommand(ctx, &game_interface_pb.SendSettingsCommandRequest{Command: command, Dimensional: dimensional})
	return err
}

// SendPlayerCommand 向租赁服发送玩家命令并无视返回值。
func (c *Commands) SendPlayerCommand(ctx context.Context, command string) error {
	_, err := c.client.SendPlayerCommand(ctx, &game_interface_pb.SendCommandRequest{Command: command})
	return err
}

// SendWSCommand 向租赁服发送 WebSocket 命令并无视返回值。
func (c *Commands) SendWSCommand(ctx context.Context, command string) error {
	_, err := c.client.SendWSCommand(ctx, &game_interface_pb.SendCommandRequest{Command: command})
	return err
}

// SendPlayerCommandWithResp 向租赁服发送玩家命令并返回命令输出。
func (c *Commands) SendPlayerCommandWithResp(ctx context.Context, command string) (*packet_pb.CommandOutput, error) {
	resp, err := c.client.SendPlayerCommandWithResponse(ctx, &game_interface_pb.SendCommandWithResponseRequest{Command: command})
	return commandOutput(resp, err)
}

// SendWSCommandWithResp 向租赁服发送 WebSocket 命令并返回命令输出。
func (c *Commands) SendWSCommandWithResp(ctx context.Context, command string) (*packet_pb.CommandOutput, error) {
	resp, err := c.client.SendWSCommandWithResponse(ctx, &game_interface_pb.SendCommandWithResponseRequest{Command: command})
	return commandOutput(resp, err)
}

// AwaitChangesGeneral 通过发送空指令等待租赁服更改同步完成。
func (c *Commands) AwaitChangesGeneral(ctx context.Context) error {
	_, err := c.client.AwaitChangesGeneral(ctx, &game_interface_pb.AwaitChangesGeneralRequest{})
	return err
}

// SendChat 发送聊天消息。
func (c *Commands) SendChat(ctx context.Context, content string) error {
	_, err := c.client.SendChat(ctx, &game_interface_pb.SendChatRequest{Content: content})
	return err
}

// Title 发送标题文本。
func (c *Commands) Title(ctx context.Context, message string) error {
	_, err := c.client.Title(ctx, &game_interface_pb.TitleRequest{Message: message})
	return err
}

func commandOutput(resp *game_interface_pb.SendCommandWithResponseResponse, err error) (*packet_pb.CommandOutput, error) {
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Output == nil {
		return &packet_pb.CommandOutput{}, nil
	}
	return resp.Output, nil
}
