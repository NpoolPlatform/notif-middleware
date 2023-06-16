package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	readstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/readstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) ExistReadState(ctx context.Context, in *npool.ExistReadStateRequest) (*npool.ExistReadStateResponse, error) {
	handler, err := readstate1.NewHandler(
		ctx,
		readstate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistReadState",
			"In", in,
			"Error", err,
		)
		return &npool.ExistReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistReadState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistReadState",
			"In", in,
			"Error", err,
		)
		return &npool.ExistReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistReadStateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistReadStateConds(ctx context.Context, in *npool.ExistReadStateCondsRequest) (*npool.ExistReadStateCondsResponse, error) {
	handler, err := readstate1.NewHandler(ctx,
		readstate1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistReadState",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistReadStateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistReadStateConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistReadState",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistReadStateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistReadStateCondsResponse{
		Info: info,
	}, nil
}
