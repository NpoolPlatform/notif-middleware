package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserAnnouncement(ctx context.Context, in *npool.CreateUserAnnouncementRequest) (*npool.CreateUserAnnouncementResponse, error) {
	req := in.GetInfo()
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithAppID(req.AppID),
		announcement1.WithUserID(req.UserID),
		announcement1.WithAnnouncementID(req.AnnouncementID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserAnnouncement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateUserAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateUserAnnouncement(ctx)
	if err != nil {
		return &npool.CreateUserAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateUserAnnouncementResponse{
		Info: info,
	}, nil
}
