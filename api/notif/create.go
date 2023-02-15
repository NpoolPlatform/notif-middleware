//nolint:nolintlint,dupl
package notif

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

//nolint:gocyclo
func ValidateCreate(in *mgrpb.NotifReq) error {
	if in == nil {
		logger.Sugar().Errorw("validate", "in", in, "error", "invalid info")
		return fmt.Errorf("invalid info")
	}
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", in.GetUserID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "LangID", in.GetLangID(), "error", err)
		return err
	}

	switch in.GetEventType() {
	case basetypes.UsedFor_WithdrawalRequest:
	case basetypes.UsedFor_WithdrawalCompleted:
	case basetypes.UsedFor_DepositReceived:
	case basetypes.UsedFor_KYCApproved:
	case basetypes.UsedFor_KYCRejected:
	case basetypes.UsedFor_Announcement:
	default:
		return fmt.Errorf("EventType is invalid")
	}

	switch in.GetChannel() {
	case channel.NotifChannel_ChannelFrontend:
	case channel.NotifChannel_ChannelEmail:
	case channel.NotifChannel_ChannelSMS:
	default:
		logger.Sugar().Errorw("validate", "Channel", in.GetChannel(), "error", "invalid channel")
		return fmt.Errorf("channel is invalid")
	}

	if in.GetTitle() == "" {
		logger.Sugar().Errorw("validate", "Title", in.GetTitle())
		return fmt.Errorf("title is invalid")
	}
	if in.GetContent() == "" {
		logger.Sugar().Errorw("validate", "Content", in.GetContent())
		return fmt.Errorf("title is invalid")
	}
	return nil
}

func (s *Server) CreateNotif(ctx context.Context, in *npool.CreateNotifRequest) (*npool.CreateNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "notif", "crud", "Create")

	if err = ValidateCreate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateNotif", "error", err)
		return &npool.CreateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := notif1.CreateNotif(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateNotif", "error", err)
		return &npool.CreateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateNotifs(ctx context.Context, in *npool.CreateNotifsRequest) (*npool.CreateNotifsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotifs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "notif", "crud", "Create")

	for _, val := range in.GetInfos() {
		if err = ValidateCreate(val); err != nil {
			logger.Sugar().Errorw("CreateNotifs", "error", err)
			return &npool.CreateNotifsResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	infos, err := notif1.CreateNotifs(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateNotifs", "error", err)
		return &npool.CreateNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifsResponse{
		Infos: infos,
	}, nil
}
