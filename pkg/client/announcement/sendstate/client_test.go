package send

import (
	"context"
	"fmt"

	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	announcementpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"github.com/stretchr/testify/assert"
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
	info, err := CreateSendState(context.Background(), &npool.SendStateReq{
		AppID:          &ret.AppID,
		UserID:         &ret.UserID,
		AnnouncementID: &ret.AnnouncementID,
	})
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
	infos, err := CreateSendStates(context.Background(), reqs)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(infos))
	_, err = DeleteSendState(context.Background(), infos[0].ID)
	assert.Nil(t, err)
}

func getSendState(t *testing.T) {
	info, err := GetSendState(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getSendStates(t *testing.T) {
	infos, _, err := GetSendStates(context.Background(), &npool.Conds{
		AppID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AnnouncementID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AnnouncementID},
		Channel:        &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.NotifChannel(basetypes.NotifChannel_value[ret.Channel]))},
		UserIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteSendState(t *testing.T) {
	info, err := DeleteSendState(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetSendState(context.Background(), info.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupSendState(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	t.Run("createSendState", createSendState)
	t.Run("createSendStates", createSendStates)
	t.Run("getSendState", getSendState)
	t.Run("getSendStates", getSendStates)
	t.Run("deleteSendState", deleteSendState)

	patch.Unpatch()
}
