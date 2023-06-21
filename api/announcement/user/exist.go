package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	amtuser1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAnnouncementUserConds(
	ctx context.Context,
	in *npool.ExistAnnouncementUserCondsRequest,
) (
	*npool.ExistAnnouncementUserCondsResponse,
	error,
) {
	handler, err := amtuser1.NewHandler(ctx,
		amtuser1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAnnouncementUserConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistAnnouncementUserCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistAnnouncementUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAnnouncementUserConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistAnnouncementUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAnnouncementUserCondsResponse{
		Info: info,
	}, nil
}
