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

func (s *Server) UpdateTx(ctx context.Context, in *npool.UpdateTxRequest) (*npool.UpdateTxResponse, error) {
	req := in.GetInfo()
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithID(req.ID),
		tx1.WithNotifState(req.NotifState),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTx",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateTxResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTx",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateTxResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTxResponse{
		Info: info,
	}, nil
}
