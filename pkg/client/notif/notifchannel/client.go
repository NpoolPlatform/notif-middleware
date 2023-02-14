package notifchannel

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/notifchannel"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/notifchannel"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get notifchannel connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func GetNotifChannelOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.NotifChannel, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetNotifChannelOnly(ctx, &npool.GetNotifChannelOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notifchannel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get notifchannel: %v", err)
	}
	return info.(*mgrpb.NotifChannel), nil
}

func GetNotifChannels(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.NotifChannel, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetNotifChannels(ctx, &npool.GetNotifChannelsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notifchannels: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get notifchannels: %v", err)
	}
	return infos.([]*mgrpb.NotifChannel), total, nil
}
