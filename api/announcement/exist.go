package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAnnouncement(ctx context.Context, in *npool.ExistAnnouncementRequest) (*npool.ExistAnnouncementResponse, error) { // nolint
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAnnouncementResponse{
		Info: exist,
	}, nil
}
