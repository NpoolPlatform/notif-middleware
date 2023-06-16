package sendstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	amtsend1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/sendstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetSendStates(ctx context.Context, in *npool.GetSendStatesRequest) (*npool.GetSendStatesResponse, error) {
	handler, err := amtsend1.NewHandler(
		ctx,
		amtsend1.WithConds(in.GetConds()),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendStates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSendStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetSendStates(ctx)
	if err != nil {
		return &npool.GetSendStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSendStatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetSendState(ctx context.Context, in *npool.GetSendStateRequest) (*npool.GetSendStateResponse, error) {
	handler, err := amtsend1.NewHandler(
		ctx,
		handler.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendState",
			"In", in,
			"error", err,
		)
		return &npool.GetSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetSendState(ctx)
	if err != nil {
		return &npool.GetSendStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSendStateResponse{
		Info: info,
	}, nil
}
