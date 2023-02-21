package replace

import (
	"strings"
	"time"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"
)

const (
	NotifTemplateVarName     = "{{ NAME }}"
	NotifTemplateVarMessage  = "{{ MESSAGE }}"
	NotifTemplateVarAmount   = "{{ AMOUNT }}"
	NotifTemplateVarCoinUnit = "{{ COIN_UNIT }}"
	NotifTemplateVarDate     = "{{ DATE }}"
	NotifTemplateVarTime     = "{{ TIME }}"
	NotifTemplateVarAddress  = "{{ ADDRESS }}"
	NotifTemplateVarCode     = "{{ CODE }}"
)

func ReplaceAll(pattern string, vars *npool.TemplateVars) string {
	if vars == nil {
		return pattern
	}

	if vars.Username != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarName, *vars.Username)
	}
	if vars.Message != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarMessage, *vars.Message)
	}
	if vars.Amount != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarAmount, *vars.Amount)
	}
	if vars.CoinUnit != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarCoinUnit, *vars.CoinUnit)
	}
	if vars.Timestamp != nil {
		datetime := time.Unix(int64(*vars.Timestamp), 0)
		date := datetime.Format("1970-01-01")
		time1 := datetime.Format("00:00:01")

		pattern = strings.ReplaceAll(pattern, NotifTemplateVarDate, date)
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarTime, time1)
	}
	if vars.Address != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarAddress, *vars.Address)
	}
	if vars.Code != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarCode, *vars.Code)
	}

	return pattern
}
