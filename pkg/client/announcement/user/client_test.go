package user

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
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
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
		EndAt:               uint32(time.Now().Add(1 * time.Hour).Unix()),
	}

	ret = npool.AnnouncementUser{
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

func setupAnnouncementUser(t *testing.T) func(*testing.T) {
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

	handler, err := amt1.NewHandler(
		context.Background(),
		amt1.WithTitle(&amt.Title),
		amt1.WithContent(&amt.Content),
		amt1.WithAppID(&amt.AppID),
		amt1.WithLangID(&amt.AppID, &amt.LangID),
		amt1.WithChannel(&amt.Channel),
		amt1.WithAnnouncementType(&amt.AnnouncementType),
		amt1.WithEndAt(&amt.EndAt),
	)
	assert.Nil(t, err)

	_amt, err := handler.CreateAnnouncement(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _amt)

	ret.AnnouncementID = _amt.ID

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
		_, _ = appusercli.DeleteUser(context.Background(), ret.AppID, ret.UserID)
		_, _ = handler.DeleteAnnouncement(context.Background())
	}
}

func createAnnouncementUser(t *testing.T) {
	info, err := CreateAnnouncementUser(context.Background(), &npool.AnnouncementUserReq{
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

func getAnnouncementUser(t *testing.T) {
	info, err := GetAnnouncementUser(context.Background(), ret.AppID, ret.ID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getAnnouncementUsers(t *testing.T) {
	infos, _, err := GetAnnouncementUsers(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteAnnouncementUser(t *testing.T) {
	info, err := DeleteAnnouncementUser(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetAnnouncementUser(context.Background(), info.AppID, info.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupAnnouncementUser(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	t.Run("createAnnouncementUser", createAnnouncementUser)
	t.Run("getAnnouncementUser", getAnnouncementUser)
	t.Run("getAnnouncementUsers", getAnnouncementUsers)
	t.Run("deleteAnnouncementUser", deleteAnnouncementUser)

	patch.Unpatch()
}
