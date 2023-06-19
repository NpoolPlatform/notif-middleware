//nolint:nolintlint,dupl
package notif

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistNotifConds(
	ctx context.Context,
	in *npool.ExistNotifCondsRequest,
) (
	*npool.ExistNotifCondsResponse,
	error,
) {
	handler, err := notif1.NewHandler(ctx,
		notif1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistNotif",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistNotifCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistNotifConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistNotif",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistNotifCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistNotifCondsResponse{
		Info: info,
	}, nil
}
