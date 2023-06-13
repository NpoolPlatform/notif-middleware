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

func CreateUserAnnouncement(ctx context.Context, in *npool.UserAnnouncementReq) (*npool.UserAnnouncement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateUserAnnouncement(ctx, &npool.CreateUserAnnouncementRequest{
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
	return info.(*npool.UserAnnouncement), nil
}

func DeleteUserAnnouncement(ctx context.Context, id string) (*npool.UserAnnouncement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteUserAnnouncement(ctx, &npool.DeleteUserAnnouncementRequest{
			Info: &npool.UserAnnouncementReq{
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
	return info.(*npool.UserAnnouncement), nil
}

func GetUserAnnouncements(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.UserAnnouncement, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUserAnnouncements(ctx, &npool.GetUserAnnouncementsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcements: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get announcements: %v", err)
	}
	return infos.([]*npool.UserAnnouncement), total, nil
}

func GetUserAnnouncement(ctx context.Context, id string) (*npool.UserAnnouncement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUserAnnouncement(ctx, &npool.GetUserAnnouncementRequest{
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
	return info.(*npool.UserAnnouncement), nil
}
