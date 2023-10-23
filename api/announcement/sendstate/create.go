package sendstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	handler1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	sendamt "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/sendstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSendState(ctx context.Context, in *npool.CreateSendStateRequest) (*npool.CreateSendStateResponse, error) {
	req := in.GetInfo()
	handler, err := sendamt.NewHandler(
		ctx,
		handler1.WithAppID(req.AppID, true),
		handler1.WithUserID(req.UserID, true),
		handler1.WithAnnouncementID(req.AppID, req.AnnouncementID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendState",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSendState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendState",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSendStateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateSendStates(ctx context.Context, in *npool.CreateSendStatesRequest) (*npool.CreateSendStatesResponse, error) {
	handler, err := sendamt.NewHandler(
		ctx,
		sendamt.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendStates",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateSendStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.CreateSendStates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendStates",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateSendStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSendStatesResponse{
		Infos: infos,
	}, nil
}
