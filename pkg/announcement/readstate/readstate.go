package readstate

import (
	"context"
	"encoding/json"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"
	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
	entannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	entreadannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/readannouncement"
)

func GetReadState(ctx context.Context, announcementID, userID string) (*npool.ReadState, error) {
	var infos []*npool.ReadState
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Announcement.
			Query().
			Where(
				entannouncement.ID(uuid.MustParse(announcementID)),
			).
			Limit(1)

		return join(stm, &userID).
			Scan(_ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetReadState", "err", err)
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}

	infos, err = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetReadState", "err", err)
		return nil, err
	}
	return infos[0], nil
}

func GetReadStates(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.ReadState, uint32, error) {
	var infos []*npool.ReadState
	var total uint32
	var err error
	var userID *string

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

			if conds.UserID != nil {
				val := conds.GetUserID().GetValue()
				userID = &val
			}
		}

		sel := join(stm, userID)
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
		logger.Sugar().Errorw("GetReadState", "err", err)
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}

	infos, err = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetReadState", "err", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func join(stm *ent.AnnouncementQuery, userID *string) *ent.AnnouncementSelect {
	return stm.Select().Modify(func(s *sql.Selector) {
		s.Select(
			sql.As(s.C(entannouncement.FieldID), "announcement_id"),
			s.C(entannouncement.FieldAppID),
			s.C(entannouncement.FieldTitle),
			s.C(entannouncement.FieldContent),
			s.C(entannouncement.FieldChannels),
			s.C(entannouncement.FieldCreatedAt),
			s.C(entannouncement.FieldUpdatedAt),
		)
		t1 := sql.Table(entreadannouncement.Table)
		s.
			LeftJoin(t1).
			On(
				s.C(entannouncement.FieldID),
				t1.C(entreadannouncement.FieldAnnouncementID),
			)
		if userID != nil {
			s.
				OnP(
					sql.EQ(t1.C(entreadannouncement.FieldUserID), *userID),
				)
		}
		s.
			AppendSelect(
				t1.C(entreadannouncement.FieldUserID),
			)
	})
}

func expand(infos []*npool.ReadState) ([]*npool.ReadState, error) {
	for key, info := range infos {
		if info.UserID != "" {
			infos[key].AlreadyRead = true
		}

		channelsStr := []string{}
		err := json.Unmarshal([]byte(infos[0].ChannelsStr), &channelsStr)
		if err != nil {
			return nil, err
		}
		channels := []channelpb.NotifChannel{}
		for _, channel := range channelsStr {
			channels = append(channels, channelpb.NotifChannel(channelpb.NotifChannel_value[channel]))
		}
		infos[key].Channels = channels
	}
	return infos, nil
}
