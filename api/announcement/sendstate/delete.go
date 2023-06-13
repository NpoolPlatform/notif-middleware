package sendstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/send"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/send"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteSendAnnouncement(ctx context.Context, in *npool.DeleteSendAnnouncementRequest) (*npool.DeleteSendAnnouncementResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := announcement1.NewHandler(
		ctx,
		handler.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSendAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSendAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteSendAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSendAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSendAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteSendAnnouncementResponse{
		Info: info,
	}, nil
}
