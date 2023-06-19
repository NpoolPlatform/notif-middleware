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

func (s *Server) CreateUserNotif(ctx context.Context, in *npool.CreateUserNotifRequest) (*npool.CreateUserNotifResponse, error) {
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
			"CreateUserNotif",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserNotif",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateUserNotifResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateUserNotifs(ctx context.Context, in *npool.CreateUserNotifsRequest) (*npool.CreateUserNotifsResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateUserNotifsResponse{
		Infos: infos,
	}, nil
}
