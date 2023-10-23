//nolint:nolintlint,dupl
package email

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	emailtemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/email"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetEmailTemplate(
	ctx context.Context,
	in *npool.GetEmailTemplateRequest,
) (
	*npool.GetEmailTemplateResponse,
	error,
) {
	handler, err := emailtemplate1.NewHandler(
		ctx,
		emailtemplate1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.GetEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetEmailTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.GetEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetEmailTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) GetEmailTemplateOnly(
	ctx context.Context,
	in *npool.GetEmailTemplateOnlyRequest,
) (
	*npool.GetEmailTemplateOnlyResponse,
	error,
) {
	handler, err := emailtemplate1.NewHandler(
		ctx,
		emailtemplate1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEmailTemplateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetEmailTemplateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetEmailTemplateOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEmailTemplateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetEmailTemplateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetEmailTemplateOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetEmailTemplates(
	ctx context.Context,
	in *npool.GetEmailTemplatesRequest,
) (
	*npool.GetEmailTemplatesResponse,
	error,
) {
	handler, err := emailtemplate1.NewHandler(
		ctx,
		emailtemplate1.WithConds(in.GetConds()),
		emailtemplate1.WithOffset(in.GetOffset()),
		emailtemplate1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEmailTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.GetEmailTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetEmailTemplates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEmailTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.GetEmailTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetEmailTemplatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
