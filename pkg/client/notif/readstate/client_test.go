package readstate

import (
	"context"
	"fmt"
	"math/rand"

	"os"
	"strconv"
	"testing"

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
	notifpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	notifmw "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif"

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
	appID     = uuid.NewString()
	nType     = notifpb.NotifType_Multicast
	notifData = notifpb.Notif{
		AppID:        appID,
		UserID:       uuid.NewString(),
		LangID:       uuid.NewString(),
		Title:        uuid.NewString(),
		Content:      uuid.NewString(),
		Channel:      basetypes.NotifChannel_ChannelEmail,
		ChannelStr:   basetypes.NotifChannel_ChannelEmail.String(),
		NotifTypeStr: nType.String(),
		NotifType:    nType,
	}

	ret = npool.ReadState{
		ID:      uuid.NewString(),
		NotifID: uuid.NewString(),
		AppID:   uuid.NewString(),
		UserID:  uuid.NewString(),
	}
)

func setupReadState(t *testing.T) func(*testing.T) {
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

	handler, err := notifmw.NewHandler(
		context.Background(),
		notifmw.WithAppID(&notifData.AppID),
		notifmw.WithUserID(&notifData.UserID),
	)
	assert.Nil(t, err)

	_amt, err := handler.CreateNotif(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _amt)

	ret.NotifID = _amt.ID

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
		_, _ = appusercli.DeleteUser(context.Background(), ret.AppID, ret.UserID)
		_, _ = handler.DeleteNotif(context.Background())
	}
}

func createReadState(t *testing.T) {
	info, err := CreateReadState(context.Background(), &npool.ReadStateReq{
		AppID:   &ret.AppID,
		UserID:  &ret.UserID,
		NotifID: &ret.NotifID,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func getReadState(t *testing.T) {
	info, err := GetReadState(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getReadStates(t *testing.T) {
	infos, _, err := GetReadStates(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteReadState(t *testing.T) {
	info, err := DeleteReadState(context.Background(), &npool.ReadStateReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetReadState(context.Background(), info.ID)
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
	t.Run("createReadState", createReadState)
	t.Run("getReadState", getReadState)
	t.Run("getReadStates", getReadStates)
	t.Run("deleteReadState", deleteReadState)

	patch.Unpatch()
}
