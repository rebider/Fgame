package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/shareboss/pbutil"
	"fgame/fgame/cross/shareboss/shareboss"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

//进入跨服boss场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()

	//TODO:shareboss:跨服世界boss场景
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossWorldBoss {
		return
	}

	//推送场景boss信息
	bossList := shareboss.GetShareBossService().GetShareBossListGroupByMap(worldbosstypes.BossTypeShareBoss, s.MapId())
	scShareBossListInfoNotice := pbutil.BuildSCShareBossListInfoNotice(bossList)
	pl.SendMsg(scShareBossListInfoNotice)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
