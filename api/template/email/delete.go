package email

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	emailtemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/email"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteEmailTemplate(
	ctx context.Context,
	in *npool.DeleteEmailTemplateRequest,
) (
	*npool.DeleteEmailTemplateResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := emailtemplate1.NewHandler(
		ctx,
		emailtemplate1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteEmailTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteEmailTemplateResponse{
		Info: info,
	}, nil
}
