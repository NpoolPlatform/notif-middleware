package contact

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	contact1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/contact"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteContact(ctx context.Context, in *npool.DeleteContactRequest) (*npool.DeleteContactResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteContact",
			"In", in,
		)
		return &npool.DeleteContactResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := contact1.NewHandler(
		ctx,
		contact1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteContact",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteContactResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteContact(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteContact",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteContactResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteContactResponse{
		Info: info,
	}, nil
}
