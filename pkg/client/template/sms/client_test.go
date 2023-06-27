package sms

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

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
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
	ret = npool.SMSTemplate{
		ID:         uuid.NewString(),
		AppID:      uuid.NewString(),
		LangID:     uuid.NewString(),
		UsedFor:    basetypes.UsedFor_KYCApproved,
		UsedForStr: basetypes.UsedFor_KYCApproved.String(),
		Subject:    "Subject " + uuid.NewString(),
		Message:    "Message " + uuid.NewString(),
	}
)

var (
	appInfo = npool.SMSTemplateReq{
		ID:      &ret.ID,
		AppID:   &ret.AppID,
		LangID:  &ret.LangID,
		UsedFor: &ret.UsedFor,
		Subject: &ret.Subject,
		Message: &ret.Message,
	}
)

var info *npool.SMSTemplate

func createSMSTemplate(t *testing.T) {
	var err error
	info, err = CreateSMSTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func createSMSTemplates(t *testing.T) {
	rets := []npool.SMSTemplate{
		{
			ID:         uuid.NewString(),
			AppID:      ret.AppID,
			LangID:     uuid.NewString(),
			UsedFor:    basetypes.UsedFor_Signin,
			UsedForStr: basetypes.UsedFor_Signin.String(),
			Subject:    "Subject1 " + uuid.NewString(),
			Message:    "Message1 " + uuid.NewString(),
		},
		{
			ID:         uuid.NewString(),
			AppID:      ret.AppID,
			LangID:     uuid.NewString(),
			UsedFor:    basetypes.UsedFor_KYCRejected,
			UsedForStr: basetypes.UsedFor_KYCRejected.String(),
			Subject:    "Subject2 " + uuid.NewString(),
			Message:    "Message2 " + uuid.NewString(),
		},
	}

	apps := []*npool.SMSTemplateReq{}
	for key := range rets {
		apps = append(apps, &npool.SMSTemplateReq{
			ID:      &rets[key].ID,
			AppID:   &rets[key].AppID,
			LangID:  &rets[key].LangID,
			UsedFor: &rets[key].UsedFor,
			Subject: &rets[key].Subject,
			Message: &rets[key].Message,
		})
	}

	infos, err := CreateSMSTemplates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateSMSTemplate(t *testing.T) {
	var err error
	info, err = UpdateSMSTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSMSTemplate(t *testing.T) {
	var err error
	info, err = GetSMSTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSMSTemplates(t *testing.T) {
	infos, total, err := GetSMSTemplates(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &ret)
	}
}

func getSMSTemplateOnly(t *testing.T) {
	var err error
	info, err = GetSMSTemplateOnly(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func existSMSTemplate(t *testing.T) {
	exist, err := ExistSMSTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existSMSTemplateConds(t *testing.T) {
	exist, err := ExistSMSTemplateConds(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteSMSTemplate(t *testing.T) {
	info, err := DeleteSMSTemplate(context.Background(), &npool.SMSTemplateReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createSMSTemplate", createSMSTemplate)
	t.Run("createSMSTemplates", createSMSTemplates)
	t.Run("getSMSTemplate", getSMSTemplate)
	t.Run("getSMSTemplates", getSMSTemplates)
	t.Run("getSMSTemplateOnly", getSMSTemplateOnly)
	t.Run("updateSMSTemplate", updateSMSTemplate)
	t.Run("existSMSTemplate", existSMSTemplate)
	t.Run("existSMSTemplateConds", existSMSTemplateConds)
	t.Run("delete", deleteSMSTemplate)
}
