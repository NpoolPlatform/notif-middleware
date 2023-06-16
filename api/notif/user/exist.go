package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	user1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) ExistUser(ctx context.Context, in *npool.ExistUserNotifRequest) (*npool.ExistUserNotifResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUser",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUser",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistUserNotifResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistUserConds(ctx context.Context, in *npool.ExistUserNotifCondsRequest) (*npool.ExistUserNotifCondsResponse, error) {
	handler, err := user1.NewHandler(ctx,
		user1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserNotif",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistUserNotifCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserNotif",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistUserNotifCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistUserNotifCondsResponse{
		Info: info,
	}, nil
}
