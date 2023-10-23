package channel

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	channel1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/channel"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateChannel(ctx context.Context, in *npool.CreateChannelRequest) (*npool.CreateChannelResponse, error) {
	req := in.GetInfo()
	if req == nil {
		return nil, fmt.Errorf("invalid arg")
	}
	handler, err := channel1.NewHandler(
		ctx,
		channel1.WithAppID(req.AppID, true),
		channel1.WithEventType(req.EventType, true),
		channel1.WithChannel(req.Channel, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateChannel",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateChannelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateChannel(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateChannel",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateChannelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateChannelResponse{
		Info: info,
	}, nil
}
