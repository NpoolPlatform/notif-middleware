package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	user1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUser(ctx context.Context, in *npool.GetUserRequest) (*npool.GetUserResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUserOnly(ctx context.Context, in *npool.GetUserOnlyRequest) (*npool.GetUserOnlyResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *npool.GetUsersRequest) (*npool.GetUsersResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
		user1.WithOffset(in.GetOffset()),
		user1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
