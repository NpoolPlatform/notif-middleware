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

func (s *Server) CreateNotifUser(ctx context.Context, in *npool.CreateNotifUserRequest) (*npool.CreateNotifUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateNotifUser",
			"In", in,
		)
		return &npool.CreateNotifUserResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(req.EntID, false),
		user1.WithAppID(req.AppID, true),
		user1.WithUserID(req.UserID, true),
		user1.WithEventType(req.EventType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNotifUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNotifUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateNotifUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNotifUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNotifUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateNotifUserResponse{
		Info: info,
	}, nil
}
