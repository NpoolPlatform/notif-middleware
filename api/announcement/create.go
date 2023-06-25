package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAnnouncement(ctx context.Context, in *npool.CreateAnnouncementRequest) (*npool.CreateAnnouncementResponse, error) {
	req := in.GetInfo()
	handler, err := amt1.NewHandler(
		ctx,
		amt1.WithAppID(req.AppID),
		amt1.WithLangID(req.LangID),
		amt1.WithTitle(req.Title),
		amt1.WithContent(req.Content),
		amt1.WithChannel(req.Channel),
		amt1.WithAnnouncementType(req.AnnouncementType),
		amt1.WithStartAt(req.StartAt, req.EndAt),
		amt1.WithEndAt(req.StartAt, req.EndAt),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAnnouncement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAnnouncement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAnnouncementResponse{
		Info: info,
	}, nil
}
