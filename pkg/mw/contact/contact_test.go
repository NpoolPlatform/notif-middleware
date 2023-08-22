package contact

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

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
	ret = npool.Contact{
		AppID:          uuid.NewString(),
		Account:        "vagrant@163.com",
		AccountType:    basetypes.SignMethod_Email,
		AccountTypeStr: basetypes.SignMethod_Email.String(),
		UsedFor:        basetypes.UsedFor_Contact,
		UsedForStr:     basetypes.UsedFor_Contact.String(),
		Sender:         "vagrant2@163.com",
	}
)

func createContact(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(&ret.AppID),
		WithAccount(&ret.Account),
		WithAccountType(&ret.AccountType),
		WithUsedFor(&ret.UsedFor),
		WithSender(&ret.Sender),
	)
	assert.Nil(t, err)

	info, err := handler.CreateContact(context.Background())
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
	ret.AccountTypeStr = basetypes.SignMethod_Email.String()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithAccount(&ret.Account),
		WithSender(&ret.Sender),
		WithAccountType(&ret.AccountType),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateContact(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getContact(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetContact(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getContacts(t *testing.T) {
	conds := &npool.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UsedFor:     &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.UsedFor)},
		AccountType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.AccountType)},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetContacts(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteContact(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteContact(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetContact(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestContact(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createContact", createContact)
	t.Run("updateContact", updateContact)
	t.Run("getContact", getContact)
	t.Run("getContacts", getContacts)
	t.Run("deleteContact", deleteContact)
}
