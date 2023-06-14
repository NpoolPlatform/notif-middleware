//nolint:nolintlint,dupl
package notif

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"
	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif"
	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateNotif(ctx context.Context, in *npool.UpdateNotifRequest) (*npool.UpdateNotifResponse, error) {
	req := in.GetInfo()
	handler, err := notif1.NewHandler(
		ctx,
		notif1.WithID(req.ID),
		notif1.WithAppID(req.AppID),
		notif1.WithLangID(req.LangID),
		notif1.WithUserID(req.UserID),
		notif1.WithEventID(req.EventID),
		notif1.WithNotified(req.Notified),
		notif1.WithEventType(req.EventType),
		notif1.WithUseTemplate(&req.UseTemplate),
		notif1.WithTitle(req.Title),
		notif1.WithContent(req.Content),
		notif1.WithChannel(req.Channel),
		notif1.WithExtra(req.Extra),
		notif1.WithType(req.Type),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNotif",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateNotif(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNotif",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNotifResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateNotifResponse{
		Info: info,
	}, nil
}

//nolint:funlen,gocyclo
func (s *Server) UpdateNotifs(ctx context.Context, in *npool.UpdateNotifsRequest) (*npool.UpdateNotifsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdatesNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "notif", "crud", "Updates")

	if len(in.GetIDs()) == 0 {
		return &npool.UpdateNotifsResponse{}, status.Error(codes.Internal, "IDs is invalid")
	}

	for _, id := range in.GetIDs() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("validate", "ID", id, "error", err)
			return &npool.UpdateNotifsResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	infos, _, err := notif1.UpdateNotifs(ctx, in.GetIDs(), in.Notified)
	if err != nil {
		logger.Sugar().Errorw("UpdatesNotif", "error", err)
		return &npool.UpdateNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateNotifsResponse{
		Infos: infos,
	}, nil
}
