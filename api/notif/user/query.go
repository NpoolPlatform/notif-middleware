//nolint:nolintlint,dupl
package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	user1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetNotifUser(ctx context.Context, in *npool.GetNotifUserRequest) (*npool.GetNotifUserResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetNotifUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifUserResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotifUserOnly(ctx context.Context, in *npool.GetNotifUserOnlyRequest) (*npool.GetNotifUserOnlyResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifUserOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifUserOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetNotifUserOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifUserOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifUserOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifUserOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotifUsers(ctx context.Context, in *npool.GetNotifUsersRequest) (*npool.GetNotifUsersResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
		user1.WithOffset(in.GetOffset()),
		user1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetNotifUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
