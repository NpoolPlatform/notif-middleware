package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAnnouncement(ctx context.Context, in *npool.DeleteAnnouncementRequest) (*npool.DeleteAnnouncementResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := amt1.NewHandler(
		ctx,
		amt1.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteAnnouncementResponse{
		Info: info,
	}, nil
}
