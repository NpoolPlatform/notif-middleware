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

func (s *Server) GetSendAnnouncements(ctx context.Context, in *npool.GetSendAnnouncementsRequest) (*npool.GetSendAnnouncementsResponse, error) {
	handler, err := announcement1.NewHandler(
		ctx,
		announcement1.WithConds(in.GetConds()),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendAnnouncements",
			"In", in,
			"Error", err,
		)
		return &npool.GetSendAnnouncementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetSendAnnouncements(ctx)
	if err != nil {
		return &npool.GetSendAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSendAnnouncementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetSendAnnouncement(ctx context.Context, in *npool.GetSendAnnouncementRequest) (*npool.GetSendAnnouncementResponse, error) {
	handler, err := announcement1.NewHandler(
		ctx,
		handler.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendAnnouncement",
			"In", in,
			"error", err,
		)
		return &npool.GetSendAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetSendAnnouncement(ctx)
	if err != nil {
		return &npool.GetSendAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSendAnnouncementResponse{
		Info: info,
	}, nil
}
