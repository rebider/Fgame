package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家使用物品
func playerItemUse(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*inventoryeventtypes.PlayerInventoryItemUseEventData)
	if !ok {
		return
	}

	itemId := useData.GetItemId()
	useNum := useData.GetUseNum()
	itemTemp := item.GetItemService().GetItem(int(itemId))
	if !itemTemp.IsYunYingItem() {
		return
	}

	// 更新排行榜信息
	updateUseItemTimesRankData(pl, useNum, itemTemp)
	return
}

// 次数排行榜
func updateUseItemTimesRankData(pl player.Player, useNum int32, itemTemp *gametemplate.ItemTemplate) {
	typ := welfaretypes.OpenActivityTypeRank
	subType := welfaretypes.OpenActivityRankSubTypeNumber
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		if !itemTemp.IsRelationToGroup(groupId) {
			continue
		}

		welfareManager.AddActivityNumRecordRecord(groupId, useNum)
	}
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeUseItem, event.EventListenerFunc(playerItemUse))
}
