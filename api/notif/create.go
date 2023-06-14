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
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithID(req.ID),
		notif1.WithAppID(req.AppID),
		notif1.WithLangID(req.LangID),
		notif1.WithUserID(req.UserID),
		notif1.WithEventID(req.EventID),
		notif1.WithNotified(req.Notified),
		notif1.WithEventType(req.EventType),
		notif1.WithUseTemplate(&req.UseTemplate),
		notif1.WithTitle(req.Title),
		notif1.WithContent(req.Content),
		notif1.WithChannel(req.Channel),
		notif1.WithExtra(req.Extra),
		notif1.WithType(req.Type),
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
		notif1.WithReqs(in.GetInfos()),
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
