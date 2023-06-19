//nolint:nolintlint,dupl
package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	readstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/readstate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetReadState(ctx context.Context, in *npool.GetReadStateRequest) (*npool.GetReadStateResponse, error) {
	handler, err := readstate1.NewHandler(
		ctx,
		readstate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadState",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetReadState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadState",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetReadStateResponse{
		Info: info,
	}, nil
}

func (s *Server) GetReadStateOnly(ctx context.Context, in *npool.GetReadStateOnlyRequest) (*npool.GetReadStateOnlyResponse, error) {
	handler, err := readstate1.NewHandler(
		ctx,
		readstate1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadStateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadStateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetReadStateOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadStateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadStateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetReadStateOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetReadStates(ctx context.Context, in *npool.GetReadStatesRequest) (*npool.GetReadStatesResponse, error) {
	handler, err := readstate1.NewHandler(
		ctx,
		readstate1.WithConds(in.GetConds()),
		readstate1.WithOffset(in.GetOffset()),
		readstate1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadStates",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetReadStates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadStates",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetReadStatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
