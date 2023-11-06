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
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateNotif",
			"In", in,
		)
		return &npool.UpdateNotifResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithID(req.ID, true),
		notif1.WithAppID(req.AppID, false),
		notif1.WithLangID(req.LangID, false),
		notif1.WithUserID(req.UserID, false),
		notif1.WithEventID(req.EventID, false),
		notif1.WithNotified(req.Notified, true),
		notif1.WithEventType(req.EventType, false),
		notif1.WithUseTemplate(req.UseTemplate, false),
		notif1.WithTitle(req.Title, false),
		notif1.WithContent(req.Content, false),
		notif1.WithChannel(req.Channel, false),
		notif1.WithExtra(req.Extra, false),
		notif1.WithNotifType(req.NotifType, false),
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
		notif1.WithReqs(in.GetInfos(), false),
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
