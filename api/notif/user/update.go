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

func (s *Server) UpdateNotifUser(ctx context.Context, in *npool.UpdateNotifUserRequest) (*npool.UpdateNotifUserResponse, error) {
	req := in.GetInfo()
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID),
		user1.WithUserID(req.UserID),
		user1.WithEventType(req.EventType),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNotifUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNotifUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateNotifUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNotifUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNotifUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateNotifUserResponse{
		Info: info,
	}, nil
}
