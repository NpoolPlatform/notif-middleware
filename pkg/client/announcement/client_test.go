package announcement

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	readstatemgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	announcementcrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement"
	readstatecrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/readstate"

	userstatemgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/user"
	userstatecrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/user"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appusercli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	appuserpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"
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
	addressFields     = []string{uuid.NewString()}
	addressFieldsS, _ = json.Marshal(addressFields)
	appID             = uuid.NewString()
	ret               = appuserpb.User{
		ID:                  uuid.NewString(),
		AppID:               appID,
		EmailAddress:        "aaa@aaa.aaa",
		PhoneNO:             "+8613612203133",
		ImportedFromAppID:   uuid.NewString(),
		Username:            "amwnrekadsf.are-",
		AddressFieldsString: string(addressFieldsS),
		AddressFields:       addressFields,
		Gender:              uuid.NewString(),
		PostalCode:          uuid.NewString(),
		Age:                 100,
		Birthday:            uint32(time.Now().Unix()),
		Avatar:              uuid.NewString(),
		Organization:        uuid.NewString(),
		FirstName:           uuid.NewString(),
		LastName:            uuid.NewString(),
		IDNumber:            uuid.NewString(),
		GoogleAuthVerified:  true,
		SigninVerifyType:    basetypes.SignMethod_Email,
		SigninVerifyTypeStr: basetypes.SignMethod_Email.String(),
		GoogleSecret:        appID,
		HasGoogleSecret:     true,
		Roles:               []string{""},
		ActionCredits:       "0",
		Banned:              true,
		BanMessage:          uuid.NewString(),
	}
	amt = npool.Announcement{
		AppID:               "",
		LangID:              uuid.NewString(),
		Title:               uuid.NewString(),
		Content:             uuid.NewString(),
		Channel:             basetypes.NotifChannel_ChannelEmail,
		ChannelStr:          basetypes.NotifChannel_ChannelEmail.String(),
		AnnouncementType:    npool.AnnouncementType_Multicast,
		AnnouncementTypeStr: npool.AnnouncementType_Multicast.String(),
		EndAt:               uint32(time.Now().Add(1 * time.Hour).Unix()),
	}
)

func setupUser(t *testing.T) func(*testing.T) {
	app1, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			ID:        &ret.AppID,
			CreatedBy: &ret.ID,
			Name:      &ret.AppID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	app2, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			ID:        &ret.ImportedFromAppID,
			CreatedBy: &ret.ID,
			Name:      &ret.ImportedFromAppID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app2)

	amt.AppID = app2.ID
	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
		_, _ = appmwcli.DeleteApp(context.Background(), ret.ImportedFromAppID)
	}
}

func createUser(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+10000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+10000) //nolint
	ret.ImportedFromAppName = ret.ImportedFromAppID
	var (
		id                = ret.ID
		appID             = ret.AppID
		importedFromAppID = ret.ImportedFromAppID
		strVal            = "AAA"
		req               = appuserpb.UserReq{
			ID:                 &id,
			AppID:              &appID,
			EmailAddress:       &ret.EmailAddress,
			PhoneNO:            &ret.PhoneNO,
			ImportedFromAppID:  &importedFromAppID,
			Username:           &ret.Username,
			AddressFields:      addressFields,
			Gender:             &ret.Gender,
			PostalCode:         &ret.PostalCode,
			Age:                &ret.Age,
			Birthday:           &ret.Birthday,
			Avatar:             &ret.Avatar,
			Organization:       &ret.Organization,
			FirstName:          &ret.FirstName,
			LastName:           &ret.LastName,
			IDNumber:           &ret.IDNumber,
			GoogleAuthVerified: &ret.GoogleAuthVerified,
			SigninVerifyType:   &ret.SigninVerifyType,
			PasswordHash:       &strVal,
			GoogleSecret:       &appID,
			ThirdPartyID:       &strVal,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &ret.Banned,
			BanMessage:         &ret.BanMessage,
		}
		ret1 = appuserpb.User{
			ID:                  ret.ID,
			AppID:               ret.AppID,
			EmailAddress:        ret.EmailAddress,
			PhoneNO:             ret.PhoneNO,
			ImportedFromAppID:   ret.ImportedFromAppID,
			ImportedFromAppName: ret.ImportedFromAppName,
			ActionCredits:       ret.ActionCredits,
			AddressFieldsString: "[]",
			AddressFields:       nil,
			SigninVerifyTypeStr: basetypes.SignMethod_Email.String(),
			SigninVerifyType:    basetypes.SignMethod_Email,
		}
	)

	info, err := appusercli.CreateUser(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret1.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret1)
		amt.UserID = info.ID
	}
}

func createAnnouncement(t *testing.T) {
	info, err := CreateAnnouncement(context.Background(), &npool.AnnouncementReq{
		AppID:            &amt.AppID,
		LangID:           &amt.LangID,
		Title:            &amt.Title,
		Content:          &amt.Content,
		Channel:          &amt.Channel,
		AnnouncementType: &amt.AnnouncementType,
		EndAt:            &amt.EndAt,
	})
	if assert.Nil(t, err) {
		amt.CreatedAt = info.CreatedAt
		amt.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &amt)
	}
}

func getAnnouncementStates(t *testing.T) {
	channel1 := channelpb.NotifChannel_ChannelEmail

	_, err := announcementcrud.Create(context.Background(), &mgrpb.AnnouncementReq{
		ID:               &data.AnnouncementID,
		AppID:            &data.AppID,
		LangID:           &data.LangID,
		Title:            &data.Title,
		Content:          &data.Content,
		Channel:          &channel1,
		EndAt:            &endAt,
		AnnouncementType: &aType,
	})
	assert.Nil(t, err)

	userID1 := uuid.NewString()
	aType1 := mgrpb.AnnouncementType_Multicast
	_, err = announcementcrud.Create(context.Background(), &mgrpb.AnnouncementReq{
		ID:               &userID,
		AppID:            &data.AppID,
		LangID:           &data.LangID,
		Title:            &data.Title,
		Content:          &data.Content,
		Channel:          &channel1,
		EndAt:            &endAt,
		AnnouncementType: &aType1,
	})
	assert.Nil(t, err)

	_, err = readstatecrud.Create(context.Background(), &readstatemgrpb.ReadStateReq{
		AppID:          &data.AppID,
		UserID:         &userID,
		AnnouncementID: &data.AnnouncementID,
	})
	assert.Nil(t, err)

	_, err = readstatecrud.Create(context.Background(), &readstatemgrpb.ReadStateReq{
		AppID:          &data.AppID,
		UserID:         &userID,
		AnnouncementID: &data.AnnouncementID,
	})
	assert.Nil(t, err)

	_, err = userstatecrud.Create(context.Background(), &userstatemgrpb.UserReq{
		AppID:          &data.AppID,
		UserID:         &userID1,
		AnnouncementID: &data.AnnouncementID,
	})
	assert.Nil(t, err)

	infos, total, err := GetAnnouncementStates(context.Background(), &npool.Conds{
		AppID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.AppID,
		},
		UserID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		LangID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.LangID,
		},
	}, 0, 10)
	if assert.Nil(t, err) {
		data.CreatedAt = infos[0].CreatedAt
		data.UpdatedAt = infos[0].UpdatedAt
		assert.NotEqual(t, total, 1)
		assert.Equal(t, &data, infos[0])
	}

	_, total, err = GetAnnouncementStates(context.Background(), &npool.Conds{
		AppID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.AppID,
		},
		UserID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		LangID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.LangID,
		},
	}, 0, 10)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 2)
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
	t.Run("getAnnouncementStates", getAnnouncementStates)
}
