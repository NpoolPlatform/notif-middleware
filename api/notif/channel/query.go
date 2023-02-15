//nolint:nolintlint,dupl
package channel

import (
	"context"

	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif/channel"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
)

func (s *Server) GetChannelOnly(
	ctx context.Context,
	in *npool.GetChannelOnlyRequest,
) (
	*npool.GetChannelOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetChannels")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	row, err := mgrcli.GetChannelOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetChannels", "error", err)
		return &npool.GetChannelOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChannelOnlyResponse{
		Info: row,
	}, nil
}

func (s *Server) GetChannels(
	ctx context.Context,
	in *npool.GetChannelsRequest,
) (
	*npool.GetChannelsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetChannels")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	rows, total, err := mgrcli.GetChannels(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetChannels", "error", err)
		return &npool.GetChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChannelsResponse{
		Infos: rows,
		Total: total,
	}, nil
}
