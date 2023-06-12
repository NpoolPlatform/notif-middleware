package send

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/send"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/send"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSendAnnouncement(ctx context.Context, in *npool.CreateSendAnnouncementRequest) (*npool.CreateSendAnnouncementResponse, error) {
	req := in.GetInfo()
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithAppID(req.AppID),
		announcement1.WithUserID(req.UserID),
		announcement1.WithAnnouncementID(req.AnnouncementID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendAnnouncement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateSendAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateSendAnnouncement(ctx)
	if err != nil {
		return &npool.CreateSendAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSendAnnouncementResponse{
		Info: info,
	}, nil
}
