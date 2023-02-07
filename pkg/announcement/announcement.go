package sendstate

import (
	"context"

	entreadannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/readannouncement"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
	entannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	entsendannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/sendannouncement"
)

func GetAnnouncements(
	ctx context.Context,
	conds *npool.Conds,
	offset, limit int32,
) (
	[]*npool.Announcement,
	uint32,
	error,
) {
	var infos []*npool.Announcement
	var total uint32
	var err error
	var userID *string
	var channel *string
	var userIDs []string
	var alreadySend *bool
	var alreadyRead *bool

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Announcement.
			Query()
		if conds != nil {
			if conds.AnnouncementID != nil {
				stm.Where(
					entannouncement.ID(uuid.MustParse(conds.GetAnnouncementID().GetValue())),
				)
			}
			if conds.AppID != nil {
				stm.Where(
					entannouncement.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
				)
			}
			if conds.EndAt != nil {
				stm.Where(
					entannouncement.EndAt(conds.GetEndAt().GetValue()),
				)
			}

			if conds.UserID != nil {
				val := conds.GetUserID().GetValue()
				userID = &val
			}

			if conds.Channel != nil {
				val := conds.GetChannel().GetValue()
				channelStr := channelpb.NotifChannel(val).String()
				channel = &channelStr
			}

			if conds.AlreadySend != nil {
				val := conds.GetAlreadySend().GetValue()
				alreadySend = &val
			}

			if conds.AlreadyRead != nil {
				val := conds.GetAlreadySend().GetValue()
				alreadySend = &val
			}
		}

		sel := join(stm, userID, channel, userIDs, alreadySend, alreadyRead)
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

func join(stm *ent.AnnouncementQuery, userID, channel *string, userIDs []string, alreadySend, alreadyRead *bool) *ent.AnnouncementSelect {
	return stm.Select().Modify(func(s *sql.Selector) {
		s.Select(
			sql.As(s.C(entannouncement.FieldID), "announcement_id"),
			s.C(entannouncement.FieldAppID),
			s.C(entannouncement.FieldTitle),
			s.C(entannouncement.FieldContent),
			s.C(entannouncement.FieldCreatedAt),
			s.C(entannouncement.FieldUpdatedAt),
		)
		t1 := sql.Table(entsendannouncement.Table)
		t2 := sql.Table(entreadannouncement.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entannouncement.FieldID),
				t1.C(entsendannouncement.FieldAnnouncementID),
			).
			LeftJoin(t2).
			On(
				s.C(entannouncement.FieldID),
				t2.C(entreadannouncement.FieldAnnouncementID),
			)
		s.AppendSelect(
			sql.As(t1.C(entsendannouncement.FieldUserID), "user_id"),
			sql.As(t2.C(entreadannouncement.FieldUserID), "read_user_id"),
			t1.C(entsendannouncement.FieldChannel),
		)
		if userID != nil {
			s.
				OnP(
					sql.EQ(t1.C(entsendannouncement.FieldUserID), *userID),
				).
				OnP(
					sql.EQ(t2.C(entreadannouncement.FieldUserID), *userID),
				)
		}
		if len(userIDs) > 0 {
			s.
				OnP(
					sql.In(t1.C(entsendannouncement.FieldUserID), userIDs),
				).
				OnP(
					sql.In(t2.C(entreadannouncement.FieldUserID), userIDs),
				)
		}
		if channel != nil {
			s.
				OnP(
					sql.EQ(t1.C(entsendannouncement.FieldChannel), *channel),
				)
		}
		if alreadySend != nil {
			if *alreadySend {
				s.
					OnP(
						sql.NEQ(t1.C(entsendannouncement.FieldUserID), ""),
					)
			} else {
				s.
					OnP(
						sql.EQ(t1.C(entsendannouncement.FieldUserID), ""),
					)
			}
		}
		if alreadyRead != nil {
			if *alreadyRead {
				s.
					OnP(
						sql.NEQ(t1.C(entreadannouncement.FieldUserID), ""),
					)
			} else {
				s.
					OnP(
						sql.EQ(t1.C(entreadannouncement.FieldUserID), ""),
					)
			}
		}
	})
}

func expand(infos []*npool.Announcement) []*npool.Announcement {
	for key, info := range infos {
		if info.UserID != "" {
			infos[key].AlreadySend = true
		}
		if info.ReadUserID != "" {
			infos[key].AlreadyRead = true
		}

		infos[key].Channel = channelpb.NotifChannel(channelpb.NotifChannel_value[infos[key].ChannelStr])
	}
	return infos
}
