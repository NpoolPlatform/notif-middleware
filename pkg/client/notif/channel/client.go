package channel

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/notif-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return fn(_ctx, cli)
}

func GetChannels(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Channel, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetChannels(ctx, &npool.GetChannelsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get channels: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get channels: %v", err)
	}
	return infos.([]*npool.Channel), total, nil
}

func GetChannelOnly(ctx context.Context, conds *npool.Conds) (*npool.Channel, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetChannels(ctx, &npool.GetChannelsRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Channel)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Channel)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Channel)[0], nil
}

func CreateChannel(ctx context.Context, in *npool.ChannelReq) (*npool.Channel, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateChannel(ctx, &npool.CreateChannelRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Channel), nil
}

func DeleteChannel(ctx context.Context, id uint32) (*npool.Channel, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteChannel(ctx, &npool.DeleteChannelRequest{
			Info: &npool.ChannelReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Channel), nil
}

func GetChannel(ctx context.Context, appID, id string) (*npool.Channel, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetChannel(ctx, &npool.GetChannelRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Channel), nil
}

func ExistChannelConds(ctx context.Context, conds *npool.ExistChannelCondsRequest) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistChannelConds(ctx, conds)
		if err != nil {
			return nil, fmt.Errorf("fail exist channel: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail exist channel: %v", err)
	}

	return info.(bool), nil
}
