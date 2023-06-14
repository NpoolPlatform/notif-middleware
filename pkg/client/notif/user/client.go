package user

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"

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

func CreateUser(ctx context.Context, req *npool.UserReq) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateUser(ctx, &npool.CreateUserRequest{
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
	return info.(*npool.User), nil
}

func CreateUsers(ctx context.Context, reqs []*npool.UserReq) ([]*npool.User, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateUsers(ctx, &npool.CreateUsersRequest{
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
	return infos.([]*npool.User), nil
}

func UpdateUser(ctx context.Context, req *npool.UserReq) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateUser(ctx, &npool.UpdateUserRequest{
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
	return info.(*npool.User), nil
}

func GetUser(ctx context.Context, id string) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUser(ctx, &npool.GetUserRequest{
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
	return info.(*npool.User), nil
}

func GetUsers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.User, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUsers(ctx, &npool.GetUsersRequest{
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
	return infos.([]*npool.User), total, nil
}

func GetUserOnly(ctx context.Context, conds *npool.Conds) (*npool.User, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUsers(ctx, &npool.GetUsersRequest{
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
	if len(infos.([]*npool.User)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.User)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.User)[0], nil
}

func DeleteUser(ctx context.Context, req *npool.UserReq) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteUser(ctx, &npool.DeleteUserRequest{
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
	return info.(*npool.User), nil
}

func ExistUserConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistUserConds(ctx, &npool.ExistUserCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get user: %v", err)
	}
	return infos.(bool), nil
}
