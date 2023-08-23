package tx

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
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
	ret = npool.Tx{
		ID:            uuid.NewString(),
		TxID:          uuid.NewString(),
		NotifState:    npool.TxState_WaitNotified,
		NotifStateStr: npool.TxState_WaitNotified.String(),
		TxType:        basetypes.TxType_TxWithdraw,
		TxTypeStr:     basetypes.TxType_TxWithdraw.String(),
	}
)

func createTx(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithTxID(&ret.TxID),
		WithNotifState(&ret.NotifState),
		WithTxType(&ret.TxType),
	)
	assert.Nil(t, err)

	info, err := handler.CreateTx(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateTx(t *testing.T) {
	ret.NotifState = npool.TxState_Notified
	ret.NotifStateStr = npool.TxState_Notified.String()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithNotifState(&ret.NotifState),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateTx(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getTx(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetTx(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getTxs(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		TxID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.TxID},
		NotifState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.NotifState)},
		TxTypes:    &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{uint32(ret.TxType)}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetTxs(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteTx(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteTx(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetTx(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createTx", createTx)
	t.Run("updateTx", updateTx)
	t.Run("getTx", getTx)
	t.Run("getTxs", getTxs)
	t.Run("deleteTx", deleteTx)
}
