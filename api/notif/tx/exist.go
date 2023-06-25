package tx

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
	tx1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/tx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) ExistTx(
	ctx context.Context,
	in *npool.ExistTxRequest,
) (
	*npool.ExistTxResponse,
	error,
) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTx",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTx",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTxConds(
	ctx context.Context,
	in *npool.ExistTxCondsRequest,
) (
	*npool.ExistTxCondsResponse,
	error,
) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTxConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistTxCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistTxConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTxConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistTxCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxCondsResponse{
		Info: info,
	}, nil
}
