package send

import (
	"context"
	"fmt"

	"time"

	"github.com/NpoolPlatform/notif-middleware/pkg/servicename"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/send"
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

func CreateSendAnnouncement(ctx context.Context, in *npool.SendAnnouncementReq) (*npool.SendAnnouncement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateSendAnnouncement(ctx, &npool.CreateSendAnnouncementRequest{
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
	return info.(*npool.SendAnnouncement), nil
}

func DeleteSendAnnouncement(ctx context.Context, id string) (*npool.SendAnnouncement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteSendAnnouncement(ctx, &npool.DeleteSendAnnouncementRequest{
			Info: &npool.SendAnnouncementReq{
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
	return info.(*npool.SendAnnouncement), nil
}

func GetSendAnnouncements(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.SendAnnouncement, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSendAnnouncements(ctx, &npool.GetSendAnnouncementsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get send announcements: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get send announcements: %v", err)
	}
	return infos.([]*npool.SendAnnouncement), total, nil
}

func GetSendAnnouncement(ctx context.Context, id string) (*npool.SendAnnouncement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSendAnnouncement(ctx, &npool.GetSendAnnouncementRequest{
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
	return info.(*npool.SendAnnouncement), nil
}
