package sms

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
	smstemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/sms"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateSMSTemplate(ctx context.Context, in *npool.UpdateSMSTemplateRequest) (*npool.UpdateSMSTemplateResponse, error) {
	req := in.GetInfo()
	handler, err := smstemplate1.NewHandler(
		ctx,
		smstemplate1.WithID(req.ID),
		smstemplate1.WithAppID(req.AppID),
		smstemplate1.WithLangID(req.LangID),
		smstemplate1.WithUsedFor(req.UsedFor),
		smstemplate1.WithSubject(req.Subject),
		smstemplate1.WithMessage(req.Message),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateSMSTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateSMSTemplateResponse{
		Info: info,
	}, nil
}
