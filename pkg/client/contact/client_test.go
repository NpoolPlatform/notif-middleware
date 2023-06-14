package contact

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	ret   = npool.Contact{
		AppID:          appID,
		Account:        "vagrant@163.com",
		AccountType:    basetypes.SignMethod_Email,
		AccountTypeStr: basetypes.SignMethod_Email.String(),
		UsedFor:        basetypes.UsedFor_Contact,
		UsedForStr:     basetypes.UsedFor_Contact.String(),
		Sender:         "vagrant2@163.com",
	}
)

func setupContact(t *testing.T) func(*testing.T) {
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

	ret.AppID = app1.ID

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
	}
}

func createContact(t *testing.T) {
	info, err := CreateContact(context.Background(), &npool.ContactReq{
		AppID:       &ret.AppID,
		Account:     &ret.Account,
		AccountType: &ret.AccountType,
		UsedFor:     &ret.UsedFor,
		Sender:      &ret.Sender,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateContact(t *testing.T) {
	ret.Account = "aaaa@123.com"
	ret.Sender = "bbbb@123.com"
	ret.AccountType = basetypes.SignMethod_Email
	info, err := UpdateContact(context.Background(), &npool.ContactReq{
		ID:          &ret.ID,
		Account:     &ret.Account,
		Sender:      &ret.Sender,
		AccountType: &ret.AccountType,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getContact(t *testing.T) {
	info, err := GetContact(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getContacts(t *testing.T) {
	infos, _, err := GetContacts(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		UsedFor: &basetypes.Int32Val{
			Op:    cruder.EQ,
			Value: int32(ret.UsedFor),
		},
		AccountType: &basetypes.Int32Val{
			Op:    cruder.EQ,
			Value: int32(ret.AccountType),
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func generateContact(t *testing.T) {
	info, err := GenerateContact(context.Background(), &npool.GenerateContactRequest{
		AppID:      ret.AppID,
		UsedFor:    ret.UsedFor,
		Sender:     ret.Sender,
		Subject:    uuid.NewString(),
		Body:       uuid.NewString(),
		SenderName: uuid.NewString(),
	})

	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func deleteContact(t *testing.T) {
	info, err := DeleteContact(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetContact(context.Background(), info.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupContact(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	t.Run("createContact", createContact)
	t.Run("updateContact", updateContact)
	t.Run("getContact", getContact)
	t.Run("getContacts", getContacts)
	t.Run("generateContact", generateContact)
	t.Run("deleteContact", deleteContact)

	patch.Unpatch()
}
