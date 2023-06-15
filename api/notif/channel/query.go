//nolint:nolintlint,dupl
package channel

import (
	"context"

	channel1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/channel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
)

func (s *Server) GetChannels(ctx context.Context, in *npool.GetChannelsRequest) (*npool.GetChannelsResponse, error) {
	handler, err := channel1.NewHandler(
		ctx,
		channel1.WithConds(in.GetConds()),
		channel1.WithOffset(in.Offset),
		channel1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetChannels",
			"In", in,
			"Error", err,
		)
		return &npool.GetChannelsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetChannels(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetChannels",
			"Req", in,
			"Error", err,
		)
		return &npool.GetChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChannelsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
