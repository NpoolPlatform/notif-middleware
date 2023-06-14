package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	user1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteUser(ctx context.Context, in *npool.DeleteUserRequest) (*npool.DeleteUserResponse, error) {
	req := in.GetInfo()
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteUserResponse{
		Info: info,
	}, nil
}
