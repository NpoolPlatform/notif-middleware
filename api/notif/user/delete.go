package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	user1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteUserNotif(ctx context.Context, in *npool.DeleteUserNotifRequest) (*npool.DeleteUserNotifResponse, error) {
	req := in.GetInfo()
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID),
		user1.WithAppID(req.AppID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserNotif",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserNotif",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteUserNotifResponse{
		Info: info,
	}, nil
}
