//nolint:nolintlint,dupl
package tx

import (
	"context"

	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif/tx"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
)

func (s *Server) GetTxs(
	ctx context.Context,
	in *npool.GetTxsRequest,
) (
	*npool.GetTxsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	rows, total, err := mgrcli.GetTxs(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetTxs", "error", err)
		return &npool.GetTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxsResponse{
		Infos: rows,
		Total: total,
	}, nil
}

func (s *Server) GetTxOnly(
	ctx context.Context,
	in *npool.GetTxOnlyRequest,
) (
	*npool.GetTxOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	row, err := mgrcli.GetTxOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetTxOnly", "error", err)
		return &npool.GetTxOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxOnlyResponse{
		Info: row,
	}, nil
}
