package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家hp变化
func playerHPChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenapvp {
		return
	}
	scArenaPlayerDataHPChanged := pbutil.BuildSCArenapvpPlayerShowDataHpChanged(pl)
	s.BroadcastMsg(scArenaPlayerDataHPChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerHPChanged, event.EventListenerFunc(playerHPChanged))
}