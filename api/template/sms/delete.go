package sms

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
	smstemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/sms"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteSMSTemplate(ctx context.Context, in *npool.DeleteSMSTemplateRequest) (*npool.DeleteSMSTemplateResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := smstemplate1.NewHandler(
		ctx,
		smstemplate1.WithID(&id, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteSMSTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteSMSTemplateResponse{
		Info: info,
	}, nil
}
