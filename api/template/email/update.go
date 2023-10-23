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

func (s *Server) UpdateEmailTemplate(
	ctx context.Context,
	in *npool.UpdateEmailTemplateRequest,
) (
	*npool.UpdateEmailTemplateResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := emailtemplate1.NewHandler(
		ctx,
		emailtemplate1.WithID(req.ID, true),
		emailtemplate1.WithAppID(req.AppID, false),
		emailtemplate1.WithUsedFor(req.UsedFor, false),
		emailtemplate1.WithDefaultToUsername(req.DefaultToUsername, false),
		emailtemplate1.WithSender(req.Sender, false),
		emailtemplate1.WithReplyTos(&req.ReplyTos, false),
		emailtemplate1.WithCcTos(&req.CCTos, false),
		emailtemplate1.WithSubject(req.Subject, false),
		emailtemplate1.WithBody(req.Body, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateEmailTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateEmailTemplateResponse{
		Info: info,
	}, nil
}
