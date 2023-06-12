package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/user"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteUserAnnouncement(ctx context.Context, in *npool.DeleteUserAnnouncementRequest) (*npool.DeleteUserAnnouncementResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := announcement1.NewHandler(
		ctx,
		handler.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteUserAnnouncement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserAnnouncement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteUserAnnouncementResponse{
		Info: info,
	}, nil
}
