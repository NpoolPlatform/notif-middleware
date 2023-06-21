package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAnnouncements(ctx context.Context, in *npool.GetAnnouncementsRequest) (*npool.GetAnnouncementsResponse, error) {
	handler, err := amt1.NewHandler(
		ctx,
		amt1.WithConds(in.GetConds()),
		amt1.WithOffset(in.Offset),
		amt1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAnnouncements",
			"In", in,
			"Error", err,
		)
		return &npool.GetAnnouncementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAnnouncements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAnnouncements",
			"In", in,
			"Error", err,
		)
		return &npool.GetAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

//nolint
func (s *Server) GetAnnouncement(ctx context.Context, in *npool.GetAnnouncementRequest) (*npool.GetAnnouncementResponse, error) {
	handler, err := amt1.NewHandler(
		ctx,
		amt1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAnnouncement",
			"In", in,
			"error", err,
		)
		return &npool.GetAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAnnouncement",
			"In", in,
			"error", err,
		)
		return &npool.GetAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementResponse{
		Info: info,
	}, nil
}
