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

func (s *Server) ExistNotifUser(ctx context.Context, in *npool.ExistNotifUserRequest) (*npool.ExistNotifUserResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistNotifUser",
			"In", in,
			"Error", err,
		)
		return &npool.ExistNotifUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistNotifUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistNotifUser",
			"In", in,
			"Error", err,
		)
		return &npool.ExistNotifUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistNotifUserResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistNotifUserConds(
	ctx context.Context,
	in *npool.ExistNotifUserCondsRequest,
) (
	*npool.ExistNotifUserCondsResponse,
	error,
) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistNotifUser",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistNotifUserCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistNotifUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistNotifUser",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistNotifUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistNotifUserCondsResponse{
		Info: info,
	}, nil
}
