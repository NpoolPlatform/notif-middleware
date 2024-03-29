//nolint
package announcement

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/notif-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
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

func CreateAnnouncement(ctx context.Context, in *npool.AnnouncementReq) (*npool.Announcement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateAnnouncement(ctx, &npool.CreateAnnouncementRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create announcement conds: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Announcement), nil
}

func DeleteAnnouncement(ctx context.Context, id uint32) (*npool.Announcement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteAnnouncement(ctx, &npool.DeleteAnnouncementRequest{
			Info: &npool.AnnouncementReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete announcement conds: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Announcement), nil
}

func UpdateAnnouncement(ctx context.Context, in *npool.AnnouncementReq) (*npool.Announcement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateAnnouncement(ctx, &npool.UpdateAnnouncementRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update announcement conds: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Announcement), nil
}

func ExistAnnouncement(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistAnnouncement(ctx, &npool.ExistAnnouncementRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail exist announcement: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func ExistAnnouncementConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistAnnouncementConds(ctx, &npool.ExistAnnouncementCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail exist announcement conds: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func GetAnnouncements(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Announcement, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncements(ctx, &npool.GetAnnouncementsRequest{
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
	return infos.([]*npool.Announcement), total, nil
}

func GetAnnouncement(ctx context.Context, id string) (*npool.Announcement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncement(ctx, &npool.GetAnnouncementRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcement: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Announcement), nil
}

func GetAnnouncementOnly(ctx context.Context, conds *npool.Conds) (*npool.Announcement, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncements(ctx, &npool.GetAnnouncementsRequest{
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
	if len(infos.([]*npool.Announcement)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Announcement)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Announcement)[0], nil
}
