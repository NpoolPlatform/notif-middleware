package template

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	channel "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	template1 "github.com/NpoolPlatform/notif-middleware/pkg/template"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GenerateText(ctx context.Context, in *npool.GenerateTextRequest) (*npool.GenerateTextResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.GenerateTextResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetLangID()); err != nil {
		return &npool.GenerateTextResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	switch in.GetEventType() {
	case basetypes.UsedFor_WithdrawalRequest:
	case basetypes.UsedFor_WithdrawalCompleted:
	case basetypes.UsedFor_DepositReceived:
	case basetypes.UsedFor_KYCApproved:
	case basetypes.UsedFor_KYCRejected:
	default:
		logger.Sugar().Errorw("GenerateText", "EventType", in.GetEventType())
		return &npool.GenerateTextResponse{}, status.Error(codes.InvalidArgument, "EventType is invalid")
	}

	switch in.GetChannel() {
	case channel.NotifChannel_ChannelFrontend:
	case channel.NotifChannel_ChannelEmail:
	case channel.NotifChannel_ChannelSMS:
	default:
		logger.Sugar().Errorw("GenerateText", "Channel", in.GetChannel())
		return &npool.GenerateTextResponse{}, status.Error(codes.InvalidArgument, "Channel is invalid")
	}

	info, err := template1.GenerateText(ctx, in.GetAppID(), in.GetLangID(), in.GetEventType(), in.GetChannel(), in.Vars)
	if err != nil {
		logger.Sugar().Errorw("GenerateText", "Error", err)
		return &npool.GenerateTextResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GenerateTextResponse{
		Info: info,
	}, nil
}
