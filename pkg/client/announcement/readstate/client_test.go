package read

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
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"
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
		EntID:               uuid.NewString(),
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

	ret = npool.ReadState{
		AppID:            announcement.AppID,
		UserID:           uuid.NewString(),
		AnnouncementID:   "",
		LangID:           announcement.LangID,
		Title:            announcement.Title,
		Content:          announcement.Content,
		Channel:          announcement.ChannelStr,
		AnnouncementType: announcement.AnnouncementTypeStr,
		EndAt:            announcement.EndAt,
	}
)

func setupReadState(t *testing.T) func(*testing.T) {
	// Create Announcement First
	announcementHandler, err := announcement1.NewHandler(
		context.Background(),
		announcement1.WithTitle(&announcement.Title, true),
		announcement1.WithContent(&announcement.Content, true),
		announcement1.WithAppID(&announcement.AppID, true),
		announcement1.WithLangID(&announcement.LangID, true),
		announcement1.WithChannel(&announcement.Channel, true),
		announcement1.WithAnnouncementType(&announcement.AnnouncementType, true),
		announcement1.WithStartAt(&announcement.StartAt, true),
		announcement1.WithEndAt(&announcement.EndAt, true),
	)
	assert.Nil(t, err)

	announcement, err := announcementHandler.CreateAnnouncement(context.Background())
	assert.Nil(t, err)
	ret.AnnouncementID = announcement.EntID

	announcementHandler.ID = &announcement.ID

	return func(*testing.T) {
		_, _ = announcementHandler.DeleteAnnouncement(context.Background())
	}
}

func createReadState(t *testing.T) {
	info, err := CreateReadState(context.Background(), &npool.ReadStateReq{
		AppID:          &ret.AppID,
		UserID:         &ret.UserID,
		AnnouncementID: &ret.AnnouncementID,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, &ret)
	}
}

func getReadState(t *testing.T) {
	info, err := GetReadState(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getReadStates(t *testing.T) {
	infos, _, err := GetReadStates(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteReadState(t *testing.T) {
	info, err := DeleteReadState(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetReadState(context.Background(), info.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupReadState(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createReadState", createReadState)
	t.Run("getReadState", getReadState)
	t.Run("getReadStates", getReadStates)
	t.Run("deleteReadState", deleteReadState)

	patch.Unpatch()
}
