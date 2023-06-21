package readstate

import (
	"context"
	"fmt"

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
	userID    = uuid.NewString()
	notifData = notifpb.Notif{
		ID:           uuid.NewString(),
		AppID:        appID,
		Notified:     false,
		UserID:       userID,
		LangID:       uuid.NewString(),
		EventID:      uuid.NewString(),
		EventType:    basetypes.UsedFor_WithdrawalRequest,
		EventTypeStr: basetypes.UsedFor_WithdrawalRequest.String(),
		UseTemplate:  true,
		Title:        "Title " + uuid.NewString(),
		Content:      "Content " + uuid.NewString(),
		Channel:      basetypes.NotifChannel_ChannelEmail,
		ChannelStr:   basetypes.NotifChannel_ChannelEmail.String(),
		NotifType:    basetypes.NotifType_NotifMulticast,
		NotifTypeStr: basetypes.NotifType_NotifMulticast.String(),
	}

	notifRet = &notifpb.NotifReq{
		ID:          &notifData.ID,
		AppID:       &notifData.AppID,
		Notified:    &notifData.Notified,
		LangID:      &notifData.LangID,
		EventID:     &notifData.EventID,
		EventType:   &notifData.EventType,
		UseTemplate: &notifData.UseTemplate,
		Title:       &notifData.Title,
		Content:     &notifData.Content,
		Channel:     &notifData.Channel,
		NotifType:   &notifData.NotifType,
	}

	ret = npool.ReadState{
		AppID:   appID,
		NotifID: "",
		UserID:  userID,
	}
)

func setupReadState(t *testing.T) func(*testing.T) {
	n1, err := notifmw.NewHandler(
		context.Background(),
		notifmw.WithID(notifRet.ID),
		notifmw.WithAppID(notifRet.AppID),
		notifmw.WithLangID(notifRet.LangID),
		notifmw.WithEventID(notifRet.EventID),
		notifmw.WithNotified(notifRet.Notified),
		notifmw.WithEventType(notifRet.EventType),
		notifmw.WithUseTemplate(notifRet.UseTemplate),
		notifmw.WithTitle(notifRet.Title),
		notifmw.WithContent(notifRet.Content),
		notifmw.WithChannel(notifRet.Channel),
		notifmw.WithNotifType(notifRet.NotifType),
	)
	assert.Nil(t, err)

	_notif, err := n1.CreateNotif(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _notif)

	ret.NotifID = _notif.ID

	return func(*testing.T) {
		_, _ = n1.DeleteNotif(context.Background())
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
	}, 0, 100)
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
