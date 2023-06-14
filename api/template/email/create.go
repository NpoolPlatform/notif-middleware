package email

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	emailtemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/email"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEmailTemplate(ctx context.Context, in *npool.CreateEmailTemplateRequest) (*npool.CreateEmailTemplateResponse, error) {
	req := in.GetInfo()
	handler, err := emailtemplate1.NewHandler(
		ctx,
		emailtemplate1.WithID(req.ID),
		emailtemplate1.WithAppID(req.AppID),
		emailtemplate1.WithLangID(req.LangID),
		emailtemplate1.WithDefaultToUsername(req.DefaultToUsername),
		emailtemplate1.WithUsedFor(req.UsedFor),
		emailtemplate1.WithSender(req.Sender),
		emailtemplate1.WithReplyTos(&req.ReplyTos),
		emailtemplate1.WithCcTos(&req.CCTos),
		emailtemplate1.WithSubject(req.Subject),
		emailtemplate1.WithBody(req.Body),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateEmailTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateEmailTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateEmailTemplates(ctx context.Context, in *npool.CreateEmailTemplatesRequest) (*npool.CreateEmailTemplatesResponse, error) {
	handler, err := emailtemplate1.NewHandler(
		ctx,
		emailtemplate1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEmailTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateEmailTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateEmailTemplates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEmailTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateEmailTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateEmailTemplatesResponse{
		Infos: infos,
	}, nil
}
