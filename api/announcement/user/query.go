package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserAnnouncements(ctx context.Context, in *npool.GetUserAnnouncementsRequest) (*npool.GetUserAnnouncementsResponse, error) {
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithConds(in.GetConds()),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserAnnouncements",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserAnnouncementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetUserAnnouncements(ctx)
	if err != nil {
		return &npool.GetUserAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUserAnnouncementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetUserAnnouncement(ctx context.Context, in *npool.GetUserAnnouncementRequest) (*npool.GetUserAnnouncementResponse, error) {
	handler, err := announcement1.NewHandler(
		ctx,
		handler.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserAnnouncement",
			"In", in,
			"error", err,
		)
		return &npool.GetUserAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetUserAnnouncement(ctx)
	if err != nil {
		return &npool.GetUserAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUserAnnouncementResponse{
		Info: info,
	}, nil
}
