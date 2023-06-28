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
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithID(req.ID),
		goodbenefit1.WithNotified(req.Notified),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodBenefit",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateGoodBenefitResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateGoodBenefit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodBenefit",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateGoodBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateGoodBenefitResponse{
		Info: info,
	}, nil
}
