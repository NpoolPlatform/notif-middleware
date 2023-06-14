package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	user1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
	req := in.GetInfo()
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID),
		user1.WithAppID(req.AppID),
		user1.WithUserID(req.UserID),
		user1.WithNotifID(req.NotifID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateUserResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateUsers(ctx context.Context, in *npool.CreateUsersRequest) (*npool.CreateUsersResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateUsersResponse{
		Infos: infos,
	}, nil
}
