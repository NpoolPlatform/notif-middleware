package template

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	template1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GenerateText(ctx context.Context, in *npool.GenerateTextRequest) (*npool.GenerateTextResponse, error) { //nolint
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.GenerateTextResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetLangID()); err != nil {
		return &npool.GenerateTextResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	switch in.GetEventType() {
	case basetypes.UsedFor_Signup:
	case basetypes.UsedFor_Signin:
	case basetypes.UsedFor_Update:
	case basetypes.UsedFor_SetWithdrawAddress:
	case basetypes.UsedFor_Withdraw:
	case basetypes.UsedFor_CreateInvitationCode:
	case basetypes.UsedFor_SetCommission:
	case basetypes.UsedFor_SetTransferTargetUser:
	case basetypes.UsedFor_Transfer:
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
	case basetypes.NotifChannel_ChannelFrontend:
	case basetypes.NotifChannel_ChannelEmail:
	case basetypes.NotifChannel_ChannelSMS:
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
