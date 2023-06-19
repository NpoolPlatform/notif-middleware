package user

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
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
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
	userID1            = uuid.NewString()
	userID2            = uuid.NewString()
	appID              = uuid.NewString()
	notifInfoMulticast = notifpb.Notif{
		ID:           uuid.NewString(),
		AppID:        appID,
		UserID:       "",
		Notified:     false,
		LangID:       uuid.NewString(),
		EventID:      uuid.NewString(),
		EventType:    basetypes.UsedFor_WithdrawalRequest,
		EventTypeStr: basetypes.UsedFor_WithdrawalRequest.String(),
		UseTemplate:  true,
		Title:        "Multicast Title " + uuid.NewString(),
		Content:      "Content " + uuid.NewString(),
		Channel:      basetypes.NotifChannel_ChannelEmail,
		ChannelStr:   basetypes.NotifChannel_ChannelEmail.String(),
		NotifType:    notifpb.NotifType_Multicast,
		NotifTypeStr: notifpb.NotifType_Multicast.String(),
	}
	notifInfoUnicast = notifpb.Notif{
		ID:           uuid.NewString(),
		AppID:        appID,
		UserID:       userID1,
		Notified:     false,
		LangID:       uuid.NewString(),
		EventID:      uuid.NewString(),
		EventType:    basetypes.UsedFor_KYCApproved,
		EventTypeStr: basetypes.UsedFor_KYCApproved.String(),
		UseTemplate:  true,
		Title:        "Unicast Title " + uuid.NewString(),
		Content:      "Content " + uuid.NewString(),
		Channel:      basetypes.NotifChannel_ChannelSMS,
		ChannelStr:   basetypes.NotifChannel_ChannelSMS.String(),
		NotifType:    notifpb.NotifType_Unicast,
		NotifTypeStr: notifpb.NotifType_Unicast.String(),
	}

	notifReqs = []*notifpb.NotifReq{
		{
			ID:          &notifInfoUnicast.ID,
			AppID:       &notifInfoUnicast.AppID,
			UserID:      &notifInfoUnicast.UserID,
			Notified:    &notifInfoUnicast.Notified,
			LangID:      &notifInfoUnicast.LangID,
			EventID:     &notifInfoUnicast.EventID,
			EventType:   &notifInfoUnicast.EventType,
			UseTemplate: &notifInfoUnicast.UseTemplate,
			Title:       &notifInfoUnicast.Title,
			Content:     &notifInfoUnicast.Content,
			Channel:     &notifInfoUnicast.Channel,
			NotifType:   &notifInfoUnicast.NotifType,
		},
		{
			ID:          &notifInfoMulticast.ID,
			AppID:       &notifInfoMulticast.AppID,
			Notified:    &notifInfoMulticast.Notified,
			LangID:      &notifInfoMulticast.LangID,
			EventID:     &notifInfoMulticast.EventID,
			EventType:   &notifInfoMulticast.EventType,
			UseTemplate: &notifInfoMulticast.UseTemplate,
			Title:       &notifInfoMulticast.Title,
			Content:     &notifInfoMulticast.Content,
			Channel:     &notifInfoMulticast.Channel,
			NotifType:   &notifInfoMulticast.NotifType,
		},
	}

	ret = npool.UserNotif{
		ID:      "",
		AppID:   appID,
		NotifID: "",
		UserID:  userID1,
	}

	rets = []npool.UserNotif{
		{
			ID:      "",
			AppID:   appID,
			NotifID: notifInfoMulticast.ID,
			UserID:  userID1,
		},
		{
			ID:      "",
			AppID:   appID,
			NotifID: notifInfoMulticast.ID,
			UserID:  userID2,
		},
	}
)

func setupNotifUser(t *testing.T) func(*testing.T) {
	n1, err := notifmw.NewHandler(
		context.Background(),
		notifmw.WithReqs(notifReqs),
	)
	assert.Nil(t, err)

	_notif, err := n1.CreateNotifs(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _notif)

	ret.NotifID = _notif[0].ID

	return func(*testing.T) {
		for _, row := range _notif {
			id, _ := uuid.Parse(row.ID)
			n1.ID = &id
			_appID, _ := uuid.Parse(row.AppID)
			n1.AppID = &_appID
			_, _ = n1.DeleteNotif(context.Background())
		}
	}
}

// nolint:gosec,vet
func createNotifUser(t *testing.T) {
	for i, item := range rets {
		info, err := CreateUser(context.Background(), &npool.UserNotifReq{
			AppID:   &item.AppID,
			UserID:  &item.UserID,
			NotifID: &item.NotifID,
		})
		if assert.Nil(t, err) {
			item.CreatedAt = info.CreatedAt
			item.UpdatedAt = info.UpdatedAt
			item.ID = info.ID
			rets[i].ID = info.ID
			assert.Equal(t, info, &item)
		}
	}
}

//nolint:vet
func getNotifUser(t *testing.T) {
	for _, item := range rets {
		info, err := GetUser(context.Background(), item.ID)
		assert.Nil(t, err)
		assert.NotNil(t, info)
	}
}

func getNotifUsers(t *testing.T) {
	infos, _, err := GetUsers(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: rets[0].ID,
		},
	}, 0, 100)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

// nolint:gosec,vet
func deleteNotifUser(t *testing.T) {
	for _, item := range rets {
		info, err := DeleteUser(context.Background(), &npool.UserNotifReq{
			ID: &item.ID,
		})
		if assert.Nil(t, err) {
			item.CreatedAt = info.CreatedAt
			item.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &item)
		}
		info, err = GetUser(context.Background(), item.ID)
		assert.Nil(t, err)
		assert.NotNil(t, info)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupNotifUser(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	t.Run("createNotifUser", createNotifUser)
	t.Run("getNotifUser", getNotifUser)
	t.Run("getNotifUsers", getNotifUsers)
	t.Run("deleteNotifUser", deleteNotifUser)

	patch.Unpatch()
}
