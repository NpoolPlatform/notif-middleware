//nolint:nolintlint,dupl
package goodbenefit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	goodbenefit1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/goodbenefit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodBenefit(
	ctx context.Context,
	in *npool.CreateGoodBenefitRequest,
) (
	*npool.CreateGoodBenefitResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithID(req.ID),
		goodbenefit1.WithGoodID(req.GoodID),
		goodbenefit1.WithGoodName(req.GoodName),
		goodbenefit1.WithAmount(req.Amount),
		goodbenefit1.WithState(req.State),
		goodbenefit1.WithMessage(req.Message),
		goodbenefit1.WithTxID(req.TxID),
		goodbenefit1.WithNotified(req.Notified),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodBenefit",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodBenefitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodBenefit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodBenefit",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodBenefitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodBenefitResponse{
		Info: info,
	}, nil
}
