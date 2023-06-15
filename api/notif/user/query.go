package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	user1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUser(ctx context.Context, in *npool.GetUserNotifRequest) (*npool.GetUserNotifResponse, error) {
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
		return &npool.GetUserNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserNotifResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUserOnly(ctx context.Context, in *npool.GetUserNotifOnlyRequest) (*npool.GetUserNotifOnlyResponse, error) {
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
		return &npool.GetUserNotifOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserNotifOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserNotifOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *npool.GetUserNotifsRequest) (*npool.GetUserNotifsResponse, error) {
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
		return &npool.GetUserNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserNotifsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
