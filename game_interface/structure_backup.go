package game_interface

import (
	"context"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	"github.com/google/uuid"
)

// StructureBackup 是远程结构备份实现。
type StructureBackup struct {
	client game_interface_pb.StructureBackupServiceClient
}

// BackupStructure 通过 structure 命令保存 pos 处的方块并返回备份 UUID。
func (s *StructureBackup) BackupStructure(ctx context.Context, pos *protocol_pb.BlockPos) (uuid.UUID, error) {
	resp, err := s.client.BackupStructure(ctx, &game_interface_pb.BackupStructureRequest{Pos: pos})
	return parseUUIDResponse(resp, err)
}

// BackupOffset 通过 structure 命令保存 pos 到 pos+offset 范围内的方块并返回备份 UUID。
func (s *StructureBackup) BackupOffset(ctx context.Context, pos *protocol_pb.BlockPos, offset *protocol_pb.BlockPos) (uuid.UUID, error) {
	resp, err := s.client.BackupOffset(ctx, &game_interface_pb.BackupOffsetRequest{Pos: pos, Offset: offset})
	return parseUUIDResponse(resp, err)
}

// RevertStructure 在 pos 处恢复 UUID 对应的结构备份。
func (s *StructureBackup) RevertStructure(ctx context.Context, uniqueID uuid.UUID, pos *protocol_pb.BlockPos) error {
	_, err := s.client.RevertStructure(ctx, &game_interface_pb.RevertStructureRequest{UUID: uniqueID.String(), Pos: pos})
	return err
}

// DeleteStructure 删除 UUID 对应的结构备份。
func (s *StructureBackup) DeleteStructure(ctx context.Context, uniqueID uuid.UUID) error {
	_, err := s.client.DeleteStructure(ctx, &game_interface_pb.DeleteStructureRequest{UUID: uniqueID.String()})
	return err
}

func parseUUIDResponse(resp *game_interface_pb.BackupStructureResponse, err error) (uuid.UUID, error) {
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.Parse(resp.UUID)
}
