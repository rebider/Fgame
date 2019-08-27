package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shangguzhilinglogic "fgame/fgame/game/shangguzhiling/logic"
	"fgame/fgame/game/shangguzhiling/pbutil"
	playershangguzhiling "fgame/fgame/game/shangguzhiling/player"
	shangguzhilingtemplate "fgame/fgame/game/shangguzhiling/template"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_UPLEVEL_TYPE), dispatch.HandlerFunc(handleShangguzhilingUplevel))
}

func handleShangguzhilingUplevel(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSShangguzhilingUplevel)
	lingshouType := shangguzhilingtypes.LingshouType(csMsg.GetType())
	autoFlag := csMsg.GetAutoFlag()
	itemId := csMsg.GetItemId()

	if !lingshouType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:上古之灵升级请求,灵兽类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = shangguzhilingUplevel(tpl, lingshouType, autoFlag, itemId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shangguzhiling:上古之灵升级请求,错误")

		return err
	}
	return nil
}

func shangguzhilingUplevel(pl player.Player, lingshouType shangguzhilingtypes.LingshouType, autoFlag bool, useItemId int32) (err error) {
	lingShouManager := pl.GetPlayerDataManager(playertypes.PlayerShangguzhilingDataManagerType).(*playershangguzhiling.PlayerShangguzhilingDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//是否解锁
	if !lingShouManager.IsLingShouUnlock(lingshouType) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:上古之灵升级请求,灵兽未解锁")
		playerlogic.SendSystemMessage(pl, lang.ShangguzhilingLingShouUnLock)
		return
	}
	obj := lingShouManager.GetLingShouObj(lingshouType)

	//相关模板
	curLevel := obj.GetLevel()
	baseTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingShouTemplate(lingshouType)
	if baseTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:上古之灵升级请求,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	levelTemp := baseTemp.GetLevelTemp(curLevel + 1)
	if levelTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
				"curLevel":     curLevel,
			}).Warn("shangguzhiling:上古之灵升级请求,下一级模板不存在(满级)")
		playerlogic.SendSystemMessage(pl, lang.ShangguzhilingLingShouFullLevel)
		return
	}

	//使用的物品ID
	canUseItemList := baseTemp.GetLingShouUpLevelUseItemList()
	flag := false
	for _, itemId := range canUseItemList {
		if itemId == useItemId {
			flag = true
		}
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
				"useItemId":    useItemId,
			}).Warn("shangguzhiling:上古之灵升级请求,选择使用的物品Id错误")
		playerlogic.SendSystemMessage(pl, lang.ShangguzhilingUseItemIdWrong)
		return
	}
	useItemTemp := item.GetItemService().GetItem(int(useItemId))
	if useItemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
				"useItemId":    useItemId,
			}).Warn("shangguzhiling:上古之灵升级请求,使用的物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//物品数量
	itemCount := inventoryManager.NumOfItems(useItemId)
	useItemCount := int32(1)
	finalUseItemNum := useItemCount
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	needGold := int64(0)
	needBindGold := int64(0)
	needSilver := int64(0)
	if itemCount < useItemCount {
		if !autoFlag {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"error":        err,
					"lingshouType": lingshouType,
					"itemCount":    itemCount,
					"useItemCount": useItemCount,
				}).Warn("shangguzhiling:上古之灵升级请求,选择使用的物品数量不足")
			playerlogic.SendSystemMessage(pl, lang.ShangguzhilingUseItemCountNotEnough)
			return
		}
		//自动进阶
		needBuyNum := useItemCount - itemCount
		finalUseItemNum = itemCount
		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItemId) {
				log.WithFields(log.Fields{
					"playerId":  pl.GetId(),
					"useItemId": useItemId,
					"autoFlag":  autoFlag,
				}).Warn("shangguzhiling:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItemId, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId":  pl.GetId(),
					"useItemId": useItemId,
					"autoFlag":  autoFlag,
				}).Warn("shangguzhiling:购买物品失败,上古之灵升级失败")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNumInvalid)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			needGold += shopNeedGold
			needBindGold += shopNeedBindGold
			needSilver += shopNeedSilver
		}
	}
	//是否足够银两
	flag = propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shangguzhiling:上古之灵，银两不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(needGold, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shangguzhiling:上古之灵，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//是否足够绑元
	needCostBindGold := needBindGold + needGold
	flag = propertyManager.HasEnoughGold(needCostBindGold, true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shangguzhiling:上古之灵，绑元不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//-----------------------分割线-----------------------

	//增加经验
	exp := int64(useItemTemp.TypeFlag1)
	allexp := exp * int64(useItemCount)
	flag = lingShouManager.AddExp(lingshouType, allexp)
	if !flag {
		panic("shangguzhiling:灵兽增加经验应该成功")
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//自动购买消耗金钱
	//消耗银两
	if needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonLingShouUplevel
		silverUseReasonText := fmt.Sprintf(silverUseReason.String(), lingshouType.String())
		flag := propertyManager.CostSilver(needSilver, silverUseReason, silverUseReasonText)
		if !flag {
			panic("shangguzhiling:消耗银两应该成功")
		}
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonLingshouUplevel
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), lingshouType.String())
		flag := propertyManager.CostGold(needGold, false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("shangguzhiling:消耗元宝应该成功")
		}
	}

	//消耗绑元
	if needBindGold > 0 {
		goldUseReason := commonlog.GoldLogReasonLingshouUplevel
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), lingshouType.String())
		flag := propertyManager.CostGold(needBindGold, true, goldUseReason, goldUseReasonText)
		if !flag {
			panic("shangguzhiling:消耗元宝应该成功")
		}
	}

	//同步物品（删掉吃掉的物品）
	useReason := commonlog.InventoryLogReasonLingShouUplevel
	useReasonText := fmt.Sprintf(useReason.String(), lingshouType.String())
	if finalUseItemNum > 0 {
		flag = inventoryManager.UseItem(useItemId, finalUseItemNum, useReason, useReasonText)
		if !flag {
			panic("inventory:移除物品应该是可以的")
		}
	}

	shangguzhilinglogic.LingShouPropertyChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCShangguzhilingUplevel(lingshouType, obj.GetLevel(), obj.GetExperience())
	pl.SendMsg(scMsg)
	return
}
