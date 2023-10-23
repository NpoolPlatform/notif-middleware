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

func (s *Server) GetNotif(ctx context.Context, in *npool.GetNotifRequest) (*npool.GetNotifResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotif",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetNotif(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotif",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotifOnly(ctx context.Context, in *npool.GetNotifOnlyRequest) (*npool.GetNotifOnlyResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetNotifOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotifs(ctx context.Context, in *npool.GetNotifsRequest) (*npool.GetNotifsResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithConds(in.GetConds()),
		notif1.WithOffset(in.GetOffset()),
		notif1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetNotifs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
