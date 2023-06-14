package contact

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	contact1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/contact"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetContacts(ctx context.Context, in *npool.GetContactsRequest) (*npool.GetContactsResponse, error) {
	handler, err := contact1.NewHandler(
		ctx,
		contact1.WithConds(in.GetConds()),
		contact1.WithOffset(in.Offset),
		contact1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetContacts",
			"In", in,
			"Error", err,
		)
		return &npool.GetContactsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetContacts(ctx)
	if err != nil {
		return &npool.GetContactsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContactsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetContact(ctx context.Context, in *npool.GetContactRequest) (*npool.GetContactResponse, error) {
	handler, err := contact1.NewHandler(
		ctx,
		contact1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetContact",
			"In", in,
			"error", err,
		)
		return &npool.GetContactResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetContact(ctx)
	if err != nil {
		return &npool.GetContactResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContactResponse{
		Info: info,
	}, nil
}
