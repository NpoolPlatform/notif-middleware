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

func (s *Server) GetSendState(ctx context.Context, in *npool.GetSendStateRequest) (*npool.GetSendStateResponse, error) {
	handler, err := sendstate1.NewHandler(
		ctx,
		sendstate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendState",
			"In", in,
			"Error", err,
		)
		return &npool.GetSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetSendState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendState",
			"In", in,
			"Error", err,
		)
		return &npool.GetSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSendStateResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSendStateOnly(ctx context.Context, in *npool.GetSendStateOnlyRequest) (*npool.GetSendStateOnlyResponse, error) {
	handler, err := sendstate1.NewHandler(
		ctx,
		sendstate1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendStateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetSendStateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetSendStateOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendStateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetSendStateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSendStateOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSendStates(ctx context.Context, in *npool.GetSendStatesRequest) (*npool.GetSendStatesResponse, error) {
	handler, err := sendstate1.NewHandler(
		ctx,
		sendstate1.WithConds(in.GetConds()),
		sendstate1.WithOffset(in.GetOffset()),
		sendstate1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendStates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSendStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetSendStates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSendStates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSendStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSendStatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
