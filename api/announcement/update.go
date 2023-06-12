package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateAnnouncement(ctx context.Context, in *npool.UpdateAnnouncementRequest) (*npool.UpdateAnnouncementResponse, error) {
	req := in.GetInfo()
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithID(req.ID),
		announcement1.WithTitle(req.Title),
		announcement1.WithContent(req.Content),
		announcement1.WithAnnouncementType(req.AnnouncementType),
		announcement1.WithEndAt(req.EndAt),
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
