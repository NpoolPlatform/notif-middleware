//nolint:nolintlint,dupl
package sms

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
	smstemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/sms"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSMSTemplate(ctx context.Context, in *npool.CreateSMSTemplateRequest) (*npool.CreateSMSTemplateResponse, error) {
	req := in.GetInfo()
	handler, err := smstemplate1.NewHandler(
		ctx,
		smstemplate1.WithID(req.ID),
		smstemplate1.WithAppID(req.AppID),
		smstemplate1.WithLangID(req.LangID),
		smstemplate1.WithSubject(req.Subject),
		smstemplate1.WithUsedFor(req.UsedFor),
		smstemplate1.WithMessage(req.Message),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateSMSTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateSMSTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateSMSTemplates(ctx context.Context, in *npool.CreateSMSTemplatesRequest) (*npool.CreateSMSTemplatesResponse, error) {
	handler, err := smstemplate1.NewHandler(
		ctx,
		smstemplate1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSMSTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSMSTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateSMSTemplates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSMSTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSMSTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateSMSTemplatesResponse{
		Infos: infos,
	}, nil
}
