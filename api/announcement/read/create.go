package read

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/read"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/read"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReadAnnouncement(ctx context.Context, in *npool.CreateReadAnnouncementRequest) (*npool.CreateReadAnnouncementResponse, error) {
	req := in.GetInfo()
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithAppID(req.AppID),
		announcement1.WithUserID(req.UserID),
		announcement1.WithAnnouncementID(req.AnnouncementID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReadAnnouncement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateReadAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateReadAnnouncement(ctx)
	if err != nil {
		return &npool.CreateReadAnnouncementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateReadAnnouncementResponse{
		Info: info,
	}, nil
}
