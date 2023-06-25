package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateAnnouncement(ctx context.Context, in *npool.UpdateAnnouncementRequest) (*npool.UpdateAnnouncementResponse, error) {
	req := in.GetInfo()
	handler, err := amt1.NewHandler(
		ctx,
		amt1.WithID(req.ID),
		amt1.WithTitle(req.Title),
		amt1.WithContent(req.Content),
		amt1.WithAnnouncementType(req.AnnouncementType),
		amt1.WithStartAt(req.StartAt, req.EndAt),
		amt1.WithEndAt(req.StartAt, req.EndAt),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAnnouncement",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAnnouncement",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAnnouncementResponse{
		Info: info,
	}, nil
}
