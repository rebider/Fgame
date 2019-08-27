package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopenlogic "fgame/fgame/game/funcopen/logic"
	"fgame/fgame/game/funcopen/pbutil"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

//玩家基础属性变更
func playerLevelChanged(target event.EventTarget, data event.EventData) (err error) {

	p := target.(player.Player)

	updateList, err := funcopenlogic.CheckFuncOpen(p)
	if err != nil {
		return
	}
	if len(updateList) != 0 {
		//TODO 优化
		//更新部分作用器属性
		p.UpdateBattleProperty(playerpropertytypes.PropertyEffectorTypeMaskAll)

		scFuncOpenUpdateList := pbutil.BuildSCFuncOpenUpdateList(updateList)
		p.SendMsg(scFuncOpenUpdateList)

	}
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerLevelChanged, event.EventListenerFunc(playerLevelChanged))
}
