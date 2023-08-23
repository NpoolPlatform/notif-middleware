package sendstate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	announcementpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	handler1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	announcement = announcementpb.Announcement{
		AppID:               uuid.NewString(),
		LangID:              uuid.NewString(),
		Title:               uuid.NewString(),
		Content:             uuid.NewString(),
		Channel:             basetypes.NotifChannel_ChannelEmail,
		ChannelStr:          basetypes.NotifChannel_ChannelEmail.String(),
		AnnouncementType:    basetypes.NotifType_NotifMulticast,
		AnnouncementTypeStr: basetypes.NotifType_NotifMulticast.String(),
		StartAt:             uint32(time.Now().Add(10 * time.Minute).Unix()),
		EndAt:               uint32(time.Now().Add(1 * time.Hour).Unix()),
	}

	ret = npool.SendState{
		AppID:            announcement.AppID,
		AnnouncementID:   "",
		UserID:           uuid.NewString(),
		LangID:           announcement.LangID,
		Title:            announcement.Title,
		Content:          announcement.Content,
		Channel:          announcement.ChannelStr,
		AnnouncementType: announcement.AnnouncementTypeStr,
		EndAt:            announcement.EndAt,
	}
)

func setupSendState(t *testing.T) func(*testing.T) {
	// Create Announcement First
	announcementHandler, err := announcement1.NewHandler(
		context.Background(),
		announcement1.WithTitle(&announcement.Title),
		announcement1.WithContent(&announcement.Content),
		announcement1.WithAppID(&announcement.AppID),
		announcement1.WithLangID(&announcement.LangID),
		announcement1.WithChannel(&announcement.Channel),
		announcement1.WithAnnouncementType(&announcement.AnnouncementType),
		announcement1.WithStartAt(&announcement.StartAt),
		announcement1.WithEndAt(&announcement.EndAt),
	)
	assert.Nil(t, err)

	announcement, err := announcementHandler.CreateAnnouncement(context.Background())
	assert.Nil(t, err)
	ret.AnnouncementID = announcement.ID

	_id, err := uuid.Parse(announcement.ID)
	assert.Nil(t, err)
	announcementHandler.ID = &_id

	return func(*testing.T) {
		_, _ = announcementHandler.DeleteAnnouncement(context.Background())
	}
}

func createSendState(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler1.WithAppID(&ret.AppID),
		handler1.WithUserID(&ret.UserID),
		handler1.WithAnnouncementID(&ret.AppID, &ret.AnnouncementID),
	)
	assert.Nil(t, err)

	info, err := handler.CreateSendState(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func createSendStates(t *testing.T) {
	userID := uuid.NewString()
	appID := uuid.NewString()
	reqs := []*npool.SendStateReq{{
		AppID:          &appID,
		UserID:         &userID,
		AnnouncementID: &ret.AnnouncementID,
	}}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(reqs),
	)
	assert.Nil(t, err)

	infos, err := handler.CreateSendStates(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(infos))

	handler2, err := NewHandler(
		context.Background(),
		handler1.WithID(&infos[0].ID),
	)
	assert.Nil(t, err)

	_, err = handler2.DeleteSendState(context.Background())
	assert.Nil(t, err)
}

func getSendState(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetSendState(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSendStates(t *testing.T) {
	conds := &npool.Conds{
		AppID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AnnouncementID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AnnouncementID},
		Channel:        &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.NotifChannel(basetypes.NotifChannel_value[ret.Channel]))},
		UserIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		handler1.WithOffset(0),
		handler1.WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetSendStates(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteSendState(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteSendState(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetSendState(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestSendState(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupSendState(t)
	defer teardown(t)

	t.Run("createSendState", createSendState)
	t.Run("createSendStates", createSendStates)
	t.Run("getSendState", getSendState)
	t.Run("getSendStates", getSendStates)
	t.Run("deleteSendState", deleteSendState)
}
