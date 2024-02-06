//nolint:nolintlint,dupl
package notif

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif"
)

func (s *Server) GenerateNotifs(ctx context.Context, in *npool.GenerateNotifsRequest) (*npool.GenerateNotifsResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithAppID(&in.AppID, true),
		notif1.WithUserID(in.UserID, false),
		notif1.WithEventType(&in.EventType, true),
		notif1.WithExtra(in.Extra, false),
		notif1.WithVars(in.Vars, false),
		notif1.WithNotifType(&in.NotifType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GenerateNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.GenerateNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.GenerateNotifs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GenerateNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.GenerateNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.GenerateNotifsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GenerateMultiNotifs(ctx context.Context, in *npool.GenerateMultiNotifsRequest) (*npool.GenerateMultiNotifsResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithAppID(&in.AppID, true),
		notif1.WithMultiNotifReqs(in.Reqs, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GenerateMultiNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.GenerateMultiNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.GenerateMultiNotifs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GenerateMultiNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.GenerateMultiNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.GenerateMultiNotifsResponse{
		Infos: infos,
	}, nil
}
