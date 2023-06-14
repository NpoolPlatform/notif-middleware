//nolint:nolintlint,dupl
package notif


import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetNotif(ctx context.Context, in *npool.GetNotifRequest) (*npool.GetNotifResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotif",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetNotif(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotif",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotifOnly(ctx context.Context, in *npool.GetNotifOnlyRequest) (*npool.GetNotifOnlyResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotif",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetNotif(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotif",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotifs(ctx context.Context, in *npool.GetNotifsRequest) (*npool.GetNotifsResponse, error) {
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithConds(in.GetConds()),
		notif1.WithOffset(in.GetOffset()),
		notif1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetNotifs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNotifs",
			"In", in,
			"Error", err,
		)
		return &npool.GetNotifsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNotifsResponse{
		Infos: infos,
		Total: total,
	}, nil
}



import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	channel "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/notif"
)

func (s *Server) GetNotif(ctx context.Context, in *npool.GetNotifRequest) (*npool.GetNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	_, err = uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetNotif", "ID", in.GetID(), "error", err)
		return &npool.GetNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Row")

	info, err := notif1.GetNotif(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetNotif", "error", err)
		return &npool.GetNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifResponse{
		Info: info,
	}, nil
}
