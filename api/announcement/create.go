package announcement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAnnouncement(ctx context.Context, in *npool.CreateAnnouncementRequest) (*npool.CreateAnnouncementResponse, error) {
	req := in.GetInfo()
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithEntID(req.EntID, false),
		announcement1.WithAppID(req.AppID, true),
		announcement1.WithLangID(req.LangID, true),
		announcement1.WithTitle(req.Title, true),
		announcement1.WithContent(req.Content, true),
		announcement1.WithChannel(req.Channel, true),
		announcement1.WithAnnouncementType(req.AnnouncementType, true),
		announcement1.WithStartAt(req.StartAt, true),
		announcement1.WithEndAt(req.EndAt, true),
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
