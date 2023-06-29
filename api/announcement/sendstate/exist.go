package sendstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	amtsendstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/sendstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) ExistSendStateConds(
	ctx context.Context,
	in *npool.ExistSendStateCondsRequest,
) (
	*npool.ExistSendStateCondsResponse,
	error,
) {
	handler, err := amtsendstate1.NewHandler(ctx,
		amtsendstate1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSendStateConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistSendStateCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistSendStateConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSendStateConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistSendStateCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistSendStateCondsResponse{
		Info: info,
	}, nil
}
