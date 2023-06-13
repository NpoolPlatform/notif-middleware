package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/read"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/read"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteReadAnnouncement(ctx context.Context, in *npool.DeleteReadAnnouncementRequest) (*npool.DeleteReadAnnouncementResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := announcement1.NewHandler(
		ctx,
		handler.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteReadAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteReadAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteReadAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteReadAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteReadAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteReadAnnouncementResponse{
		Info: info,
	}, nil
}
