package goodbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	goodbenefit1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/goodbenefit"
	"github.com/google/uuid"
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
	yesterday = uint32(time.Now().AddDate(0, 0, -1).Unix())
	ret       = npool.GoodBenefit{
		GoodID:      uuid.NewString(),
		GoodName:    uuid.NewString(),
		Amount:      "100",
		State:       basetypes.Result_Success,
		StateStr:    basetypes.Result_Success.String(),
		Message:     uuid.NewString(),
		BenefitDate: yesterday,
		TxID:        uuid.NewString(),
		Notified:    false,
	}

	ret2 = npool.GoodBenefit{
		GoodID:      uuid.NewString(),
		GoodName:    uuid.NewString(),
		Amount:      "10",
		State:       basetypes.Result_Success,
		StateStr:    basetypes.Result_Success.String(),
		Message:     uuid.NewString(),
		BenefitDate: uint32(time.Now().Add(-3 * time.Minute).Unix()),
		TxID:        uuid.NewString(),
		Notified:    false,
	}
)

func createGoodBenefit(t *testing.T) {
	info, err := CreateGoodBenefit(context.Background(), &npool.GoodBenefitReq{
		GoodID:      &ret.GoodID,
		GoodName:    &ret.GoodName,
		Amount:      &ret.Amount,
		State:       &ret.State,
		Message:     &ret.Message,
		BenefitDate: &ret.BenefitDate,
		TxID:        &ret.TxID,
		Notified:    &ret.Notified,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}

	_info, err := CreateGoodBenefit(context.Background(), &npool.GoodBenefitReq{
		GoodID:      &ret2.GoodID,
		GoodName:    &ret2.GoodName,
		Amount:      &ret2.Amount,
		State:       &ret2.State,
		Message:     &ret2.Message,
		BenefitDate: &ret2.BenefitDate,
		TxID:        &ret2.TxID,
		Notified:    &ret2.Notified,
	})
	if assert.Nil(t, err) {
		ret2.CreatedAt = _info.CreatedAt
		ret2.UpdatedAt = _info.UpdatedAt
		ret2.ID = _info.ID
		assert.Equal(t, _info, &ret2)
	}
}

func updateGoodBenefit(t *testing.T) {
	ret.Notified = true
	info, err := UpdateGoodBenefit(context.Background(), &npool.GoodBenefitReq{
		ID:       &ret.ID,
		Notified: &ret.Notified,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getGoodBenefit(t *testing.T) {
	info, err := GetGoodBenefit(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getGoodBenefits(t *testing.T) {
	infos, _, err := GetGoodBenefits(context.Background(), &npool.Conds{
		Notified: &basetypes.BoolVal{
			Op:    cruder.EQ,
			Value: ret.Notified,
		},
		BenefitDateStart: &basetypes.Uint32Val{
			Op:    cruder.LTE,
			Value: yesterday,
		},
		BenefitDateEnd: &basetypes.Uint32Val{
			Op:    cruder.GTE,
			Value: uint32(time.Now().Unix()),
		},
	}, 0, 2)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 2)
	}
}

func deleteGoodBenefit(t *testing.T) {
	handler, err := goodbenefit1.NewHandler(
		context.Background(),
		goodbenefit1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteGoodBenefit(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = handler.GetGoodBenefit(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)

	_id, err := uuid.Parse(ret2.ID)
	assert.Nil(t, err)
	handler.ID = &_id

	_info, err := handler.DeleteGoodBenefit(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _info)

	_info, err = handler.GetGoodBenefit(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, _info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createGoodBenefit", createGoodBenefit)
	t.Run("updateGoodBenefit", updateGoodBenefit)
	t.Run("getGoodBenefits", getGoodBenefit)
	t.Run("getGoodBenefits", getGoodBenefits)
	t.Run("deleteGoodBenefit", deleteGoodBenefit)
}
