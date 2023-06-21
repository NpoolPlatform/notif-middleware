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

func (s *Server) UpdateNotif(ctx context.Context, in *npool.UpdateNotifRequest) (*npool.UpdateNotifResponse, error) {
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
		notif1.WithUseTemplate(req.UseTemplate),
		notif1.WithTitle(req.Title),
		notif1.WithContent(req.Content),
		notif1.WithChannel(req.Channel),
		notif1.WithExtra(req.Extra),
		notif1.WithNotifType(req.NotifType),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNotif",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateNotif(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNotif",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateNotifResponse{
		Info: info,
	}, nil
}

//nolint:funlen,gocyclo
func (s *Server) UpdateNotifs(ctx context.Context, in *npool.UpdateNotifsRequest) (*npool.UpdateNotifsResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.UpdateNotifs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateNotifsResponse{
		Infos: infos,
	}, nil
}
