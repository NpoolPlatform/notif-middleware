//nolint:nolintlint,dupl
package sendstate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"
	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/announcement/sendstate"

	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"

	sendstate1 "github.com/NpoolPlatform/notif-middleware/pkg/announcement/sendstate"
	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"

	"github.com/google/uuid"
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

func (s *Server) GetSendStates(ctx context.Context, in *npool.GetSendStatesRequest) (*npool.GetSendStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSendStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err = validateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("validateConds", "Conds", in.GetConds(), "error", err)
		return &npool.GetSendStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Rows")

	rows, total, err := sendstate1.GetSendStates(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetSendStates", "error", err)
		return &npool.GetSendStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSendStatesResponse{
		Infos: rows,
		Total: total,
	}, nil
}

func (s *Server) CreateSendState(ctx context.Context, in *npool.CreateSendStateRequest) (*npool.CreateSendStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSendStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateSendState", "AppID", in.GetAppID(), "error", err)
		return &npool.CreateSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("CreateSendState", "UserID", in.GetUserID(), "error", err)
		return &npool.CreateSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetAnnouncementID()); err != nil {
		logger.Sugar().Errorw("CreateSendState", "AnnouncementID", in.GetAnnouncementID(), "error", err)
		return &npool.CreateSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	switch in.GetChannel() {
	case channel.NotifChannel_ChannelEmail:
	case channel.NotifChannel_ChannelSMS:
	default:
		logger.Sugar().Errorw("CreateSendState", "Channel", in.GetChannel())
		return &npool.CreateSendStateResponse{}, status.Error(codes.InvalidArgument, "channel is invalid")
	}
	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Rows")

	exist, err := mgrcli.ExistSendStateConds(ctx, &mgrpb.Conds{
		AppID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		UserID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: in.GetUserID(),
		},
		AnnouncementID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAnnouncementID(),
		},
		Channel: &valuedef.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(in.GetChannel().Number()),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateSendState", "err", err.Error())
		return &npool.CreateSendStateResponse{}, status.Error(codes.Internal, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("CreateSendState", "err", "send state already exist")
		return &npool.CreateSendStateResponse{}, status.Error(codes.InvalidArgument, "send state already exist")
	}

	err = sendstate1.CreateSendState(
		ctx,
		in.GetAppID(),
		in.GetUserID(),
		in.GetAnnouncementID(),
		in.GetChannel(),
	)
	if err != nil {
		logger.Sugar().Errorw("GetSendStates", "error", err)
		return &npool.CreateSendStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSendStateResponse{}, nil
}

func (s *Server) CreateSendStates(ctx context.Context, in *npool.CreateSendStatesRequest) (*npool.CreateSendStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSendStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	for _, info := range in.GetInfos() {
		if _, err := uuid.Parse(info.GetAppID()); err != nil {
			logger.Sugar().Errorw("CreateSendState", "AppID", info.GetAppID(), "error", err)
			return &npool.CreateSendStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		if _, err := uuid.Parse(info.GetUserID()); err != nil {
			logger.Sugar().Errorw("CreateSendState", "UserID", info.GetUserID(), "error", err)
			return &npool.CreateSendStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		if _, err := uuid.Parse(info.GetAnnouncementID()); err != nil {
			logger.Sugar().Errorw("CreateSendState", "AnnouncementID", info.GetAnnouncementID(), "error", err)
			return &npool.CreateSendStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		switch info.GetChannel() {
		case channel.NotifChannel_ChannelEmail:
		case channel.NotifChannel_ChannelSMS:
		default:
			logger.Sugar().Errorw("CreateSendState", "Channel", info.GetChannel())
			return &npool.CreateSendStatesResponse{}, status.Error(codes.InvalidArgument, "channel is invalid")
		}
		span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Rows")
	}
	err = sendstate1.CreateSendStates(
		ctx,
		in.GetInfos(),
	)
	if err != nil {
		logger.Sugar().Errorw("GetSendStates", "error", err)
		return &npool.CreateSendStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSendStatesResponse{}, nil
}
