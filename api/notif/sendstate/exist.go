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

func (s *Server) ExistSendState(ctx context.Context, in *npool.ExistSendStateRequest) (*npool.ExistSendStateResponse, error) {
	handler, err := sendstate1.NewHandler(
		ctx,
		sendstate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSendState",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistSendState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSendState",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistSendStateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSendStateConds(
	ctx context.Context,
	in *npool.ExistSendStateCondsRequest,
) (
	*npool.ExistSendStateCondsResponse,
	error,
) {
	handler, err := sendstate1.NewHandler(ctx,
		sendstate1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSendState",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistSendStateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistSendStateConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSendState",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistSendStateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSendStateCondsResponse{
		Info: info,
	}, nil
}
