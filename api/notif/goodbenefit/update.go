//nolint:nolintlint,dupl
package goodbenefit

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	goodbenefit1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/goodbenefit"
)

func (s *Server) UpdateGoodBenefit(
	ctx context.Context,
	in *npool.UpdateGoodBenefitRequest,
) (
	*npool.UpdateGoodBenefitResponse,
	error,
) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateGoodBenefit",
			"In", in,
		)
		return &npool.UpdateGoodBenefitResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithID(req.ID, true),
		goodbenefit1.WithGenerated(req.Generated, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodBenefit",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateGoodBenefitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGoodBenefit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodBenefit",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateGoodBenefitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateGoodBenefitResponse{
		Info: info,
	}, nil
}
