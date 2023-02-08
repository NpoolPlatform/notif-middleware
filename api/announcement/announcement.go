//nolint:nolintlint,dupl
package announcement

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/google/uuid"

	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/announcement"

	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
)

//nolint:gocyclo
func validateConds(in *npool.Conds) error {
	if in == nil {
		return nil
	}
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "ID", in.GetID().GetValue(), "error", err)
			return err
		}
	}
	if in.AppID != nil {
		if _, err := uuid.Parse(in.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "AppID", in.GetAppID().GetValue(), "error", err)
			return err
		}
	}
	if in.UserID != nil {
		if _, err := uuid.Parse(in.GetUserID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "UserID", in.GetUserID().GetValue(), "error", err)
			return err
		}
	}
	if in.AnnouncementID != nil {
		if _, err := uuid.Parse(in.GetAnnouncementID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "AnnouncementID", in.GetAnnouncementID().GetValue(), "error", err)
			return err
		}
	}
	if in.Channel != nil {
		switch in.GetChannel().GetValue() {
		case uint32(channel.NotifChannel_ChannelEmail):
		case uint32(channel.NotifChannel_ChannelSMS):
		default:
			return fmt.Errorf("channel is invalid")
		}
	}
	if in.UserIDs != nil {
		for _, v := range in.GetUserIDs().GetValue() {
			if _, err := uuid.Parse(v); err != nil {
				logger.Sugar().Errorw("validateConds", "UserID", v, "error", err)
				return err
			}
		}
	}
	return nil
}

func (s *Server) GetAnnouncements(ctx context.Context, in *npool.GetAnnouncementsRequest) (*npool.GetAnnouncementsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAnnouncements")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	span = commontracer.TraceInvoker(span, "announcement/announcement", "crud", "Rows")

	err = validateConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAnnouncements", "error", err)
		return &npool.GetAnnouncementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	rows, total, err := announcement1.GetAnnouncements(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetAnnouncements", "error", err)
		return &npool.GetAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementsResponse{
		Infos: rows,
		Total: total,
	}, nil
}
