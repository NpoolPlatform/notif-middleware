package sendstate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"

	servicename "github.com/NpoolPlatform/notif-middleware/pkg/servicename"
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

func CreateSendState(ctx context.Context, req *npool.SendStateReq) (*npool.SendState, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateSendState(ctx, &npool.CreateSendStateRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.SendState), nil
}

func CreateSendStates(ctx context.Context, reqs []*npool.SendStateReq) ([]*npool.SendState, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateSendStates(ctx, &npool.CreateSendStatesRequest{
			Infos: reqs,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.SendState), nil
}

func UpdateSendState(ctx context.Context, req *npool.SendStateReq) (*npool.SendState, error) {
	return nil, nil
}

func GetSendState(ctx context.Context, id string) (*npool.SendState, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSendState(ctx, &npool.GetSendStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.SendState), nil
}

func GetSendStates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.SendState, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSendStates(ctx, &npool.GetSendStatesRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.SendState), total, nil
}

func GetSendStateOnly(ctx context.Context, conds *npool.Conds) (*npool.SendState, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSendStates(ctx, &npool.GetSendStatesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2, //nolint
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.SendState)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.SendState)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.SendState)[0], nil
}

func DeleteSendState(ctx context.Context, req *npool.SendStateReq) (*npool.SendState, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteSendState(ctx, &npool.DeleteSendStateRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.SendState), nil
}

func ExistSendStateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistSendStateConds(ctx, &npool.ExistSendStateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get sendstate: %v", err)
	}
	return infos.(bool), nil
}
