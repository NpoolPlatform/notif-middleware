package sms

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
	smstemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/sms"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) ExistSMSTemplate(ctx context.Context, in *npool.ExistSMSTemplateRequest) (*npool.ExistSMSTemplateResponse, error) {
	handler, err := smstemplate1.NewHandler(
		ctx,
		smstemplate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistSMSTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSMSTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSMSTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistSMSTemplateResponse{
		Info: exist,
	}, nil
}
