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
	userID3            = uuid.NewString()
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
		NotifType:    basetypes.NotifType_NotifMulticast,
		NotifTypeStr: basetypes.NotifType_NotifMulticast.String(),
	}
	notifInfoUnicast = notifpb.Notif{
		ID:           uuid.NewString(),
		AppID:        appID,
		UserID:       userID1,
		Notified:     true,
		LangID:       uuid.NewString(),
		EventID:      uuid.NewString(),
		EventType:    basetypes.UsedFor_KYCApproved,
		EventTypeStr: basetypes.UsedFor_KYCApproved.String(),
		UseTemplate:  true,
		Title:        "Unicast Title " + uuid.NewString(),
		Content:      "Content " + uuid.NewString(),
		Channel:      basetypes.NotifChannel_ChannelSMS,
		ChannelStr:   basetypes.NotifChannel_ChannelSMS.String(),
		NotifType:    basetypes.NotifType_NotifUnicast,
		NotifTypeStr: basetypes.NotifType_NotifUnicast.String(),
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

	ret = npool.NotifUser{
		ID:        "",
		AppID:     appID,
		EventType: notifInfoUnicast.EventType,
		UserID:    userID1,
	}

	rets = []npool.NotifUser{
		{
			ID:        uuid.NewString(),
			AppID:     appID,
			EventType: notifInfoMulticast.EventType,
			UserID:    userID2,
		},
		{
			ID:        uuid.NewString(),
			AppID:     appID,
			EventType: notifInfoMulticast.EventType,
			UserID:    userID3,
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

func createNotifUser(t *testing.T) {
	info, err := CreateNotifUser(context.Background(), &npool.NotifUserReq{
		AppID:     &ret.AppID,
		UserID:    &ret.UserID,
		EventType: &ret.EventType,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

//nolint:vet
func getNotifUser(t *testing.T) {
	for _, item := range rets {
		info, err := GetNotifUser(context.Background(), item.ID)
		assert.Nil(t, err)
		assert.NotNil(t, info)
	}
}

func getNotifUsers(t *testing.T) {
	infos, _, err := GetNotifUsers(context.Background(), &npool.Conds{
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
	info, err := DeleteNotifUser(context.Background(), &npool.NotifUserReq{
		ID:    &ret.ID,
		AppID: &ret.AppID,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
	info, err = GetNotifUser(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)

	for _, item := range rets {
		info, err := DeleteNotifUser(context.Background(), &npool.NotifUserReq{
			ID:    &item.ID,
			AppID: &item.AppID,
		})
		if assert.Nil(t, err) {
			item.CreatedAt = info.CreatedAt
			item.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &item)
		}
		info, err = GetNotifUser(context.Background(), item.ID)
		assert.Nil(t, err)
		assert.Nil(t, info)
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
