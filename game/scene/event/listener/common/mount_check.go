package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//同步玩家
func playerMountCheck(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	s := p.GetScene()
	if s == nil {
		return
	}
	if !p.IsMove() {
		return
	}
	if !p.IsGuaJi() {
		return
	}
	if s.MapTemplate().LimitRideHorse == 0 {
		//骑马
		if p.IsMountHidden() && p.GetMountId() != 0 {
			p.MountHidden(false)
		}

	}
	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerMountCheck, event.EventListenerFunc(playerMountCheck))
}