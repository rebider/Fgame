package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
)

//玩家取消排队
func godSiegeCancleLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	eventData := data.(*godsiegeeventtypes.GodSiegeCancleLineUpEventData)
	pos := eventData.GetPos()
	godType := eventData.GetGodType()
	godsiegelogic.BroadGodSiegeLineUpChanged(int32(godType), pos, lineList)
	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegeCancleLineUp, event.EventListenerFunc(godSiegeCancleLineUp))
}
