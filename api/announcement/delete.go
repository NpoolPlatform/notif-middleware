package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAnnouncement(ctx context.Context, in *npool.DeleteAnnouncementRequest) (*npool.DeleteAnnouncementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAnnouncement",
			"In", in,
		)
		return &npool.DeleteAnnouncementResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithID(req.ID, true),
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
