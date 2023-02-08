//nolint:dupl
package announcements

import (
	"context"
	"fmt"
	"time"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get announcements connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func GetAnnouncements(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Announcement, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncements(ctx, &npool.GetAnnouncementsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcementss: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get announcementss: %v", err)
	}
	return infos.([]*npool.Announcement), total, nil
}

func GetAnnouncementStates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Announcement, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncementStates(ctx, &npool.GetAnnouncementStatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcementss: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get announcementss: %v", err)
	}
	return infos.([]*npool.Announcement), total, nil
}
