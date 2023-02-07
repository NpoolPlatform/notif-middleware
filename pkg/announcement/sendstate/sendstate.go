package sendstate

import (
	"context"

	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/announcement/sendstate"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"
	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
	entannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	entsendannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/sendannouncement"
)

func CreateSendState(
	ctx context.Context,
	appID, userID, announcementID string,
	channel channelpb.NotifChannel,
) error {
	_, err := mgrcli.CreateSendState(ctx, &mgrpb.SendStateReq{
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
		stm := cli.Debug().
			SendAnnouncement.
			Query()
		if conds != nil {
			if conds.AnnouncementID != nil {
				stm.Where(
					entsendannouncement.AnnouncementID(uuid.MustParse(conds.GetAnnouncementID().GetValue())),
				)
			}
			if conds.AppID != nil {
				stm.Where(
					entsendannouncement.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
				)
			}
			if conds.UserID != nil {
				stm.Where(
					entsendannouncement.UserID(uuid.MustParse(conds.GetUserID().GetValue())),
				)
			}

			if conds.Channel != nil {
				val := conds.GetChannel().GetValue()
				channelStr := channelpb.NotifChannel(val).String()
				stm.Where(
					entsendannouncement.Channel(channelStr),
				)
			}
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
			)
	})
}

func expand(infos []*npool.SendState) []*npool.SendState {
	for key := range infos {
		infos[key].Channel = channelpb.NotifChannel(channelpb.NotifChannel_value[infos[key].ChannelStr])
	}
	return infos
}
