package sendstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"
	sendstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/sendstate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateSendState(ctx context.Context, in *npool.UpdateSendStateRequest) (*npool.UpdateSendStateResponse, error) {
	req := in.GetInfo()
	handler, err := sendstate1.NewHandler(
		ctx,
		sendstate1.WithID(req.ID),
		sendstate1.WithAppID(req.AppID),
		sendstate1.WithUserID(req.UserID),
		sendstate1.WithNotifID(req.NotifID),
		sendstate1.WithChannel(req.Channel),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSendState",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateSendState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSendState",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateSendStateResponse{
		Info: info,
	}, nil
}
