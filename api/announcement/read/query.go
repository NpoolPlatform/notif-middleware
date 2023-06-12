package read

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/read"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/read"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetReadAnnouncements(ctx context.Context, in *npool.GetReadAnnouncementsRequest) (*npool.GetReadAnnouncementsResponse, error) {
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithConds(in.GetConds()),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadAnnouncements",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadAnnouncementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetReadAnnouncements(ctx)
	if err != nil {
		return &npool.GetReadAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadAnnouncementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetReadAnnouncement(ctx context.Context, in *npool.GetReadAnnouncementRequest) (*npool.GetReadAnnouncementResponse, error) {
	handler, err := announcement1.NewHandler(
		ctx,
		handler.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadAnnouncement",
			"In", in,
			"error", err,
		)
		return &npool.GetReadAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetReadAnnouncement(ctx)
	if err != nil {
		return &npool.GetReadAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadAnnouncementResponse{
		Info: info,
	}, nil
}
