package notif

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteNotif(ctx context.Context, in *npool.DeleteNotifRequest) (*npool.DeleteNotifResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteNotif",
			"In", in,
		)
		return &npool.DeleteNotifResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteNotif",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteNotif(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteNotif",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteNotifResponse{
		Info: info,
	}, nil
}
