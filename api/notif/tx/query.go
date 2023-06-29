//nolint:nolintlint,dupl
package tx

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"

	tx1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/tx"
)

func (s *Server) GetTxs(ctx context.Context, in *npool.GetTxsRequest) (*npool.GetTxsResponse, error) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithConds(in.GetConds()),
		tx1.WithOffset(in.Offset),
		tx1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTxs",
			"In", in,
			"Error", err,
		)
		return &npool.GetTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetTxs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTxs",
			"In", in,
			"Error", err,
		)
		return &npool.GetTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetTx(ctx context.Context, in *npool.GetTxRequest) (*npool.GetTxResponse, error) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTx",
			"In", in,
			"error", err,
		)
		return &npool.GetTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTxs",
			"In", in,
			"Error", err,
		)
		return &npool.GetTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTxOnly(ctx context.Context, in *npool.GetTxOnlyRequest) (*npool.GetTxOnlyResponse, error) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTxOnly",
			"In", in,
			"error", err,
		)
		return &npool.GetTxOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetTxOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTxOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetTxOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxOnlyResponse{
		Info: info,
	}, nil
}
