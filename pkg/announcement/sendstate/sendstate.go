package sendstate

import (
	"context"
	"fmt"

	announcementpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"
	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/announcement/sendstate"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"
	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	"github.com/NpoolPlatform/notif-manager/pkg/db"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
	entannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	entsendannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/sendannouncement"

	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/sendstate"
)

func CreateSendState(
	ctx context.Context,
	appID, userID, announcementID string,
	channel channelpb.NotifChannel,
) error {
	exist, err := mgrcli.ExistSendStateConds(ctx, &mgrpb.Conds{
		AppID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UserID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		AnnouncementID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: announcementID,
		},
		Channel: &valuedef.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(channel.Number()),
		},
	})
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("send state already exist")
	}

	_, err = mgrcli.CreateSendState(ctx, &mgrpb.SendStateReq{
		AppID:          &appID,
		UserID:         &userID,
		AnnouncementID: &announcementID,
		Channel:        &channel,
	})
	if err != nil {
		return err
	}
	return nil
}

func CreateSendStates(
	ctx context.Context,
	infos []*mgrpb.SendStateReq,
) error {
	_, err := mgrcli.CreateSendStates(ctx, infos)
	if err != nil {
		return err
	}
	return nil
}

func GetSendStates(
	ctx context.Context,
	conds *npool.Conds,
	offset, limit int32,
) (
	[]*npool.SendState,
	uint32,
	error,
) {
	var infos []*npool.SendState
	var total uint32

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&mgrpb.Conds{
			ID:             conds.ID,
			AppID:          conds.AppID,
			UserID:         conds.UserID,
			AnnouncementID: conds.AnnouncementID,
			Channel:        conds.Channel,
			EndAt:          conds.EndAt,
			UserIDs:        conds.UserIDs,
		}, cli)
		if err != nil {
			return err
		}

		sel := join(stm)
		_total, err := sel.Count(ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)
		return sel.
			Offset(int(offset)).
			Limit(int(limit)).
			Modify(func(s *sql.Selector) {
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetSendState", "err", err)
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}

	infos = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetSendState", "err", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func join(stm *ent.SendAnnouncementQuery) *ent.SendAnnouncementSelect {
	return stm.Select().Modify(func(s *sql.Selector) {
		s.Select(
			s.C(entsendannouncement.FieldAppID),
			s.C(entsendannouncement.FieldUserID),
			s.C(entsendannouncement.FieldChannel),
		)
		t1 := sql.Table(entannouncement.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entsendannouncement.FieldAnnouncementID),
				t1.C(entannouncement.FieldID),
			).
			AppendSelect(
				sql.As(t1.C(entannouncement.FieldID), "announcement_id"),
				t1.C(entannouncement.FieldTitle),
				t1.C(entannouncement.FieldContent),
				t1.C(entannouncement.FieldType),
			)
	})
}

func expand(infos []*npool.SendState) []*npool.SendState {
	for key := range infos {
		infos[key].AnnouncementType = announcementpb.AnnouncementType(announcementpb.AnnouncementType_value[infos[key].AnnouncementTypeStr])
		infos[key].Channel = channelpb.NotifChannel(channelpb.NotifChannel_value[infos[key].ChannelStr])
	}
	return infos
}
