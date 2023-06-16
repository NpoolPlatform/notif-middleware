//nolint:nolintlint,dupl
package sendstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"
	sendstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/sendstate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSendState(ctx context.Context, in *npool.CreateSendStateRequest) (*npool.CreateSendStateResponse, error) {
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
			"CreateSendState",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateSendState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendState",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateSendStateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateSendStates(ctx context.Context, in *npool.CreateSendStatesRequest) (*npool.CreateSendStatesResponse, error) {
	handler, err := sendstate1.NewHandler(
		ctx,
		sendstate1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendStates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSendStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateSendStates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendStates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSendStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateSendStatesResponse{
		Infos: infos,
	}, nil
}
