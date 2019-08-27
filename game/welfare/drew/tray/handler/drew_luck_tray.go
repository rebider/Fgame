package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	drewcommontypes "fgame/fgame/game/welfare/drew/common/types"
	drewtraytemplate "fgame/fgame/game/welfare/drew/tray/template"
	drewtraytypes "fgame/fgame/game/welfare/drew/tray/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MERGE_ACTIVITY_LUCKY_DREW_ATTEND_TYPE), dispatch.HandlerFunc(handleAttendLuckyTray))

}

const (
	defaultNum = int32(1)
	batchNum   = int32(10)
)

//幸运抽奖
func handleAttendLuckyTray(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:幸运抽奖")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMergeActivityLuckDrewAttend := msg.(*uipb.CSMergeActivityLuckDrewAttend)
	groupId := csMergeActivityLuckDrewAttend.GetGroupId()
	attendType := drewcommontypes.LuckyDrewAttendType(csMergeActivityLuckDrewAttend.GetAttendType())

	if !attendType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"attendType": attendType,
			}).Warn("welfare:幸运抽奖错误,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = attendLuckyTray(tpl, groupId, attendType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("welfare:处理幸运抽奖,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理幸运抽奖完成")
	return nil

}

//幸运抽奖逻辑
func attendLuckyTray(pl player.Player, groupId int32, attendType drewcommontypes.LuckyDrewAttendType) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeTray

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:幸运抽奖错误，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	//是否足够元宝
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:幸运抽奖错误，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	groupTemp := groupInterface.(*drewtraytemplate.GroupTemplateDrewTray)
	needGold := int64(groupTemp.GetLuckyDrewNeedGold(attendType))
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(needGold, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("welfare:幸运抽奖错误，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	attendNum := defaultNum
	if attendType == drewcommontypes.LuckyDrewTypeBatch {
		attendNum = batchNum
	}
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if attendNum > 0 {
		if inventoryManager.GetEmptySlots(inventorytypes.BagTypePrim) < attendNum {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("welfare:充值抽奖错误，空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnoughSlot, fmt.Sprintf("%d", attendNum))
			return
		}
	}

	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*drewtraytypes.LuckyDrewInfo)
	luckTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplate(groupId)
	if luckTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:幸运抽奖错误，抽奖模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
	}
	timesMap := luckTemp.GetRewDropByTimesMap()
	timesDescList := luckTemp.GetTimesDesc()

	// 计算物品
	var totalItemList []*droptemplate.DropItemData
	var dropItemList []*droptemplate.DropItemData
	var extraItemList []*droptemplate.DropItemData
	curAttendNum := info.AttendTimes
	for index := int32(0); index < attendNum; index++ {
		curAttendNum += 1
		dropId := luckTemp.DropId
		for _, times := range timesDescList {
			ret := curAttendNum % int32(times)
			if ret == 0 {
				dropId = timesMap[int32(times)]
				break
			}
		}

		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
		if dropData == nil {
			log.WithField("dropId", dropId).Warn("掉落包随机为空")
			continue
		}
		dropData.BindType = itemtypes.ItemBindTypeUnBind
		totalItemList = append(totalItemList, dropData)
		dropItemList = append(totalItemList, dropData)
		//额外奖励
		giveItemMap := luckTemp.GetGiveItemMap()
		if len(giveItemMap) > 0 {
			giveItemDataList := droptemplate.ConvertToItemDataList(giveItemMap, itemtypes.ItemBindTypeUnBind)
			totalItemList = append(totalItemList, giveItemDataList...)
			extraItemList = append(totalItemList, giveItemDataList...)
		}
		//道具公告
		itemId := dropData.GetItemId()
		num := dropData.GetNum()
		inventorylogic.PrecioustemBroadcast(pl, itemId, num, lang.InventoryLuckyTrayItemNotice)
	}

	var newItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(totalItemList) != 0 {
		newItemList, resMap = droplogic.SeperateItemDatas(totalItemList)
	}

	// 背包空间
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"len":      len(newItemList),
			}).Warn("welfare:幸运抽奖错误，空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonDrewUse
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), subType)
		flag := propertyManager.CostGold(needGold, false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("welfare:消耗元宝应该成功")
		}
	}

	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonOpenActivityRew
		silverReason := commonlog.SilverLogReasonOpenActivityRew
		levelReason := commonlog.LevelLogReasonOpenActivityRew
		goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
		silverReasonText := fmt.Sprintf(silverReason.String(), typ, subType)
		levelReasonText := fmt.Sprintf(levelReason.String(), typ, subType)
		err = droplogic.AddRes(pl, resMap, goldReason, goldReasonText, silverReason, silverReasonText, levelReason, levelReasonText)
		if err != nil {
			return
		}
	}

	if len(newItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
		flag = inventoryManager.BatchAddOfItemLevel(newItemList, itemGetReason, itemGetReasonText)
		if !flag {
			panic("welfare:增加物品应该成功")
		}
	}

	// 更新次数
	info.AttendTimes += attendNum
	welfareManager.UpdateObj(obj)
	eventData := welfareeventtypes.CreatePlayerAttendDrewEventData(groupId, attendNum)
	gameevent.Emit(welfareeventtypes.EventTypeAttendDrew, pl, eventData)

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMergeActivityLuckDrewAttends := pbutil.BuildSCMergeActivityLuckDrewAttends(dropItemList, extraItemList, groupId)
	pl.SendMsg(scMergeActivityLuckDrewAttends)
	return
}
