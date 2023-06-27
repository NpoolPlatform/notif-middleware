package send

import (
	"context"
	"fmt"
	"math/rand"

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

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appusercli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	appuserpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	amtpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
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
	appID = uuid.NewString()
	amt   = amtpb.Announcement{
		AppID:               appID,
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
		AppID:            appID,
		AnnouncementID:   "",
		UserID:           "",
		LangID:           amt.LangID,
		Title:            amt.Title,
		Content:          amt.Content,
		Channel:          amt.ChannelStr,
		AnnouncementType: amt.AnnouncementTypeStr,
		EndAt:            amt.EndAt,
	}
)

func setupSendState(t *testing.T) func(*testing.T) {
	app1, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			ID:        &appID,
			CreatedBy: &appID,
			Name:      &appID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	var (
		id           = uuid.NewString()
		appID        = app1.ID
		emailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+1000000) //nolint
		passwordHash = uuid.NewString()
		req          = &appuserpb.UserReq{
			ID:           &id,
			AppID:        &appID,
			EmailAddress: &emailAddress,
			PasswordHash: &passwordHash,
		}
	)

	user, err := appusercli.CreateUser(context.Background(), req)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	ret.UserID = user.ID

	// Create Announcement First
	amtHandler, err := announcement1.NewHandler(
		context.Background(),
		announcement1.WithTitle(&amt.Title),
		announcement1.WithContent(&amt.Content),
		announcement1.WithAppID(&amt.AppID),
		announcement1.WithLangID(&amt.LangID),
		announcement1.WithChannel(&amt.Channel),
		announcement1.WithAnnouncementType(&amt.AnnouncementType),
		announcement1.WithStartAt(&amt.StartAt),
		announcement1.WithEndAt(&amt.EndAt),
	)
	assert.Nil(t, err)

	announcement, err := amtHandler.CreateAnnouncement(context.Background())
	assert.Nil(t, err)
	ret.AnnouncementID = announcement.ID

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
		_, _ = appusercli.DeleteUser(context.Background(), ret.AppID, ret.UserID)
		_, _ = amtHandler.DeleteAnnouncement(context.Background())
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
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
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
