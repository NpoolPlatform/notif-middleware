//nolint:nolintlint,dupl
package txnotifstate

import (
	"context"

	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif/txnotifstate"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/txnotifstate"
)

func (s *Server) CreateTxNotifState(
	ctx context.Context,
	in *npool.CreateTxNotifStateRequest,
) (
	*npool.CreateTxNotifStateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	row, err := mgrcli.CreateTxNotifState(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.CreateTxNotifStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxNotifStateResponse{
		Info: row,
	}, nil
}

func (s *Server) UpdateTxNotifState(
	ctx context.Context,
	in *npool.UpdateTxNotifStateRequest,
) (
	*npool.UpdateTxNotifStateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	row, err := mgrcli.UpdateTxNotifState(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.UpdateTxNotifStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTxNotifStateResponse{
		Info: row,
	}, nil
}

func (s *Server) GetTxNotifStates(
	ctx context.Context,
	in *npool.GetTxNotifStatesRequest,
) (
	*npool.GetTxNotifStatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	rows, total, err := mgrcli.GetTxNotifStates(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetTxNotifStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxNotifStatesResponse{
		Infos: rows,
		Total: total,
	}, nil
}

func (s *Server) GetTxNotifStateOnly(
	ctx context.Context,
	in *npool.GetTxNotifStateOnlyRequest,
) (
	*npool.GetTxNotifStateOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	row, err := mgrcli.GetTxNotifStateOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetTxNotifStateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxNotifStateOnlyResponse{
		Info: row,
	}, nil
}
