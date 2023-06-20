package sendstate

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
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"
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
		NotifType:    notifpb.NotifType_Multicast,
		NotifTypeStr: notifpb.NotifType_Multicast.String(),
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

	ret = npool.SendState{
		AppID:      appID,
		NotifID:    "",
		UserID:     userID,
		Channel:    notifData.Channel,
		ChannelStr: notifData.ChannelStr,
	}
)

func setupSendState(t *testing.T) func(*testing.T) {
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

func createSendState(t *testing.T) {
	info, err := CreateSendState(context.Background(), &npool.SendStateReq{
		AppID:   &ret.AppID,
		UserID:  &ret.UserID,
		NotifID: &ret.NotifID,
		Channel: &ret.Channel,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
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
	}, 0, 100)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteSendState(t *testing.T) {
	info, err := DeleteSendState(context.Background(), &npool.SendStateReq{
		ID: &ret.ID,
	})
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
	t.Run("getSendState", getSendState)
	t.Run("getSendStates", getSendStates)
	t.Run("deleteSendState", deleteSendState)

	patch.Unpatch()
}
