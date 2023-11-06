package notif

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
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
	ret = &npool.Notif{
		EntID:        uuid.NewString(),
		AppID:        uuid.NewString(),
		UserID:       uuid.NewString(),
		Notified:     false,
		LangID:       uuid.NewString(),
		EventID:      uuid.NewString(),
		EventType:    basetypes.UsedFor_KYCApproved,
		EventTypeStr: basetypes.UsedFor_KYCApproved.String(),
		UseTemplate:  true,
		Title:        "Title " + uuid.NewString(),
		Content:      "Content " + uuid.NewString(),
		Channel:      basetypes.NotifChannel_ChannelSMS,
		ChannelStr:   basetypes.NotifChannel_ChannelSMS.String(),
		NotifType:    basetypes.NotifType_NotifMulticast,
		NotifTypeStr: basetypes.NotifType_NotifMulticast.String(),
	}

	retReq = &npool.NotifReq{
		EntID:       &ret.EntID,
		AppID:       &ret.AppID,
		UserID:      &ret.UserID,
		Notified:    &ret.Notified,
		LangID:      &ret.LangID,
		EventID:     &ret.EventID,
		EventType:   &ret.EventType,
		UseTemplate: &ret.UseTemplate,
		Title:       &ret.Title,
		Content:     &ret.Content,
		Channel:     &ret.Channel,
		NotifType:   &ret.NotifType,
	}

	rets = []npool.Notif{
		{
			EntID:        uuid.NewString(),
			AppID:        ret.AppID,
			UserID:       uuid.NewString(),
			Notified:     false,
			LangID:       uuid.NewString(),
			EventID:      uuid.NewString(),
			EventType:    basetypes.UsedFor_WithdrawalRequest,
			EventTypeStr: basetypes.UsedFor_WithdrawalRequest.String(),
			UseTemplate:  true,
			Title:        "Title1 " + uuid.NewString(),
			Content:      "Content1 " + uuid.NewString(),
			Channel:      basetypes.NotifChannel_ChannelSMS,
			ChannelStr:   basetypes.NotifChannel_ChannelSMS.String(),
			NotifType:    basetypes.NotifType_NotifUnicast,
			NotifTypeStr: basetypes.NotifType_NotifUnicast.String(),
		},
		{
			EntID:        uuid.NewString(),
			AppID:        ret.AppID,
			UserID:       uuid.NewString(),
			Notified:     false,
			LangID:       uuid.NewString(),
			EventID:      uuid.NewString(),
			EventType:    basetypes.UsedFor_KYCRejected,
			EventTypeStr: basetypes.UsedFor_KYCRejected.String(),
			UseTemplate:  true,
			Title:        "Title2 " + uuid.NewString(),
			Content:      "Content2 " + uuid.NewString(),
			Channel:      basetypes.NotifChannel_ChannelSMS,
			ChannelStr:   basetypes.NotifChannel_ChannelSMS.String(),
			NotifType:    basetypes.NotifType_NotifMulticast,
			NotifTypeStr: basetypes.NotifType_NotifMulticast.String(),
		},
	}

	retsReq = []*npool.NotifReq{
		{
			EntID:       &rets[0].EntID,
			AppID:       &rets[0].AppID,
			UserID:      &rets[0].UserID,
			Notified:    &rets[0].Notified,
			LangID:      &rets[0].LangID,
			EventID:     &rets[0].EventID,
			EventType:   &rets[0].EventType,
			UseTemplate: &rets[0].UseTemplate,
			Title:       &rets[0].Title,
			Content:     &rets[0].Content,
			Channel:     &rets[0].Channel,
			NotifType:   &rets[0].NotifType,
		},
		{
			EntID:       &rets[1].EntID,
			AppID:       &rets[1].AppID,
			UserID:      &rets[1].UserID,
			Notified:    &rets[1].Notified,
			LangID:      &rets[1].LangID,
			EventID:     &rets[1].EventID,
			EventType:   &rets[1].EventType,
			UseTemplate: &rets[1].UseTemplate,
			Title:       &rets[1].Title,
			Content:     &rets[1].Content,
			Channel:     &rets[1].Channel,
			NotifType:   &rets[1].NotifType,
		},
	}
)

func createNotif(t *testing.T) {
	info, err := CreateNotif(context.Background(), retReq)
	if assert.Nil(t, err) {
		ret.UserID = info.UserID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.Notified = info.Notified
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func createNotifs(t *testing.T) {
	infos, err := CreateNotifs(context.Background(), retsReq)
	fmt.Println("creates err: ", err)
	if assert.Nil(t, err) {
		for key := range infos {
			if infos[key].EntID == rets[0].EntID {
				rets[0].ID = infos[key].ID
			}
			if infos[key].EntID == rets[1].EntID {
				rets[1].ID = infos[key].ID
			}
		}
		assert.Equal(t, len(infos), 2)
	}
}

func updateNotif(t *testing.T) {
	ret.Notified = true
	retReq.ID = &ret.ID
	info, err := UpdateNotif(context.Background(), retReq)
	if assert.Nil(t, err) {
		ret.UserID = info.UserID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.Notified = info.Notified
		assert.Equal(t, ret, info)
	}
}

func updateNotifs(t *testing.T) {
	b := true
	updReq := []*npool.NotifReq{
		{
			ID:       &rets[0].ID,
			Notified: &b,
		},
		{
			ID:       &rets[1].ID,
			Notified: &b,
		},
	}
	infos, err := UpdateNotifs(context.Background(), updReq)
	fmt.Println("UPDATEs err: ", err)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getNotif(t *testing.T) {
	info, err := GetNotif(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret.String(), info.String())
	}
}

func getNotifs(t *testing.T) {
	infos, total, err := GetNotifs(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
	}, 0, 3)
	fmt.Println("GETS err: ", err)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(3))
		assert.Equal(t, len(infos), 3)
	}
}

func getNotifOnly(t *testing.T) {
	info, err := GetNotifOnly(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func existNotifConds(t *testing.T) {
	exist, err := ExistNotifConds(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		LangID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.LangID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteNotif(t *testing.T) {
	info, err := DeleteNotif(context.Background(), &npool.NotifReq{
		ID: &ret.ID,
	})
	assert.Nil(t, err)
	info, err = GetNotif(context.Background(), info.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
	for key := range rets {
		info, err := DeleteNotif(context.Background(), &npool.NotifReq{
			ID: &rets[key].ID,
		})
		if assert.Nil(t, err) {
			assert.NotEqual(t, info, nil)
		}
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createNotif", createNotif)
	t.Run("createNotifs", createNotifs)
	t.Run("updateNotif", updateNotif)
	t.Run("updateNotifs", updateNotifs)
	t.Run("getNotif", getNotif)
	t.Run("getNotifs", getNotifs)
	t.Run("getNotifOnly", getNotifOnly)
	t.Run("existNotifConds", existNotifConds)
	t.Run("deleteNotif", deleteNotif)
}
