package goodbenefit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	goodbenefit1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/goodbenefit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteGoodBenefit(
	ctx context.Context,
	in *npool.DeleteGoodBenefitRequest,
) (
	*npool.DeleteGoodBenefitResponse,
	error,
) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteGoodBenefit",
			"In", in,
		)
		return &npool.DeleteGoodBenefitResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodBenefit",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodBenefitResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteGoodBenefit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodBenefit",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodBenefitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteGoodBenefitResponse{
		Info: info,
	}, nil
}
