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

func (c *Commands) SendSettingsCommand(ctx context.Context, command string, dimensional bool) error {
	_, err := c.client.SendSettingsCommand(ctx, &game_interface_pb.SendSettingsCommandRequest{Command: command, Dimensional: dimensional})
	return err
}

func (c *Commands) SendPlayerCommand(ctx context.Context, command string) error {
	_, err := c.client.SendPlayerCommand(ctx, &game_interface_pb.SendCommandRequest{Command: command})
	return err
}

func (c *Commands) SendWSCommand(ctx context.Context, command string) error {
	_, err := c.client.SendWSCommand(ctx, &game_interface_pb.SendCommandRequest{Command: command})
	return err
}

func (c *Commands) SendPlayerCommandWithResp(ctx context.Context, command string) (*packet_pb.CommandOutput, error) {
	resp, err := c.client.SendPlayerCommandWithResponse(ctx, &game_interface_pb.SendCommandWithResponseRequest{Command: command})
	return commandOutput(resp, err)
}

func (c *Commands) SendWSCommandWithResp(ctx context.Context, command string) (*packet_pb.CommandOutput, error) {
	resp, err := c.client.SendWSCommandWithResponse(ctx, &game_interface_pb.SendCommandWithResponseRequest{Command: command})
	return commandOutput(resp, err)
}

func (c *Commands) AwaitChangesGeneral(ctx context.Context) error {
	_, err := c.client.AwaitChangesGeneral(ctx, &game_interface_pb.AwaitChangesGeneralRequest{})
	return err
}

func (c *Commands) SendChat(ctx context.Context, content string) error {
	_, err := c.client.SendChat(ctx, &game_interface_pb.SendChatRequest{Content: content})
	return err
}

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
