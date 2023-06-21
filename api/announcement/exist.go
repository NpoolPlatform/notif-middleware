package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) ExistAnnouncement(
	ctx context.Context,
	in *npool.ExistAnnouncementRequest,
) (
	*npool.ExistAnnouncementResponse,
	error,
) {
	handler, err := amt1.NewHandler(
		ctx,
		amt1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAnnouncementResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAnnouncementConds(
	ctx context.Context,
	in *npool.ExistAnnouncementCondsRequest,
) (
	*npool.ExistAnnouncementCondsResponse,
	error,
) {
	handler, err := amt1.NewHandler(
		ctx,
		amt1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAnnouncementConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAnnouncementCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistAnnouncementConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAnnouncementConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAnnouncementCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAnnouncementCondsResponse{
		Info: exist,
	}, nil
}
