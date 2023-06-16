package readstate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"

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

func CreateReadState(ctx context.Context, req *npool.ReadStateReq) (*npool.ReadState, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateReadState(ctx, &npool.CreateReadStateRequest{
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
	return info.(*npool.ReadState), nil
}

func CreateReadStates(ctx context.Context, reqs []*npool.ReadStateReq) ([]*npool.ReadState, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateReadStates(ctx, &npool.CreateReadStatesRequest{
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
	return infos.([]*npool.ReadState), nil
}

func UpdateReadState(ctx context.Context, req *npool.ReadStateReq) (*npool.ReadState, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateReadState(ctx, &npool.UpdateReadStateRequest{
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
	return info.(*npool.ReadState), nil
}

func GetReadState(ctx context.Context, id string) (*npool.ReadState, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetReadState(ctx, &npool.GetReadStateRequest{
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
	return info.(*npool.ReadState), nil
}

func GetReadStates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.ReadState, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetReadStates(ctx, &npool.GetReadStatesRequest{
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
	return infos.([]*npool.ReadState), total, nil
}

func GetReadStateOnly(ctx context.Context, conds *npool.Conds) (*npool.ReadState, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetReadStates(ctx, &npool.GetReadStatesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.ReadState)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.ReadState)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.ReadState)[0], nil
}

func DeleteReadState(ctx context.Context, req *npool.ReadStateReq) (*npool.ReadState, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteReadState(ctx, &npool.DeleteReadStateRequest{
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
	return info.(*npool.ReadState), nil
}

func ExistReadStateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistReadStateConds(ctx, &npool.ExistReadStateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get readstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get readstate: %v", err)
	}
	return infos.(bool), nil
}
