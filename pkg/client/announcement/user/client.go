package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/notif-middleware/pkg/servicename"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
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

func CreateAnnouncementUser(ctx context.Context, in *npool.AnnouncementUserReq) (*npool.AnnouncementUser, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateAnnouncementUser(ctx, &npool.CreateAnnouncementUserRequest{
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
	return info.(*npool.AnnouncementUser), nil
}

func DeleteAnnouncementUser(ctx context.Context, id uint32) (*npool.AnnouncementUser, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteAnnouncementUser(ctx, &npool.DeleteAnnouncementUserRequest{
			Info: &npool.AnnouncementUserReq{
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
	return info.(*npool.AnnouncementUser), nil
}

func GetAnnouncementUsers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.AnnouncementUser, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncementUsers(ctx, &npool.GetAnnouncementUsersRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcement users: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get announcement users: %v", err)
	}
	return infos.([]*npool.AnnouncementUser), total, nil
}

func GetAnnouncementUser(ctx context.Context, appID, id string) (*npool.AnnouncementUser, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncementUser(ctx, &npool.GetAnnouncementUserRequest{
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
	return info.(*npool.AnnouncementUser), nil
}

func ExistAnnouncementUserConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistAnnouncementUserConds(ctx, &npool.ExistAnnouncementUserCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail exist announcement user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("ail exist announcement user: %v", err)
	}

	return info.(bool), nil
}
