package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"
	amtreadstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/readstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistReadStateConds(
	ctx context.Context,
	in *npool.ExistReadStateCondsRequest,
) (
	*npool.ExistReadStateCondsResponse,
	error,
) {
	handler, err := amtreadstate1.NewHandler(ctx,
		amtreadstate1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistReadStateConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistReadStateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistReadStateConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistReadStateConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistReadStateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistReadStateCondsResponse{
		Info: info,
	}, nil
}
