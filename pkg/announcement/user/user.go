package user

import (
	"context"
	"encoding/json"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/user"
	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"

	announcementpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"
	entannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	entuserannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/userannouncement"
)

func GetUsers(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.User, uint32, error) {
	var infos []*npool.User
	var total uint32
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.Debug().
			UserAnnouncement.
			Query()
		if conds != nil {
			if conds.AppID != nil {
				stm.Where(
					entuserannouncement.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
				)
			}
			if conds.UserID != nil {
				stm.Where(
					entuserannouncement.UserID(uuid.MustParse(conds.GetUserID().GetValue())),
				)
			}
			if len(conds.GetIDs().GetValue()) > 0 {
				ids := []uuid.UUID{}
				for _, idStr := range conds.GetIDs().GetValue() {
					id, err := uuid.Parse(idStr)
					if err != nil {
						return err
					}
					ids = append(ids, id)
				}
				stm.Where(
					entuserannouncement.IDIn(ids...),
				)
			}
			if conds.AnnouncementID != nil {
				stm.Where(
					entuserannouncement.AnnouncementID(uuid.MustParse(conds.GetAnnouncementID().GetValue())),
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
		logger.Sugar().Errorw("GetUser", "err", err)
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}

	infos, err = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetUser", "err", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func join(stm *ent.UserAnnouncementQuery) *ent.UserAnnouncementSelect {
	return stm.Select().Modify(func(s *sql.Selector) {
		s.Select(
			sql.As(s.C(entuserannouncement.FieldID), "id"),
			sql.As(s.C(entuserannouncement.FieldAnnouncementID), "announcement_id"),
			s.C(entuserannouncement.FieldAppID),
			s.C(entuserannouncement.FieldUserID),
			s.C(entuserannouncement.FieldCreatedAt),
			s.C(entuserannouncement.FieldUpdatedAt),
		)
		t1 := sql.Table(entannouncement.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entuserannouncement.FieldAnnouncementID),
				t1.C(entannouncement.FieldID),
			)
		s.
			AppendSelect(
				t1.C(entannouncement.FieldTitle),
				t1.C(entannouncement.FieldContent),
				t1.C(entannouncement.FieldChannels),
				t1.C(entannouncement.FieldType),
			)
	})
}

func expand(infos []*npool.User) ([]*npool.User, error) {
	for key := range infos {
		channelsStr := []string{}
		if infos[key].ChannelsStr != "" {
			err := json.Unmarshal([]byte(infos[key].ChannelsStr), &channelsStr)
			if err != nil {
				return nil, err
			}
		}

		channels := []channelpb.NotifChannel{}
		for _, channel := range channelsStr {
			channels = append(channels, channelpb.NotifChannel(channelpb.NotifChannel_value[channel]))
		}

		infos[key].AnnouncementType = announcementpb.AnnouncementType(announcementpb.AnnouncementType_value[infos[key].AnnouncementTypeStr])
		infos[key].Channels = channels
	}
	return infos, nil
}
