package contact

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	contact1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/contact"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateContact(ctx context.Context, in *npool.UpdateContactRequest) (*npool.UpdateContactResponse, error) {
	req := in.GetInfo()
	handler, err := contact1.NewHandler(
		ctx,
		contact1.WithID(req.ID),
		contact1.WithSender(req.Sender),
		contact1.WithAccount(req.Account),
		contact1.WithAccountType(req.AccountType),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateContact",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateContactResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateContact(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateContact",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateContactResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateContactResponse{
		Info: info,
	}, nil
}
