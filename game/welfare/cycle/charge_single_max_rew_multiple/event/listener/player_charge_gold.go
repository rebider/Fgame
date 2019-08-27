package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	cyclechargesinglemaxrewmultipletemplate "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/template"
	cyclechargesinglemaxrewmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值元宝
func playerChargeSingleCycleMultiple(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	chargeNum, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeCycleCharge
	subType := welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple

	//每日单笔充值
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.RefreshActivityData(typ, subType)

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*cyclechargesinglemaxrewmultipletypes.CycleSingleChargeMaxRewMultipleInfo)
		if chargeNum > info.MaxSingleChargeNum {
			info.MaxSingleChargeNum = chargeNum
		}

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*cyclechargesinglemaxrewmultipletemplate.GroupTemplateCycleSingleMaxRewMultiple)
		descTempList := groupTemp.GetCurDayTempDescList(info.CycleDay)
		for _, temp := range descTempList {
			needGold := temp.Value2
			if chargeNum < needGold {
				continue
			}

			info.AddCanRewRecord(needGold)
			break
		}

		welfareManager.UpdateObj(obj)
		canRewRecord := info.LeftCanReceiveRewards()
		scMsg := pbutil.BuildSCOpenActivityCycleSingleChargeMaxRewMultipleInfoNotice(groupId, info.MaxSingleChargeNum, canRewRecord)
		pl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeSingleCycleMultiple))
}