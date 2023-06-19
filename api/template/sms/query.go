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

func (s *Server) GetSMSTemplate(ctx context.Context, in *npool.GetSMSTemplateRequest) (*npool.GetSMSTemplateResponse, error) {
	handler, err := smstemplate1.NewHandler(
		ctx,
		smstemplate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.GetSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetSMSTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.GetSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSMSTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSMSTemplateOnly(ctx context.Context, in *npool.GetSMSTemplateOnlyRequest) (*npool.GetSMSTemplateOnlyResponse, error) {
	handler, err := smstemplate1.NewHandler(
		ctx,
		smstemplate1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSMSTemplateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetSMSTemplateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetSMSTemplateOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSMSTemplateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetSMSTemplateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSMSTemplateOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSMSTemplates(ctx context.Context, in *npool.GetSMSTemplatesRequest) (*npool.GetSMSTemplatesResponse, error) {
	handler, err := smstemplate1.NewHandler(
		ctx,
		smstemplate1.WithConds(in.GetConds()),
		smstemplate1.WithOffset(in.GetOffset()),
		smstemplate1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSMSTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSMSTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetSMSTemplates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSMSTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSMSTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSMSTemplatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
