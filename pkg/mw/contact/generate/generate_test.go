package generate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	contactmw "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	contact "github.com/NpoolPlatform/notif-middleware/pkg/mw/contact"
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
	ret = contactmw.Contact{
		AppID:          uuid.NewString(),
		Account:        "vagrant1@163.com",
		AccountType:    basetypes.SignMethod_Email,
		AccountTypeStr: basetypes.SignMethod_Email.String(),
		UsedFor:        basetypes.UsedFor_Contact,
		UsedForStr:     basetypes.UsedFor_Contact.String(),
		Sender:         "vagrant3@163.com",
	}

	subject    = uuid.NewString()
	body       = uuid.NewString()
	senderName = uuid.NewString()
)

func setupGenerateContact(t *testing.T) func(*testing.T) {
	handler, err := NewHandler(
		context.Background(),
		contact.WithAppID(&ret.AppID),
		contact.WithAccount(&ret.Account),
		contact.WithAccountType(&ret.AccountType),
		contact.WithUsedFor(&ret.UsedFor),
		contact.WithSender(&ret.Sender),
	)
	assert.Nil(t, err)

	_, err = handler.CreateContact(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = handler.DeleteContact(context.Background())
	}
}

func generateContact(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		contact.WithAppID(&ret.AppID),
		contact.WithUsedFor(&ret.UsedFor),
		contact.WithSender(&ret.Sender),
		WithSubject(&subject),
		WithBody(&body),
		WithSenderName(&senderName),
	)
	assert.Nil(t, err)

	info, err := handler.GenerateContact(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func TestGenerateContact(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupGenerateContact(t)
	defer teardown(t)

	t.Run("generateContact", generateContact)
}
