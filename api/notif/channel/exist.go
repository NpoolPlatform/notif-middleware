package channel

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	channel1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/channel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) ExistChannelConds(
	ctx context.Context,
	in *npool.ExistChannelCondsRequest,
) (
	*npool.ExistChannelCondsResponse,
	error,
) {
	handler, err := channel1.NewHandler(
		ctx,
		channel1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistChannelConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistChannelCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistChannelConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistChannelConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistChannelCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistChannelCondsResponse{
		Info: info,
	}, nil
}
