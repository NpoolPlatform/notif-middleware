//nolint:nolintlint,dupl
package notif

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateNotif(ctx context.Context, in *npool.CreateNotifRequest) (*npool.CreateNotifResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateNotif",
			"In", in,
		)
		return &npool.CreateNotifResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithEntID(req.EntID, false),
		notif1.WithAppID(req.AppID, true),
		notif1.WithLangID(req.LangID, true),
		notif1.WithUserID(req.UserID, false),
		notif1.WithEventID(req.EventID, true),
		notif1.WithNotified(req.Notified, false),
		notif1.WithEventType(req.EventType, true),
		notif1.WithUseTemplate(req.UseTemplate, false),
		notif1.WithTitle(req.Title, true),
		notif1.WithContent(req.Content, true),
		notif1.WithChannel(req.Channel, true),
		notif1.WithExtra(req.Extra, false),
		notif1.WithNotifType(req.NotifType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNotif",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateNotif(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNotif",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateNotifResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateNotifs(ctx context.Context, in *npool.CreateNotifsRequest) (*npool.CreateNotifsResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateNotifs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateNotifsResponse{
		Infos: infos,
	}, nil
}
