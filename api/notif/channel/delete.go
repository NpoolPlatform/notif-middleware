package channel

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	channel1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/channel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteChannel(ctx context.Context, in *npool.DeleteChannelRequest) (*npool.DeleteChannelResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteChannel",
			"In", in,
		)
		return &npool.DeleteChannelResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := channel1.NewHandler(
		ctx,
		channel1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteChannel",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteChannelResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteChannel(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteChannel",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteChannelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteChannelResponse{
		Info: info,
	}, nil
}
